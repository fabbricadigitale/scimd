package validation

import (
	"fmt"
	"reflect"

	"github.com/fabbricadigitale/scimd/schemas"
	validator "gopkg.in/go-playground/validator.v9"
)

var attrName = func(fl validator.FieldLevel) bool {
	field := fl.Field()
	parent := fl.Parent()

	if parent.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Invalid parent type %T: must be a struct", parent.Interface()))
	}

	typeField := parent.FieldByName("Type")

	if typeField == (reflect.Value{}) {
		panic(fmt.Sprintf("Field Type not found in the Struct"))
	}

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if typeField.String() == "reference" {
			return str == schemas.ReferenceAttrName
		}
		return schemas.AttrNameRegexp.MatchString(str)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
