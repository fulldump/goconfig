<img src="logo.png">

<p align="center">
<a href="https://travis-ci.org/fulldump/goconfig"><img src="https://travis-ci.org/fulldump/goconfig.svg?branch=master"></a>
<a href="https://goreportcard.com/report/fulldump/goconfig"><img src="http://goreportcard.com/badge/fulldump/goconfig"></a>
<a href="https://godoc.org/github.com/fulldump/goconfig"><img src="https://godoc.org/github.com/fulldump/goconfig?status.svg" alt="GoDoc"></a>
</p>

Goconfig is an extremely simple configuration library for your Go programs.
Make your configuration flags compact and easy to read.

Arguments are parsed from command line with the standard `flag` library.

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [How to use](#how-to-use)
- [Supported types](#supported-types)
- [Builtin flags](#builtin-flags)
  - [-help](#-help)
  - [-config](#-config)
- [Contribute](#contribute)
- [Testing](#testing)
- [Example project](#example-project)

<!-- /MarkdownTOC -->


# How to use

Define your structure with **descriptions**:

```go
type myconfig struct {
	Name      string `usage:"The name of something"`
	EnableLog bool   `usage:"Enable logging into logdb" json:"enable_log"`
	MaxProcs  int    `usage:"Maximum number of procs"`
	UsersDB   db
	LogDB     db
}

type db struct {
	Host string `usage:"Host where db is located"`
	User string `usage:"Database user"`
	Pass string `usage:"Database password"`
}
```

Instance your config with **default values**:

```go
c := &myconfig{
	EnableLog: true,
	UsersDB: db{
		Host: "localhost",
		User: "root",
		Pass: "123456",
	},
}
```

**Fill** your config:
```go
goconfig.Read(c)
```

How the `-help` looks like:

```
Usage of example:
  -enablelog
    	Enable logging into logdb (default true)
  -logdb.host string
    	Host where db is located (default "localhost")
  -logdb.pass string
    	Database password (default "123456")
  -logdb.user string
    	Database user (default "root")
  -maxprocs int
    	Maximum number of procs
  -name string
    	The name of something
  -usersdb.host string
    	Host where db is located
  -usersdb.pass string
    	Database password
  -usersdb.user string
    	Database user
```


# Supported types

Mainly almost all types from `flag` library are supported:

* bool
* float64
* int64
* int
* string
* uint64
* uint
* struct (hyerarchical keys)

For the moment `duration` type is not supported.

Type `slice` or `array` is also being considered.


# Builtin flags

## -help

Since `flag` library is using the key `-help` to show usage, Goconf is behaving
in the same way.

## -config

Builtin flag `-config` allow read configuration from a file. For the example
configuration above, this is a sample config.json file:

```json
{
  "name": "Fulanito",
  "usersdb": {
    "host": "localhost",
    "user": "admin",
    "pass": "123"
  }
}
```

Configuration precedence is as follows (higher to lower):
* Arg command line
* Json config file
* Default value


# Contribute

Feel free to fork, make changes and pull-request to master branch.

If you prefer, [create a new issue](https://github.com/fulldump/goconfig/releases/new)
or email me for new features, issues or whatever.


# Testing

This command will pass all tests.

No tests are expected for the moment.

```sh
make
```


# Example project

This project includes a sample project with a sample configuration. To make the binary:

```sh
make example
```
