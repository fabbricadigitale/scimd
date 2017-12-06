package validation

import (
	"fmt"
	"reflect"

	u "github.com/leodido/go-urn"
	validator "gopkg.in/go-playground/validator.v9"
)

var urn = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		_, match := u.Parse(str)

		return match
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
