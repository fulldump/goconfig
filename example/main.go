package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fulldump/goconfig"
)

// My custom configuration options
type myconfig struct {
	Name      string `usage:"The name of something"`
	EnableLog bool   `usage:"Enable logging into logdb"`
	MaxProcs  int    `usage:"Maximum number of procs"`
	UsersDB   db
	LogDB     db
}

// Reusable configuration structure
type db struct {
	Host    string        `usage:"Host where db is located"`
	User    string        `usage:"Database user"`
	Pass    string        `usage:"Database password"`
	Timeout time.Duration `usage:"Timeout duration"`
}

func main() {
	// Default configuration
	c := &myconfig{
		EnableLog: true,
		LogDB: db{
			Host:    "localhost",
			User:    "root",
			Pass:    "123456",
			Timeout: 5 * time.Second,
		},
	}

	if err := goconfig.Load(c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", c)
}
