package request

import (
	"reflect"
	"errors"
	)

var (
	IncompleteRequestError error = errors.New("incomplete request")
)

func IsIncomplete(req interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(req))

	for i := 0; i < v.NumField(); i++ {
		// Skip field if not required
		if v.Type().Field(i).Tag.Get("required") != "true" {
			continue
		}

		// If required field is empty return true
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return true
			}
		case reflect.Int64:
			if field.Int() == 0 {
				return true
			}
		case reflect.Float64:
			if field.Float() == 0 {
				return true
			}
		}
	}

	return false
}