package goconfig

import (
	"flag"
	"os"
	"reflect"
	"strings"
	"fmt"
)

var values = map[string]interface{}{}

func FillArgs(c interface{}, args []string) {
	var f = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.Usage = func() {}
	f.SetOutput(os.Stdout)

	// Default config flag
	f.String("config", "", "Configuration JSON file")

	traverse(c, func(i item) {
		name_path := strings.ToLower(strings.Join(i.Path, "."))

		if reflect.Bool == i.Kind {
			f.BoolVar(i.Ptr.(*bool), name_path, i.Value.Interface().(bool), i.Usage)

		} else if reflect.Float64 == i.Kind {
			f.Float64Var(i.Ptr.(*float64), name_path, i.Value.Interface().(float64), i.Usage)

		} else if reflect.Int64 == i.Kind {
			f.Int64Var(i.Ptr.(*int64), name_path, i.Value.Interface().(int64), i.Usage)

		} else if reflect.Int == i.Kind {
			f.IntVar(i.Ptr.(*int), name_path, i.Value.Interface().(int), i.Usage)

		} else if reflect.String == i.Kind {
			f.StringVar(i.Ptr.(*string), name_path, i.Value.Interface().(string), i.Usage)

		} else if reflect.Uint64 == i.Kind {
			f.Uint64Var(i.Ptr.(*uint64), name_path, i.Value.Interface().(uint64), i.Usage)

		} else if reflect.Uint == i.Kind {
			f.UintVar(i.Ptr.(*uint), name_path, i.Value.Interface().(uint), i.Usage)

		} else if reflect.Slice == i.Kind {
			panic("Slice is not supported by goconfig at this moment.")

		} else {
			panic("Kind `" + i.Kind.String() +
				"` is not supported by goconfig (field `" + i.FieldName + "`)")
		}

	})

	if err := f.Parse(args); err != nil && err == flag.ErrHelp {
		f.SetOutput(os.Stderr)
		fmt.Fprint(f.Output(), "Usage of goconfig:\n\n")
		f.PrintDefaults()
		os.Exit(1)
	}
}
