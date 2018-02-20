package validation

import (
	"fmt"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var notStartsWith = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		pre := fl.Param()
		return !strings.HasPrefix(str, pre)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
