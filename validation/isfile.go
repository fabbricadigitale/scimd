package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

var re = regexp.MustCompile(`^\.?(.*)`)

var isFile = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		pre := fl.Param()

		file, err := os.Open(str)
		if err != nil {
			return false
		}
		defer file.Close()
		info, err := file.Stat()
		if err != nil || info.IsDir() {
			return false
		}

		// Eventually check also the file extension
		if pre != "" {
			ext := filepath.Ext(info.Name())

			return ext == re.ReplaceAllString(pre, `.$1`)
		}

		return true
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
