package validation

import (
	"fmt"
	"os"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

func isDir(path string) bool {
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

// isDirectory checks whether the given path is a directory or not
var isDirectory = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return isDir(field.String())
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
