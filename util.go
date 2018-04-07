package goconfig

import (
	"reflect"
	"os"
	"encoding/json"
	"fmt"
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


//fillStruct conversion function
func fillStruct(m map[string]interface{}, s interface{}) {
	j, err := json.Marshal(m)
	if err!=nil{
		fmt.Println(err)
	}
	json.Unmarshal(j, s)
}

func normalizeMap(in map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for i, variable := range in {
		switch val := variable.(type) {
		case map[interface{}]interface{}:
			res[i] = normalizeMap(normalizeGenericMap(val))
		case *map[string]interface{}:
			res[i] = normalizeMap(*val)
		default:
			res[i] = val
		}
	}
	return res
}

func normalizeGenericMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, valChild := range in {
		res[k.(string)] = valChild
	}
	return res
}