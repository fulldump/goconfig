package goconfig

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {

	c := struct {
		Value string
	}{}

	os.Setenv("VALUE", "env")

	Read(&c)

	AssertEqual(t, c.Value, "env")

}

func TestReadWithError(t *testing.T) {

	c := struct {
		Value string
	}{}

	os.Setenv("VALUE", "env")

	err := readWithError(&c)
	AssertNil(t, err)

	AssertEqual(t, c.Value, "env")
}

func TestReadWithError_MalformedEnv(t *testing.T) {

	c := struct {
		Value []string
	}{}

	os.Setenv("VALUE", "}")

	err := readWithError(&c)
	AssertNotNil(t, err)

	AssertEqual(t, err.Error(), "Config env error: "+
		"'VALUE' should be a JSON array: invalid character '}' looking for "+
		"beginning of value")
}

func TestReadWithError_MalformedArg(t *testing.T) {

	c := struct {
		Value []string
	}{}

	os.Unsetenv("VALUE")
	os.Args = []string{"cmd", "-value", "}"}

	err := readWithError(&c)
	AssertNotNil(t, err)

	AssertEqual(t, err.Error(), "Config arg error: "+
		"'Value' should be a JSON array: invalid character '}' looking for "+
		"beginning of value")
}

func TestReadWithError_FileError(t *testing.T) {

	c := struct{}{}

	os.Args = []string{"cmd", "-config", "/"}

	err := readWithError(&c)
	AssertNotNil(t, err)

	AssertEqual(t, err.Error(), "Config file error: "+
		"read /: is a directory")
}
