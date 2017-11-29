package validation

import (
	"fmt"
	"reflect"

	"github.com/fabbricadigitale/scimd/api/attr"

	validator "gopkg.in/go-playground/validator.v9"
)

var attrPath = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		return attr.Parse(str).Valid()
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
