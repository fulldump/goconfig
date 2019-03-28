package goconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func FillEnvironments(c interface{}) (err error) {

	traverse(c, func(i item) {
		env := strings.ToUpper(strings.Join(i.Path, "_"))
		value := os.Getenv(env)

		if "" == value {
			return
		}

		if reflect.TypeOf(time.Duration(0)) == i.Value.Type() {
			if d, err := unmarshalDurationString(value); err == nil {
				v := int64(d)
				set(i.Ptr, &v)
			}

		} else if reflect.Bool == i.Kind {
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

		} else if reflect.Slice == i.Kind {
			jsonErr := json.Unmarshal([]byte(value), i.Ptr)
			if jsonErr != nil {
				err = errors.New(fmt.Sprintf(
					"'%s' should be a JSON array: %s",
					env, jsonErr.Error(),
				))
			}

		}

	})

	return
}
