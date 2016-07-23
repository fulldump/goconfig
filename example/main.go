package main

import (
	"fmt"

	"github.com/fulldump/goconfig"
)

type myconfig struct {
	Name     string `The name of something`
	Logfile  string `Log file`
	MaxProcs int    `Maximum number of procs`
}

func main() {

	c := &myconfig{
		Name:    "default name",
		Logfile: "defaultlogfile.log",
	}

	goconfig.Read(c)

	fmt.Printf("%#v\n", c)
}
