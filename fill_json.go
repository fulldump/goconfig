package goconfig

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"
	"time"
)

func FillJson(c interface{}, filename string) error {
	if "" == filename {
		return nil
	}

	if c == nil {
		return errors.New("config target cannot be nil")
	}

	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if !reflect.TypeOf(c).Implements(unmarshalerType) {
		if err := validateConfigTarget(c); err != nil {
			return err
		}
	}

	data, err := os.ReadFile(filename)
	if nil != err {
		return err
	}

	return unmarshalJSON(data, c)
}

func unmarshalJSON(data []byte, c interface{}) error {
	if c == nil {
		return errors.New("config target cannot be nil")
	}

	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

	if reflect.TypeOf(c).Implements(unmarshalerType) {
		if err := json.Unmarshal(data, c); err != nil {
			return errors.New("Bad json file: " + err.Error())
		}

	} else {
		if err := validateConfigTarget(c); err != nil {
			return err
		}

		var values map[string]json.RawMessage
		if err := json.Unmarshal(data, &values); err != nil {
			return errors.New("Bad json file: " + err.Error())
		}
		for k, v := range values {
			k = strings.ToLower(k)
			values[k] = v
		}

		traverse_json(c, func(i item) {
			tag := i.Tags.Get("json")
			if len(tag) > 0 {
				if i := strings.Index(tag, ","); i != -1 {
					tag = tag[:i]
				}

				if tag == "-" {
					return
				}
			}

			// If the field is an anonymous struct without tag,
			// treat its fields as part of the current level
			if i.Anonymous && tag == "" && (i.Kind == reflect.Struct || (i.Kind == reflect.Ptr && i.Value.Type().Elem().Kind() == reflect.Struct)) {
				if i.Kind == reflect.Ptr {
					if i.Value.IsNil() {
						i.Value.Set(reflect.New(i.Value.Type().Elem()))
					}
					unmarshalJSON(data, i.Value.Interface())
					return
				}

				unmarshalJSON(data, i.Ptr)
				return
			}

			var value json.RawMessage
			if v, ok := values[tag]; ok {
				value = v
			} else if v, ok := values[i.FieldName]; ok {
				value = v
			} else if v, ok := values[strings.ToLower(i.FieldName)]; ok {
				value = v
			} else {
				return
			}

			if reflect.PtrTo(i.Value.Type()).Implements(unmarshalerType) {
				json.Unmarshal(value, i.Ptr)

			} else if i.Value.Kind() == reflect.Struct {
				unmarshalJSON(value, i.Ptr)

			} else if i.Value.Kind() == reflect.Ptr && i.Value.Type().Elem().Kind() == reflect.Struct {
				if i.Value.IsNil() {
					i.Value.Set(reflect.New(i.Value.Type().Elem()))
				}
				unmarshalJSON(value, i.Value.Interface())

			} else if reflect.TypeOf(time.Duration(0)) == i.Value.Type() {
				var d time.Duration
				// try nanosecond int, then duration string
				if err := json.Unmarshal(value, &d); err != nil {
					tmp := ""
					if err := json.Unmarshal(value, &tmp); err != nil {
						return
					}

					if d, err = unmarshalDurationString(tmp); err != nil {
						return
					}
				}

				v := int64(d)
				set(i.Ptr, &v)

			} else {
				json.Unmarshal(value, i.Ptr)
			}
		})
	}

	return nil
}
