package validation

import (
	"fmt"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// depsOn is the validation function for validating if the field specified exists in the current struct and if it has been initialized
var depsOn = func(fl validator.FieldLevel) bool {

	if fl.Parent().Kind() != reflect.Struct {
		panic(fmt.Sprintf("Invalid parent type %T: must be a struct", fl.Parent().Interface()))
	}

	if fl.Param() == "" {
		panic(fmt.Sprintf("Invalid tag composition"))
	}

	val, _, found := fl.GetStructFieldOK()

	if found {
		switch val.Kind() {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
			return !val.IsNil()
		default:
			return true
		}
	}

	panic(fmt.Sprintf("Field not found"))

}
