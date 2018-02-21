package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// pathExists checks whether the given path exists or not
var pathExists = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		if _, err := os.Stat(field.String()); os.IsNotExist(err) {
			return false
		}
		return true
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
