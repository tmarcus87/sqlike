package statement

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrorMustBeAPointer   = errors.New("must be a pointer")
	ErrorMustBeANonNilPtr = errors.New("must be a non-nil pointer")
)

func getOrderedColumnName(value interface{}) ([]string, error) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	res := make([]string, 0)
	for i := 0; i < v.Type().NumField(); i++ {
		f := v.Type().Field(i)
		if tag, ok := f.Tag.Lookup("sqlike"); ok {
			res = append(res, strings.ToLower(tag))
		} else {
			res = append(res, strings.ToLower(f.Name))
		}
	}
	return res, nil
}

func getColumnName2FieldValueMap(value interface{}) (map[string]reflect.Value, error) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	name2value := make(map[string]reflect.Value)
	for i := 0; i < v.Type().NumField(); i++ {
		f := v.Type().Field(i)
		if tag, ok := f.Tag.Lookup("sqlike"); ok {
			name2value[strings.ToLower(tag)] = v.Field(i)
		} else {
			name2value[strings.ToLower(f.Name)] = v.Field(i)
		}
	}

	return name2value, nil
}
