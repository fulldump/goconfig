package goconfig

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

const test_fill_json = `{
	"MyBoolTrue": true,
	"MyBoolFalse": false,
	"MyString": "HelloWorld",
	"MyFloat64": 1.23,
	"MyFloat32": 1.23,
	"MyInt64": 123,
	"MyInt32": 123,
	"MyInt": 8888,
	"MyUint64": 64,
	"MyUint32": 32,
	"MyUint": 4444,
	"MyEmpty": "",
	"MyStruct": {
		"MyItem": "nested"
	},
	"MyDuration": 15000000000,
	"MyDurationString": "15s",
	"MyDurationNanoString": "15000000000",
	"my_tag": "tag",
	"myalternatecase": "lower",
	"MyArray": [
		1,
		2
	],
	"MyUnmarshaler": "unmarshal"
}`

func TestFillJson(t *testing.T) {
	f, err := ioutil.TempFile("", "test-fill-json-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err = f.WriteString(test_fill_json); err != nil {
		t.Fatal(err)
	}

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
		MyDuration           time.Duration
		MyDurationString     time.Duration
		MyDurationNanoString time.Duration
		MyTag                string `json:"my_tag"`
		MyAlternateCase      string
		MyArray              []int
		MyUnmarshaler        testJsonUnmarshaler
	}{}

	err = FillJson(&c, f.Name())
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

	if c.MyDuration.String() != "15s" {
		t.Error("MyDuration should be 15s")
	}

	if c.MyDurationString.String() != "15s" {
		t.Error("MyDurationString should be 15s")
	}

	if c.MyDurationNanoString.String() != "15s" {
		t.Error("MyDurationNanoString should be 15s")
	}

	if c.MyTag != "tag" {
		t.Error("MyTag should be 'tag'")
	}

	if c.MyAlternateCase != "lower" {
		t.Error("MyAlternateCase should be 'lower'")
	}

	if !reflect.DeepEqual(c.MyArray, []int{1, 2}) {
		t.Error("MyArray should be '[1 2]'")
	}

	if c.MyUnmarshaler != "unmarshal" {
		t.Error("MyUnmarshaler should be 'unmarshal'")
	}
}

func TestFillJson_UnexistingFilename(t *testing.T) {

	c := struct{}{}

	err := FillJson(&c, "/")
	AssertNotNil(t, err)

}

type testJsonUnmarshaler string

func (c *testJsonUnmarshaler) UnmarshalJSON(data []byte) error {
	tmp := ""
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = testJsonUnmarshaler(tmp)
	return nil
}

func TestFillJson_AnonymousStruct(t *testing.T) {
	f, err := ioutil.TempFile("", "test-fill-json-anon-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err = f.WriteString(`{"Name":"Fulanez"}`); err != nil {
		t.Fatal(err)
	}

	type AnonymousStruct struct {
		Name string
	}

	cfg := struct {
		AnonymousStruct
	}{}

	err = FillJson(&cfg, f.Name())
	AssertNil(t, err)

	if cfg.Name != "Fulanez" {
		t.Error("Name should be 'Fulanez'")
	}
}
