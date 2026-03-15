package goconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type postFillKind int

const (
	postFillDuration postFillKind = iota + 1
	postFillSlice
	postFillFloat32
	postFillInt32
	postFillUint32
)

type postFillArgs struct {
	item
	Raw  *string
	Kind postFillKind
}

func FillArgs(c interface{}, args []string) error {
	return fillArgs(c, args, defaultProgramName, defaultConfigFlagName)
}

func fillArgs(c interface{}, args []string, programName string, configFlagName string) error {
	if err := validateConfigTarget(c); err != nil {
		return err
	}

	if programName == "" {
		programName = defaultProgramName
	}

	var f = flag.NewFlagSet(programName, flag.ContinueOnError)
	f.Usage = func() {}
	f.SetOutput(io.Discard)

	// Default config flag
	if configFlagName != "" {
		f.String(configFlagName, "", "Configuration JSON file")
	}

	post := []postFillArgs{}

	traverse(c, func(i item) {
		name_path := strings.ToLower(strings.Join(i.Path, "."))
		env_name := strings.ToUpper(strings.Join(i.Path, "_"))

		usage := i.Usage
		if usage != "" {
			usage += " "
		}
		usage += "[env " + env_name + "]"

		if reflect.TypeOf(time.Duration(0)) == i.Value.Type() {
			value := ""
			f.StringVar(&value, name_path, i.Value.Interface().(time.Duration).String(), usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
				Kind: postFillDuration,
			})

		} else if reflect.Bool == i.Kind {
			f.BoolVar(i.Ptr.(*bool), name_path, i.Value.Interface().(bool), usage)

		} else if reflect.Float64 == i.Kind {
			f.Float64Var(i.Ptr.(*float64), name_path, i.Value.Interface().(float64), usage)

		} else if reflect.Float32 == i.Kind {
			value := strconv.FormatFloat(float64(i.Value.Interface().(float32)), 'g', -1, 32)
			f.StringVar(&value, name_path, value, usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
				Kind: postFillFloat32,
			})

		} else if reflect.Int64 == i.Kind {
			f.Int64Var(i.Ptr.(*int64), name_path, i.Value.Interface().(int64), usage)

		} else if reflect.Int32 == i.Kind {
			value := strconv.FormatInt(int64(i.Value.Interface().(int32)), 10)
			f.StringVar(&value, name_path, value, usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
				Kind: postFillInt32,
			})

		} else if reflect.Int == i.Kind {
			f.IntVar(i.Ptr.(*int), name_path, i.Value.Interface().(int), usage)

		} else if reflect.String == i.Kind {
			f.StringVar(i.Ptr.(*string), name_path, i.Value.Interface().(string), usage)

		} else if reflect.Uint64 == i.Kind {
			f.Uint64Var(i.Ptr.(*uint64), name_path, i.Value.Interface().(uint64), usage)

		} else if reflect.Uint32 == i.Kind {
			value := strconv.FormatUint(uint64(i.Value.Interface().(uint32)), 10)
			f.StringVar(&value, name_path, value, usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
				Kind: postFillUint32,
			})

		} else if reflect.Uint == i.Kind {
			f.UintVar(i.Ptr.(*uint), name_path, i.Value.Interface().(uint), usage)

		} else if reflect.Slice == i.Kind {

			b, _ := json.Marshal(i.Value.Interface())

			value := ""
			f.StringVar(&value, name_path, string(b), usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
				Kind: postFillSlice,
			})

		} else {
			panic("Kind `" + i.Kind.String() +
				"` is not supported by goconfig (field `" + i.FieldName + "`)")
		}

	})

	knownArgs := filterKnownArgs(args, f)

	if err := f.Parse(knownArgs); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			m := bytes.NewBufferString("Usage of " + programName + ":\n\n")
			f.SetOutput(m)
			f.PrintDefaults()
			return errors.New(m.String())
		}

		return err
	}

	// Postprocess flags: unsupported flags needs to be declared as string
	// and parsed later. Here is the place.
	for _, p := range post {
		switch p.Kind {
		case postFillDuration:
			d, err := unmarshalDurationString(*p.Raw)
			if err != nil {
				return fmt.Errorf(
					"'%s' should be nanoseconds or a time.Duration string: %s",
					p.FieldName, err.Error(),
				)
			}
			p.Value.SetInt(int64(d))

		case postFillSlice:
			err := json.Unmarshal([]byte(*p.Raw), p.Ptr)
			if err != nil {
				return errors.New(fmt.Sprintf(
					"'%s' should be a JSON array: %s",
					p.FieldName, err.Error(),
				))
			}

		case postFillFloat32:
			v, err := strconv.ParseFloat(*p.Raw, 32)
			if err != nil {
				return fmt.Errorf("'%s' should be a valid float32: %s", p.FieldName, err.Error())
			}
			w := float32(v)
			set(p.Ptr, &w)

		case postFillInt32:
			v, err := strconv.ParseInt(*p.Raw, 10, 32)
			if err != nil {
				return fmt.Errorf("'%s' should be a valid int32: %s", p.FieldName, err.Error())
			}
			w := int32(v)
			set(p.Ptr, &w)

		case postFillUint32:
			v, err := strconv.ParseUint(*p.Raw, 10, 32)
			if err != nil {
				return fmt.Errorf("'%s' should be a valid uint32: %s", p.FieldName, err.Error())
			}
			w := uint32(v)
			set(p.Ptr, &w)

		default:
			return fmt.Errorf("unsupported post-processing type for field '%s'", p.FieldName)

		}
	}

	return nil
}

func filterKnownArgs(args []string, f *flag.FlagSet) []string {
	known := make([]string, 0, len(args))

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--" {
			break
		}

		if !strings.HasPrefix(arg, "-") || arg == "-" {
			continue
		}

		name := strings.TrimLeft(arg, "-")
		inlineValue := false
		if idx := strings.Index(name, "="); idx != -1 {
			name = name[:idx]
			inlineValue = true
		}

		if name == "h" || name == "help" {
			known = append(known, "-help")
			continue
		}

		fl := f.Lookup(name)
		if fl == nil {
			continue
		}

		known = append(known, arg)
		if inlineValue {
			continue
		}

		if b, ok := fl.Value.(interface{ IsBoolFlag() bool }); ok && b.IsBoolFlag() {
			if i+1 < len(args) && !isLikelyFlag(args[i+1]) {
				known = append(known, args[i+1])
				i++
			}
			continue
		}

		if i+1 < len(args) {
			known = append(known, args[i+1])
			i++
		}
	}

	return known
}

func isLikelyFlag(token string) bool {
	return strings.HasPrefix(token, "-") && token != "-"
}
