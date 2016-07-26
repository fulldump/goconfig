<img src="logo.png">

Goconfig is an extremely simple configuration library for your Go programs.
Make your configuration flags compact and easy to read.

Arguments are parsed from command line with the standard `flag` library.

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [How to use](#how-to-use)
- [Supported types](#supported-types)
- [Builtin configuration keys](#builtin-configuration-keys)
- [Contribute](#contribute)
- [Testing](#testing)
- [Example project](#example-project)

<!-- /MarkdownTOC -->


# How to use

Define your structure with **descriptions**:

```go
type myconfig struct {
	Name      string `The name of something`
	EnableLog bool   `Enable logging into logdb`
	MaxProcs  int    `Maximum number of procs`
	UsersDB   db
	LogDB     db
}

type db struct {
	Host string `Host where db is located`
	User string `Database user`
	Pass string `Database password`
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


# Builtin configuration keys

Since `flag` library is using the key `-help` to show usage, Goconf is behaving
in the same way.


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
