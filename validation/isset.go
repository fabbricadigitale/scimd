package validation

import (
	validator "gopkg.in/go-playground/validator.v9"
)

var depsOn = func(fl validator.FieldLevel) bool {
	_, _, ok := fl.GetStructFieldOK()

	return ok

}
