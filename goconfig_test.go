package goconfig

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
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

func TestFillJsonAnonymousStructs(t *testing.T) {
	f, err := ioutil.TempFile("", "test-nested-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(`{"Name":"Fulanez", "timeout":20000000, "server":{"Other": {"name":"Gonzo", "timeout":"3s"} }}`); err != nil {
		t.Fatal(err)
	}

	type AnonymousStruct struct {
		Name    string
		Timeout time.Duration
	}

	cfg := struct {
		AnonymousStruct
		Server struct {
			Port  int
			Other AnonymousStruct
		}
	}{}

	err = FillJson(&cfg, f.Name())
	AssertNil(t, err)
	if cfg.Name != "Fulanez" {
		t.Errorf("expected 'Fulanez', got '%s'", cfg.Name)
	}
	if cfg.Timeout.String() != "20ms" {
		t.Errorf("expected '20ms', got '%s'", cfg.Timeout.String())
	}
	if cfg.Server.Other.Name != "Gonzo" {
		t.Errorf("expected 'Gonzo', got '%s'", cfg.Server.Other.Name)
	}
	if cfg.Server.Other.Timeout.String() != "3s" {
		t.Errorf("expected '3s', got '%s'", cfg.Server.Other.Timeout.String())
	}
}
