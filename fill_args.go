package goconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

var values = map[string]interface{}{}

type postFillArgs struct {
	item
	Raw *string
}

func FillArgs(c interface{}, args []string) error {
	var f = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.Usage = func() {}
	f.SetOutput(os.Stdout)

	// Default config flag
	f.String("config", "", "Configuration JSON file")

	post := []postFillArgs{}

	traverse(c, func(i item) {
		name_path := strings.ToLower(strings.Join(i.Path, "."))
		env_name := strings.ToUpper(strings.Join(i.Path, "_"))

		usage := i.Usage
		if usage != "" {
			usage += " "
		}
		usage += "[env " + env_name + "]"

		if reflect.TypeOf(time.Duration(0)) == i.Value.Type() {
			value := ""
			f.StringVar(&value, name_path, i.Value.Interface().(time.Duration).String(), usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
			})

		} else if reflect.Bool == i.Kind {
			f.BoolVar(i.Ptr.(*bool), name_path, i.Value.Interface().(bool), usage)

		} else if reflect.Float64 == i.Kind {
			f.Float64Var(i.Ptr.(*float64), name_path, i.Value.Interface().(float64), usage)

		} else if reflect.Int64 == i.Kind {
			f.Int64Var(i.Ptr.(*int64), name_path, i.Value.Interface().(int64), usage)

		} else if reflect.Int == i.Kind {
			f.IntVar(i.Ptr.(*int), name_path, i.Value.Interface().(int), usage)

		} else if reflect.String == i.Kind {
			f.StringVar(i.Ptr.(*string), name_path, i.Value.Interface().(string), usage)

		} else if reflect.Uint64 == i.Kind {
			f.Uint64Var(i.Ptr.(*uint64), name_path, i.Value.Interface().(uint64), usage)

		} else if reflect.Uint == i.Kind {
			f.UintVar(i.Ptr.(*uint), name_path, i.Value.Interface().(uint), usage)

		} else if reflect.Slice == i.Kind {

			b, _ := json.Marshal(i.Value.Interface())

			value := ""
			f.StringVar(&value, name_path, string(b), usage)

			post = append(post, postFillArgs{
				Raw:  &value,
				item: i,
			})

		} else {
			panic("Kind `" + i.Kind.String() +
				"` is not supported by goconfig (field `" + i.FieldName + "`)")
		}

	})

	if err := f.Parse(args); err != nil && err == flag.ErrHelp {
		m := bytes.NewBufferString("Usage of goconfig:\n\n")
		f.SetOutput(m)
		f.PrintDefaults()
		return errors.New(m.String())
	}

	// Postprocess flags: unsupported flags needs to be declared as string
	// and parsed later. Here is the place.
	for _, p := range post {
		// Special case for durations
		if reflect.TypeOf(time.Duration(0)) == p.Value.Type() {
			d, err := unmarshalDurationString(*p.Raw)
			if err != nil {
				return fmt.Errorf(
					"'%s' should be nanoseconds or a time.Duration string: %s",
					p.FieldName, err.Error(),
				)
			}
			p.Value.SetInt(int64(d))

			continue
		}

		err := json.Unmarshal([]byte(*p.Raw), p.Ptr)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"'%s' should be a JSON array: %s",
				p.FieldName, err.Error(),
			))
		}
	}

	return nil
}
