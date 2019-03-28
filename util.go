package goconfig

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func set(dest, value interface{}) {
	dest_v := reflect.ValueOf(dest)
	dest_t := reflect.TypeOf(dest)
	value_v := reflect.ValueOf(value)

	dest_v.Elem().Set(value_v.Convert(dest_t).Elem())
}

func unmarshalDurationString(s string) (time.Duration, error) {
	// nanoseconds stored as a string
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Duration(i), nil
	}

	// duration string
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	return 0, fmt.Errorf("invalid time.Duration format, use a duration string (like '15s') or nanoseconds")
}
