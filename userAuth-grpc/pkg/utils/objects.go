package utils

import (
	"fmt"
	"reflect"
)

func CopyFields(source, destination interface{}) error {
	srcValue := reflect.ValueOf(source).Elem()
	destValue := reflect.ValueOf(destination).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.Field(i)

		if srcField.Type() == destField.Type() && destField.CanSet() {
			destField.Set(srcField)
		} else {
			return fmt.Errorf("field mismatch or unexported field: %s", srcValue.Type().Field(i).Name)
		}
	}

	return nil
}
