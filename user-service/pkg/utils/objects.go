package utils

import (
	"reflect"
)

func CopyFields(source, destination interface{}) error {
	srcValue := reflect.ValueOf(source).Elem()
	destValue := reflect.ValueOf(destination).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.FieldByName(srcValue.Type().Field(i).Name)

		if destField.IsValid() && destField.CanSet() && srcField.Type() == destField.Type() {
			destField.Set(srcField)
		}
	}

	return nil
}
