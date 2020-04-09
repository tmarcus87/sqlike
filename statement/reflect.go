package statement

import (
	"errors"
	"github.com/iancoleman/strcase"
	"reflect"
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
		res = append(res, getColumnName(f))
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
		name2value[getColumnName(f)] = v.Field(i)
	}

	return name2value, nil
}

func getColumnName(f reflect.StructField) string {
	if tag, ok := f.Tag.Lookup("sqlike"); ok {
		return tag
	}
	return strcase.ToSnake(f.Name)
}
