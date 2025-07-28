package goconfig

import (
	"reflect"
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
	Tags      reflect.StructTag
	Anonymous bool
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

		pr := p // parents to pass recursively
		if !field.Anonymous {
			pr = append(p, strings.ToLower(name))
		}
		name_path := strings.Join(p, ".")

		if reflect.Struct == kind {
			traverse_recursive(ptr, f, pr)

		} else if reflect.Slice == kind {
			//panic("Slice is not supported by goconfig at this moment.")
			f(item{
				FieldName: name,
				Usage:     usage,
				Ptr:       ptr,
				Kind:      kind,
				Path:      pr,
				Value:     value,
				Anonymous: field.Anonymous,
			})

			traverse_array(ptr, f, pr)
		} else {
			f(item{
				FieldName: name,
				Usage:     usage,
				Ptr:       ptr,
				Kind:      kind,
				Path:      pr,
				Value:     value,
				Anonymous: field.Anonymous,
			})

		}

		values[name_path] = ptr

		//p = p[0 : len(p)-1]

	}

}

func traverse_array(c interface{}, f callback, p []string) {

}

func traverse_json(c interface{}, f callback) {

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
		tags := field.Tag

		f(item{
			FieldName: name,
			Usage:     usage,
			Ptr:       ptr,
			Kind:      kind,
			Value:     value,
			Tags:      tags,
			Anonymous: field.Anonymous,
		})
	}
}
