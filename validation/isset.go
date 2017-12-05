package validation

import (
	"fmt"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

var depsOn = func(fl validator.FieldLevel) bool {

	if fl.Parent().Kind() != reflect.Struct {
		panic(fmt.Sprintf("Invalid parent type %T: must be a struct", fl.Parent().Interface()))
	}

	_, _, ok := fl.GetStructFieldOK()

	return ok

}
