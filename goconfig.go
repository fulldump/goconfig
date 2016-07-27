package goconfig

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

var values = map[string]interface{}{}

func Read(c interface{}) {

	path := []string{}

	// Add builtin -config flag
	filename := ""
	flag.StringVar(&filename, "config", filename, "Configuration JSON file")

	parse_struct(c, path)
	flag.Parse()

	// Put default values
	flag.VisitAll(func(f *flag.Flag) {
		if v, e := values[f.Name]; e {
			Set(v, f.Value)
		}
	})

	// Read from file JSON
	read_json(c, filename)

	// Overwrite configuration with command line args:
	flag.Visit(func(f *flag.Flag) {
		if v, e := values[f.Name]; e {
			Set(v, f.Value)
		}
	})

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
			flag.Bool(name_path, value.Interface().(bool), usage)

		} else if reflect.Float64 == kind {
			flag.Float64(name_path, value.Interface().(float64), usage)

		} else if reflect.Int64 == kind {
			flag.Int64(name_path, value.Interface().(int64), usage)

		} else if reflect.Int == kind {
			flag.Int(name_path, value.Interface().(int), usage)

		} else if reflect.String == kind {
			flag.String(name_path, value.Interface().(string), usage)

		} else if reflect.Uint64 == kind {
			flag.Uint64(name_path, value.Interface().(uint64), usage)

		} else if reflect.Uint == kind {
			flag.Uint(name_path, value.Interface().(uint), usage)

		} else if reflect.Struct == kind {
			parse_struct(ptr, p)

		} else if reflect.Slice == kind {
			panic("Slice is not supported by goconfig at this moment.")

		} else {
			panic("Kind `" + kind.String() +
				"` is not supported by goconfig (field `" + name + "`)")
		}

		values[name_path] = ptr

		p = p[0 : len(p)-1]

	}

}

func read_json(c interface{}, filename string) {
	if "" == filename {
		return
	}

	data, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Println("Unable to read config file `" + filename + "`!")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &c)
	if nil != err {
		fmt.Println("Config file should be a valid JSON")
		os.Exit(1)
	}

	// f, err := os.Open(filename)
	// if nil != err {
	// 	fmt.Println("Unable to read config file `" + filename + "`!")
	// 	os.Exit(1)
	// }
	// json.NewDecoder(f).Decode(&c)
}

func Set(dest, value interface{}) {
	dest_v := reflect.ValueOf(dest)
	dest_t := reflect.TypeOf(dest)
	value_v := reflect.ValueOf(value)

	dest_v.Elem().Set(value_v.Convert(dest_t).Elem())
}
