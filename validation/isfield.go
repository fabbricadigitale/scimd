package validation

import (
	"fmt"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// isField is the validation function for validating if a field is equal to a value;
// both are specified in the current field's value following this pattern: "attr:value"
var isField = func(fl validator.FieldLevel) bool {

	params := strings.Split(fl.Param(), ":")

	if params != nil && len(params) != 2 {
		panic(fmt.Sprintf("Bad tag composition"))
	}

	// FieldByName panics if fl.Parent() is not a struct and returns the zero Value if no field is found
	return fl.Parent().FieldByName(params[0]).String() == params[1]

}
