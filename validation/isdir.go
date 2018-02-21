package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// IsDir checks whether the given path is a directory or not
func IsDir(path string) bool {
	file, err := os.Open(path)
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

// isDirectory ...
var isDirectory = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return IsDir(field.String())
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
