package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// pathExists ...
var pathExists = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if _, err := os.Stat(str); os.IsNotExist(err) {
			return false
		}
		return true
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
