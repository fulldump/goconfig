package goconfig

import "reflect"

func set(dest, value interface{}) {
	dest_v := reflect.ValueOf(dest)
	dest_t := reflect.TypeOf(dest)
	value_v := reflect.ValueOf(value)

	dest_v.Elem().Set(value_v.Convert(dest_t).Elem())
}
