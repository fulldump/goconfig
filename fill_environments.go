package goconfig

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func FillEnvironments(c interface{}) {
	traverse(c, func(i item) {
		env := strings.ToUpper(strings.Join(i.Path, "_"))
		value := os.Getenv(env)

		if "" == value {
			return
		}

		if reflect.Bool == i.Kind {
			if v, err := strconv.ParseBool(value); nil == err {
				set(i.Ptr, &v)
			}

		} else if reflect.Float64 == i.Kind {
			if v, err := strconv.ParseFloat(value, 64); nil == err {
				set(i.Ptr, &v)
			}

		} else if reflect.Float32 == i.Kind {
			if v, err := strconv.ParseFloat(value, 32); nil == err {
				w := float32(v)
				set(i.Ptr, &w)
			}

		} else if reflect.Int64 == i.Kind {
			if v, err := strconv.ParseInt(value, 10, 64); nil == err {
				set(i.Ptr, &v)
			}

		} else if reflect.Int32 == i.Kind {
			if v, err := strconv.ParseInt(value, 10, 64); nil == err {
				w := int32(v)
				set(i.Ptr, &w)
			}

		} else if reflect.Int == i.Kind {
			if v, err := strconv.ParseInt(value, 10, strconv.IntSize); nil == err {
				w := int(v)
				set(i.Ptr, &w)
			}

		} else if reflect.String == i.Kind {
			set(i.Ptr, &value)

		} else if reflect.Uint64 == i.Kind {
			if v, err := strconv.ParseUint(value, 10, 64); nil == err {
				set(i.Ptr, &v)
			}

		} else if reflect.Uint32 == i.Kind {
			if v, err := strconv.ParseUint(value, 10, 32); nil == err {
				w := uint32(v)
				set(i.Ptr, &w)
			}

		} else if reflect.Uint == i.Kind {
			if v, err := strconv.ParseUint(value, 10, strconv.IntSize); nil == err {
				w := uint(v)
				set(i.Ptr, &w)
			}

		}

	})
}
