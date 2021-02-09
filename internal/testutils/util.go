package testutils

import "reflect"

func Ptr(val interface{}) interface{} {
	p := reflect.New(reflect.TypeOf(val))
	p.Elem().Set(reflect.ValueOf(val))
	return p.Interface()
}
