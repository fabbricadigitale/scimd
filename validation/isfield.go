package validation

import (
	"fmt"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// isField is the validation function for validating if a field is equal to a value;
// both are specified in the current field's value following this pattern: "attr:value"
var isField = func(fl validator.FieldLevel) bool {

	parent := fl.Parent()
	params := strings.Split(fl.Param(), ":")

	if len(params) != 2 {
		panic(fmt.Sprintf("Bad tag composition"))
	}
	if params[0] == "" || params[1] == "" {
		panic(fmt.Sprintf("No parameters specified"))
	}
	if parent.FieldByName(params[0]).Kind() == reflect.Invalid {
		panic(fmt.Sprintf("Invalid field"))
	}

	// FieldByName panics if fl.Parent() is not a struct and returns the zero Value if no field is found
	return parent.FieldByName(params[0]).String() == params[1]

	// TODO: this validator works only with string value; it needs to extend the matching with any type of value
}
