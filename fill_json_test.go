package goconfig

import "testing"

func TestFillJson(t *testing.T) {

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
	}{}

	FillJson(c, "")

}
