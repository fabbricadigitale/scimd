package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// PathExists checks whether the given path exists or not
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// pathExists ...
var pathExists = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return PathExists(field.String())
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
