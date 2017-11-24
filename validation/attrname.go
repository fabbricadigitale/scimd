package validation

import (
	"fmt"
	"reflect"
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// AttrNameExpr is the source text used to compile a regular expression macching a SCIM attribute name
const AttrNameExpr = `[A-Za-z][\-$_0-9A-Za-z]*`

// AttrNameRegexp is the compiled Regexp built from AttrNameExpr
var AttrNameRegexp = regexp.MustCompile("^" + AttrNameExpr + "$")

var attrName = func(fl validator.FieldLevel) bool {
	field := fl.Field()
	parent := fl.Parent()

	if parent.Kind() != reflect.Struct {
		panic(fmt.Sprintf("It has to be used in a Struct"))
	}

	typeField := parent.FieldByName("Type")

	if typeField == (reflect.Value{}) {
		panic(fmt.Sprintf("No field Type"))
	}

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if typeField.String() == "reference" {
			return str == "$ref"
		}
		return AttrNameRegexp.MatchString(str)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
