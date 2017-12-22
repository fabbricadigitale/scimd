package validation

import (
	"fmt"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

// topStructName is the validation function for validating if a top struct's name is equal to a value
var topStructName = func(fl validator.FieldLevel) bool {

	top := fl.Top()
	param := fl.Param()

	if fl.Parent().Kind() != reflect.Struct {
		panic(fmt.Sprintf("Invalid parent type %T: must be a struct", fl.Parent().Interface()))
	}

	if param == "" {
		panic(fmt.Sprintf("Invalid param value"))
	}

	if top.Type().Name() != param {
		return false
	}

	return true
}
