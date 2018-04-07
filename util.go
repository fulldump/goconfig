package goconfig

import (
	"reflect"
	"os"
)

func set(dest, value interface{}) {
	dest_v := reflect.ValueOf(dest)
	dest_t := reflect.TypeOf(dest)
	value_v := reflect.ValueOf(value)
	dest_v.Elem().Set(value_v.Convert(dest_t).Elem())
}

//isFile returns true when the provided path belongs to a file
func isFile(pth string) bool {
	fi, err := os.Stat(pth)
	return err == nil && !fi.IsDir()
}



