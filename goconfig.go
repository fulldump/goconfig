package goconfig

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

func Read(c interface{}) {

	parse_struct(c)

	flag.Parse()

	fmt.Println("FU")

}

func parse_struct(c interface{}) {
	t := reflect.ValueOf(c)

	// Follow pointers
	for reflect.Ptr == t.Kind() {
		t = t.Elem()
	}

	n := t.NumField()
	for i := 0; i < n; i++ {
		field := t.Type().Field(i)
		name := field.Name
		value := t.Field(i)
		usage := string(field.Tag)

		zz := value.Addr().Interface()

		flag.StringVar(zz.(*string), strings.ToLower(name), value.Interface().(string), usage)
	}

}
