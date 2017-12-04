package goconfig

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFillJson(t *testing.T) {

	type testFillJson struct {
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
		MyPointer *string
	}

	c := &testFillJson{}

	j := `{
	"MyBoolTrue": true,
	"MyBoolFalse": false,
	"MyString": "Hello world",
	"MyFloat64": 1.23,
	"MyFloat32": 3.21,
	"MyInt64": 123,
	"MyInt32": 321,
	"MyInt": 8888,
	"MyUint64": 64,
	"MyUint32": 32,
	"MyUint": 4444,
	"MyStruct": {
		"MyItem": "nested"
	},
	"MyEmpty": "",
	"MyPointer": "replaced pointer"
}`

	filled_pointer := "replaced pointer"
	expected := &testFillJson{
		MyBoolTrue:  true,
		MyBoolFalse: false,
		MyString:    "Hello world",
		MyFloat64:   1.23,
		MyFloat32:   3.21,
		MyInt64:     123,
		MyInt32:     321,
		MyInt:       8888,
		MyUint64:    64,
		MyUint32:    32,
		MyUint:      4444,
		MyStruct: struct {
			MyItem string
		}{
			MyItem: "nested",
		},
		MyEmpty:   "",
		MyPointer: &filled_pointer,
	}

	filename := "TestFillJson.testcase.json"

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if nil != err {
		t.Skip("Can not write sample json:", err)
		return
	}

	f.WriteString(j)
	f.Sync()
	f.Close()

	FillJson(c, filename)

	if !reflect.DeepEqual(c, expected) {
		t.Error("Expected result does not match\n", expected, "\n", c)
		return
	}

	fmt.Println(c)
}
