package goconfig

import (
	"errors"
	"reflect"
)

func validateConfigTarget(c interface{}) error {
	if c == nil {
		return errors.New("config target cannot be nil")
	}

	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr {
		return errors.New("config target must be a pointer")
	}

	if v.IsNil() {
		return errors.New("config target cannot be nil")
	}

	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("config target must point to a struct")
	}

	return nil
}
