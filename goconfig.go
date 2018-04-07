package goconfig

import (
	"flag"
	"io/ioutil"
	"os"
	"fmt"
)

func Read(c interface{}) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)
	filename := f.String("config", "", "-usage-")
	f.Parse(os.Args[1:])

	if len(*filename) > 0 {
		p, err := getProvider(*filename)
		if err != nil {
			fmt.Printf("It fails silentlty %s ", err.Error())
		} else {
			p.fill(c)
		}
	}
	// Overwrite configuration with environment vars:
	FillEnvironments(c)

	// Overwrite configuration with command line args:
	FillArgs(c, os.Args[1:])

}
