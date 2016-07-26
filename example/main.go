package main

import (
	"fmt"

	"github.com/fulldump/goconfig"
)

// My custom configuration options
type myconfig struct {
	Name      string `The name of something`
	EnableLog bool   `Enable logging into logdb`
	MaxProcs  int    `Maximum number of procs`
	UsersDB   db
	LogDB     db
}

// Reusable configuration structure
type db struct {
	Host string `Host where db is located`
	User string `Database user`
	Pass string `Database password`
}

func main() {

	// Default configuration
	c := &myconfig{
		EnableLog: true,
		LogDB: db{
			Host: "localhost",
			User: "root",
			Pass: "123456",
		},
	}

	goconfig.Read(c)

	fmt.Printf("%#v\n", c)
}
