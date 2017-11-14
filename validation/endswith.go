package validation

import (
	"fmt"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var endsWith = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		suf := fl.Param()
		return strings.HasSuffix(str, suf)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
