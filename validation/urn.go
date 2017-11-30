package validation

import (
	"fmt"
	"reflect"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas"
	validator "gopkg.in/go-playground/validator.v9"
)

var urn = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if attr.Parse(str).Valid() != true {
			panic(fmt.Sprintf("Invalid URN composition: %+v", field.Interface()))
		}
		return schemas.URIRegexp.MatchString(str)
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
