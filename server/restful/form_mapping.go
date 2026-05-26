package restful

import (
	"reflect"
)

func mapForm(ptr interface{}, form map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// if "form" tag is nil, we inspect if the field is a struct.
// this would not make sense for JSON parsing but it does for a form
// since data is flatten

func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setBoolField(val string, field reflect.Value) error { _ = "STUB: not implemented"; return nil }

func setFloatField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}
