package goconfig

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	defaultProgramName        = "goconfig"
	defaultConfigFlagName     = "config"
	defaultImplicitConfigFile = "config.json"
)

// Option customizes Load behaviour.
type Option func(*Options)

// Options controls how configuration is loaded.
type Options struct {
	Args               []string
	ProgramName        string
	ConfigFile         string
	ConfigFlagName     string
	AutoConfigFile     bool
	AutoConfigFilename string
	EnvLookup          func(string) (string, bool)
}

// WithArgs sets the command-line arguments used by Load.
func WithArgs(args []string) Option {
	argsCopy := append([]string(nil), args...)
	return func(o *Options) {
		o.Args = argsCopy
	}
}

// WithProgramName sets the name used in generated help output.
func WithProgramName(name string) Option {
	return func(o *Options) {
		o.ProgramName = name
	}
}

// WithConfigFile sets the JSON file used by Load.
func WithConfigFile(filename string) Option {
	return func(o *Options) {
		o.ConfigFile = filename
	}
}

// WithConfigFlagName changes the command-line flag used to point to a config file.
func WithConfigFlagName(name string) Option {
	return func(o *Options) {
		o.ConfigFlagName = name
	}
}

// WithImplicitConfigFile enables and sets the fallback JSON config filename.
func WithImplicitConfigFile(filename string) Option {
	return func(o *Options) {
		o.AutoConfigFile = true
		o.AutoConfigFilename = filename
	}
}

// WithoutImplicitConfigFile disables automatic loading of config.json.
func WithoutImplicitConfigFile() Option {
	return func(o *Options) {
		o.AutoConfigFile = false
	}
}

// WithEnvLookup sets a custom environment lookup function.
func WithEnvLookup(lookup func(string) (string, bool)) Option {
	return func(o *Options) {
		o.EnvLookup = lookup
	}
}

// Load populates c from JSON config file, environment variables and command-line flags.
//
// Priority order (highest to lowest): flags > environment > JSON file > defaults.
func Load(c interface{}, opts ...Option) error {
	if err := validateConfigTarget(c); err != nil {
		return err
	}

	options := defaultOptions()
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}

	filename, err := resolveConfigFile(options)
	if err != nil {
		return errors.New("Config arg error: " + err.Error())
	}

	if err := FillJson(c, filename); err != nil {
		return errors.New("Config file error: " + err.Error())
	}

	if err := fillEnvironmentsWithLookup(c, options.EnvLookup); err != nil {
		return errors.New("Config env error: " + err.Error())
	}

	if err := fillArgs(c, options.Args, options.ProgramName, options.ConfigFlagName); err != nil {
		return errors.New("Config arg error: " + err.Error())
	}

	return nil
}

func defaultOptions() Options {
	programName := defaultProgramName
	if len(os.Args) > 0 && os.Args[0] != "" {
		programName = os.Args[0]
	}

	args := []string{}
	if len(os.Args) > 1 {
		args = append(args, os.Args[1:]...)
	}

	return Options{
		Args:               args,
		ProgramName:        programName,
		ConfigFlagName:     defaultConfigFlagName,
		AutoConfigFile:     true,
		AutoConfigFilename: defaultImplicitConfigFile,
		EnvLookup:          os.LookupEnv,
	}
}

func resolveConfigFile(options Options) (string, error) {
	if options.ConfigFile != "" {
		return options.ConfigFile, nil
	}

	if options.ConfigFlagName != "" {
		fromArgs, ok, err := getStringFlagValue(options.Args, options.ConfigFlagName)
		if err != nil {
			return "", err
		}
		if ok {
			return fromArgs, nil
		}
	}

	if options.AutoConfigFile && options.AutoConfigFilename != "" {
		if _, err := os.Stat(options.AutoConfigFilename); err == nil {
			return options.AutoConfigFilename, nil
		}
	}

	return "", nil
}

func getStringFlagValue(args []string, flagName string) (string, bool, error) {
	if flagName == "" {
		return "", false, nil
	}

	short := "-" + flagName
	long := "--" + flagName

	found := false
	value := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == short || arg == long:
			if i+1 >= len(args) {
				return "", false, fmt.Errorf("flag needs an argument: %s", short)
			}
			value = args[i+1]
			found = true
			i++

		case strings.HasPrefix(arg, short+"="):
			value = strings.TrimPrefix(arg, short+"=")
			found = true

		case strings.HasPrefix(arg, long+"="):
			value = strings.TrimPrefix(arg, long+"=")
			found = true
		}
	}

	return value, found, nil
}
