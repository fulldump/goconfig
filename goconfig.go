package goconfig

import (
	"flag"
	"reflect"
	"strings"
)

func Read(c interface{}) {

	path := []string{}

	parse_struct(c, path)

	flag.Parse()

}

func parse_struct(c interface{}, p []string) {
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
		ptr := value.Addr().Interface()
		kind := value.Kind()

		p = append(p, strings.ToLower(name))
		name_path := strings.Join(p, ".")

		if false {
			// Do nothing

		} else if reflect.Bool == kind {
			flag.BoolVar(ptr.(*bool), name_path, value.Interface().(bool), usage)

		} else if reflect.Float64 == kind {
			flag.Float64Var(ptr.(*float64), name_path, value.Interface().(float64), usage)

		} else if reflect.Int64 == kind {
			flag.Int64Var(ptr.(*int64), name_path, value.Interface().(int64), usage)

		} else if reflect.Int == kind {
			flag.IntVar(ptr.(*int), name_path, value.Interface().(int), usage)

		} else if reflect.String == kind {
			flag.StringVar(ptr.(*string), name_path, value.Interface().(string), usage)

		} else if reflect.Uint64 == kind {
			flag.Uint64Var(ptr.(*uint64), name_path, value.Interface().(uint64), usage)

		} else if reflect.Uint == kind {
			flag.UintVar(ptr.(*uint), name_path, value.Interface().(uint), usage)

		} else if reflect.Struct == kind {
			// fmt.Println("TODO: Struct")
			parse_struct(ptr, p)

		} else if reflect.Slice == kind {
			panic("Slice is not supported by goconfig at this moment.")

		} else {
			panic("Kind `" + kind.String() + "` is not supported by goconfig (field `" + name + "`)")
		}

		p = p[0 : len(p)-1]

	}

}
