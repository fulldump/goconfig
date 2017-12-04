package goconfig

import (
	"reflect"
	"strconv"
	"strings"
)

type callback func(i item)

type item struct {
	FieldName string
	Usage     string
	Ptr       interface{}
	Kind      reflect.Kind
	Path      []string
	Value     reflect.Value
}

func traverse(c interface{}, f callback) {
	traverse_recursive(c, f, []string{})
}

func traverse_recursive(c interface{}, f callback, p []string) {

	t := reflect.ValueOf(c)

	// Follow pointers
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		name := field.Name
		value := t.Field(i)
		usage := field.Tag.Get("usage")
		ptr := value.Addr().Interface()
		kind := value.Kind()

		p = append(p, strings.ToLower(name))
		name_path := strings.Join(p, ".")

		if reflect.Struct == kind {
			traverse_recursive(ptr, f, p)
		} else if reflect.Slice == kind {
			traverse_recursive_slice(ptr, f, p)
		} else {
			f(item{
				FieldName: name,
				Usage:     usage,
				Ptr:       ptr,
				Kind:      kind,
				Path:      p,
				Value:     value,
			})
		}

		values[name_path] = ptr

		p = p[0 : len(p)-1]

	}

}

func traverse_recursive_slice(c interface{}, f callback, p []string) {

	value := reflect.ValueOf(c)

	// Follow pointers
	for reflect.Ptr == value.Kind() {
		value = value.Elem()
	}

	for v := 0; v < value.Len(); v++ {
		name := strconv.Itoa(v)
		value := value.Index(v)
		ptr := value.Addr().Interface()
		path := append(p, name)
		f(item{
			FieldName: name,
			Usage:     "",
			Ptr:       ptr,
			Kind:      value.Kind(),
			Path:      path,
			Value:     value,
		})
		name_path := strings.Join(path, ".")
		values[name_path] = ptr
	}
	//panic("Slice is not supported by goconfig at this moment.")

}
