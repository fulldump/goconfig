package goconfig

import "testing"

func TestFillArgs(t *testing.T) {
	c := struct {
		MyBoolTrue  bool
		MyBoolFalse bool
		MyString    string
		MyFloat64   float64
		MyInt64     int64
		MyInt       int
		MyUint64    uint64
		MyUint      uint
	}{}

	args := []string{
		"-mybooltrue",
		"-myboolfalse=false",
		"-mystring", "HelloWorld",
		"-myfloat64", "1.23",
		"-myint64", "123",
		"-myint", "8888",
		"-myuint64", "64",
		"-myuint", "4444",
	}

	FillArgs(&c, args)

	if c.MyBoolTrue != true {
		t.Error("MyBoolTrue should be true")
	}

	if c.MyBoolFalse != false {
		t.Error("MyBoolFalse should be false")
	}

	if c.MyString != "HelloWorld" {
		t.Error("MyString should be 'HelloWorld'")
	}

	if c.MyFloat64 != 1.23 {
		t.Error("MyFloat64 should be 1.23")
	}

	if c.MyInt64 != 123 {
		t.Error("MyInt64 should be 123")
	}

	if c.MyInt != 8888 {
		t.Error("MyInt should be 8888")
	}

	if c.MyUint64 != 64 {
		t.Error("MyUint64 should be 64")
	}

	if c.MyUint != 4444 {
		t.Error("MyUint should be 4444")
	}

}
