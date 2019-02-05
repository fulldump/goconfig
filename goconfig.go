package goconfig

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"
)

func Read(c interface{}) {

	if err := readWithError(c); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

}

func readWithError(c interface{}) error {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)
	filename := f.String("config", "", "-usage-")
	f.Parse(os.Args[1:])

	// Read from file JSON
	if err := FillJson(c, *filename); err != nil {
		return errors.New("Config file error: " + err.Error())
	}

	// Overwrite configuration with environment vars:
	if err := FillEnvironments(c); err != nil {
		return errors.New("Config env error: " + err.Error())
	}

	// Overwrite configuration with command line args:
	if err := FillArgs(c, os.Args[1:]); err != nil {
		return errors.New("Config arg error: " + err.Error())
	}

	return nil
}
