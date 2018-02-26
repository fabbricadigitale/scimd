package validation

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

var hasSubDirectory = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		dir := field.String()
		sub := fl.Param()
		if sub == "" {
			// When no parameter is given check whether it contains almost one subdirectory
			files, _ := ioutil.ReadDir("./")
			for _, f := range files {
				if f.IsDir() {
					return true
				}
			}

			return false
		}

		return isDir(filepath.Join(dir, sub))
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
