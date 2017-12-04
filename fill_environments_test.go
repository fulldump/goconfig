package goconfig

import (
	"os"
	"testing"
)

func TestFillEnvironments(t *testing.T) {

	filled_string := "filled pointer"

	c := struct {
		MyBoolTrue  bool
		MyBoolFalse bool
		MyString    string
		MyFloat64   float64
		MyFloat32   float32
		MyInt64     int64
		MyInt32     int32
		MyInt       int
		MyUint64    uint64
		MyUint32    uint32
		MyUint      uint
		MyEmpty     string
		MyPointer   *string
		MyStruct    struct {
			MyItem string
		}
	}{
		MyPointer: &filled_string,
	}

	os.Setenv("MYBOOLTRUE", "true")
	os.Setenv("MYBOOLFALSE", "false")
	os.Setenv("MYSTRING", "Hello world")
	os.Setenv("MYFLOAT64", "1.23")
	os.Setenv("MYFLOAT32", "3.21")
	os.Setenv("MYINT64", "123")
	os.Setenv("MYINT32", "321")
	os.Setenv("MYINT", "8888")
	os.Setenv("MYUINT64", "64")
	os.Setenv("MYUINT32", "32")
	os.Setenv("MYUINT", "4444")
	os.Setenv("MYSTRUCT_MYITEM", "nested")
	os.Setenv("MYEMPTY", "")
	os.Setenv("MYPOINTER", "replaced pointer")

	FillEnvironments(&c)

	if c.MyBoolTrue != true {
		t.Error("MyBoolTrue should be true")
	}

	if c.MyBoolFalse != false {
		t.Error("MyBoolFalse should be false")
	}

	if c.MyString != "Hello world" {
		t.Error("MyString should be 'Hello World'")
	}

	if c.MyFloat64 != 1.23 {
		t.Error("MyFloat64 should be 1.23")
	}

	if c.MyFloat32 != 3.21 {
		t.Error("MyFloat32 should be 3.21")
	}

	if c.MyInt64 != 123 {
		t.Error("MyInt64 should be 123")
	}

	if c.MyInt32 != 321 {
		t.Error("MyInt32 should be 321")
	}

	if c.MyInt != 8888 {
		t.Error("MyInt should be 8888")
	}

	if c.MyUint64 != 64 {
		t.Error("MyUint64 should be 64")
	}

	if c.MyUint32 != 32 {
		t.Error("MyUint32 should be 32")
	}

	if c.MyUint != 4444 {
		t.Error("MyUint should be 4444")
	}

	if c.MyEmpty != "" {
		t.Error("MyEmpty should be ''")
	}

	if c.MyStruct.MyItem != "nested" {
		t.Error("MyStruct.MyItem should be 'nested'")
	}

	if *c.MyPointer != "filled pointer" {
		t.Error("MyPointer do not match")
	}

}
