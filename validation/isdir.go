package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// isDirectory ...
var isDirectory = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		file, err := os.Open(str)
		if err != nil {
			return false
		}
		defer file.Close()
		info, err := file.Stat()
		if err != nil {
			return false
		}
		return info.IsDir()
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
