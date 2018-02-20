package validation

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// Validator is a singleton instance providing validation functionalities
var Validator *validator.Validate

func init() {
	Validator = validator.New()
	Validator.RegisterValidation("startswith", startsWith)
	Validator.RegisterValidation("nstartswith", notStartsWith)
	Validator.RegisterValidation("endswith", endsWith)
	Validator.RegisterValidation("urn", urn)
	Validator.RegisterValidation("attrname", attrName)
	Validator.RegisterValidation("attrpath", attrPath)
	Validator.RegisterValidation("isfield", isField)
	Validator.RegisterValidation("isset", depsOn)
	Validator.RegisterValidation("uniqueattr", uniqueAttr)
}
