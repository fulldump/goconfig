package goconfig

import (
	"os"
	"testing"
	"time"
)

func TestFillEnvironments(t *testing.T) {
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
		MyStruct    struct {
			MyItem string
		}
		MyDurationNano   time.Duration
		MyDurationString time.Duration
	}{}

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
	os.Setenv("MYDURATIONNANO", "15000000000")
	os.Setenv("MYDURATIONSTRING", "15s")

	err := FillEnvironments(&c)
	AssertNil(t, err)

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

	if c.MyDurationNano.String() != "15s" {
		t.Error("MyDurationNano should be '15s'")
	}

	if c.MyDurationString.String() != "15s" {
		t.Error("MyDurationString should be '15s'")
	}

}

func TestFillEnvironmentsWithArray(t *testing.T) {

	c := struct {
		MyStringArray []string
	}{
		[]string{"four", "five", "six"},
	}

	os.Setenv("MYSTRINGARRAY", `["one", "two", "three"]`)

	err := FillEnvironments(&c)
	AssertNil(t, err)

	AssertEqual(t, c.MyStringArray, []string{"one", "two", "three"})

}

func TestFillEnvironmentsWithArrayMalformed(t *testing.T) {

	c := struct {
		MyStringArray []string
	}{}

	os.Setenv("MYSTRINGARRAY", `}`)

	err := FillEnvironments(&c)
	AssertNotNil(t, err)

	AssertEqual(t, err.Error(), "'MYSTRINGARRAY' should be a JSON"+
		" array: invalid character '}' looking for beginning of value")

}

func TestFillEnvironmentsEmptyOverride(t *testing.T) {
	c := struct {
		Value string
	}{Value: "default"}

	os.Setenv("VALUE", "")

	err := FillEnvironments(&c)
	AssertNil(t, err)

	AssertEqual(t, c.Value, "")
}
