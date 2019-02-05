package goconfig

import (
	"testing"
)

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
		MyStruct    struct {
			MyItem string
		}
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
		"-mystruct.myitem", "nested",
	}

	err := FillArgs(&c, args)
	AssertNil(t, err)

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

	if c.MyStruct.MyItem != "nested" {
		t.Error("MyStruct.MyItem should be 'nested'")
	}

}

func TestFillArgsWithArrayDefinedUndefined(t *testing.T) {

	c := struct {
		Defined   []string
		Undefined []string
	}{
		Defined: []string{"default", "values"},
	}

	args := []string{
		`-defined=["one","two","three"]`,
		`-undefined=["a", "b", "c"]`,
	}

	err := FillArgs(&c, args)
	AssertNil(t, err)

	AssertEqual(t, c.Defined, []string{"one", "two", "three"})
	AssertEqual(t, c.Undefined, []string{"a", "b", "c"})

}

func TestFillArgsWithArrayPointers(t *testing.T) {

	c := struct {
		Pointers []*string
	}{}

	args := []string{
		`-pointers=["one","two","three"]`,
	}

	err := FillArgs(&c, args)
	AssertNil(t, err)

	one := "one"
	two := "two"
	three := "three"

	AssertEqual(t, c.Pointers, []*string{&one, &two, &three})

}

func TestFillArgsWithArrayScalars(t *testing.T) {

	c := struct {
		// Bool
		MyBool []bool

		// String
		MyString []string

		// Numbers
		MyFloat64 []float64
		MyInt64   []int64
		MyInt     []int
		MyUint64  []uint64
		MyUint    []uint
	}{}

	args := []string{
		// Bool
		"-mybool", "[true, false, true]",

		// String
		"-mystring", `["one", "two", "three"]`,

		// Numbers
		"-myfloat64", "[1.23, 1.24]",
		"-myint64", "[123, 124]",
		"-myint", "[8888, 9999]",
		"-myuint64", "[64, 65]",
		"-myuint", "[4444, 5555]",
	}

	err := FillArgs(&c, args)
	AssertNil(t, err)

	// Bool
	AssertEqual(t, c.MyBool, []bool{true, false, true})

	// String
	AssertEqual(t, c.MyString, []string{"one", "two", "three"})

	// Numbers
	AssertEqual(t, c.MyFloat64, []float64{1.23, 1.24})
	AssertEqual(t, c.MyInt64, []int64{123, 124})
	AssertEqual(t, c.MyInt, []int{8888, 9999})
	AssertEqual(t, c.MyUint64, []uint64{64, 65})
	AssertEqual(t, c.MyUint, []uint{4444, 5555})

}

func TestFillArgsWithArrayStructs(t *testing.T) {

	type mystruct struct {
		Name string
		Age  int
	}

	c := struct {
		MyStruct []mystruct
	}{}

	args := []string{
		"-mystruct", `[{"name":"Fulanez", "age": 33}, {"name":"Menganez", "age": 22}]`,
	}

	err := FillArgs(&c, args)
	AssertNil(t, err)

	AssertEqual(t, c.MyStruct, []mystruct{
		{Name: "Fulanez", Age: 33},
		{Name: "Menganez", Age: 22},
	})

}

func TestFillArgsWithArrayMalformed(t *testing.T) {

	c := struct {
		MyArray []string
	}{}

	args := []string{
		"-myarray", `[1,2,3]`,
	}

	err := FillArgs(&c, args)
	AssertNotNil(t, err)

	AssertEqual(t, err.Error(), "'MyArray' should be a JSON "+
		"array: json: cannot unmarshal number into Go value of type string")

}

func TestFillArgsParseFail(t *testing.T) {

	c := struct{}{}

	args := []string{
		"-help",
	}

	err := FillArgs(&c, args)
	AssertNotNil(t, err)

}
