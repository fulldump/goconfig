package goconfig

import (
	"flag"
	"io/ioutil"
	"os"
)

func Read(c interface{}) {

	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)
	filename := f.String("config", "", "-usage-")
	f.Parse(os.Args[1:])

	// Read from file JSON
		FillJson(c, *filename)

	// Overwrite configuration with environment vars:
	FillEnvironments(c)

	// Overwrite configuration with command line args:
	FillArgs(c, os.Args[1:])

}
