package goconfig

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

func FillJson(c interface{}, filename string) error {

	if "" == filename {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if nil != err {
		return err
	}

	return unmarshalJSON(data, c)
}

func unmarshalJSON(data []byte, c interface{}) error {
	if reflect.TypeOf(c).Implements(reflect.TypeOf(new(json.Unmarshaler)).Elem()) {
		if err := json.Unmarshal(data, c); err != nil {
			return errors.New("Bad json file: " + err.Error())
		}

	} else {
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
			}

			// If the field is an anonymous struct without tag,
			// treat its fields as part of the current level
			if i.Anonymous && tag == "" && (i.Kind == reflect.Struct || (i.Kind == reflect.Ptr && i.Value.Type().Elem().Kind() == reflect.Struct)) {
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

			unmarshaler := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

			if reflect.PtrTo(i.Value.Type()).Implements(unmarshaler) {
				json.Unmarshal(value, i.Ptr)

			} else if i.Value.Kind() == reflect.Struct {
				unmarshalJSON(value, i.Ptr)

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
