<img src="logo.png">

Goconfig is an extremely simple configuration library for your Go programs.

Arguments are parsed from command line with the standard `flag` library.

IMPORTANT NOTE: This is work in progress.

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [How to use](#how-to-use)
- [Contribute](#contribute)
	- [Testing](#testing)

<!-- /MarkdownTOC -->

# How to use

Define your structure with **descriptions**:

```go
type myconfig struct {
	Name     string `The name of something`
	Logfile  string `Log file`
	MaxProcs int    `Maximum number of procs`
}
```

Instance your config with **default values**:

```go
c := &myconfig{
	Name:     "default name",
	Logfile:  "defaultlogfile.log",
	MaxProcs: 8,
}
```

**Fill** your config up:
```go
goconfig.Read(c)
```

# Contribute

TODO this

## Testing

This command will pass all tests

```sh
make
```
