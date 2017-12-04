package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fabbricadigitale/scimd/schemas"
	validator "gopkg.in/go-playground/validator.v9"
)

var urn = func(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if strings.HasPrefix(str, schemas.InvalidURNPrefix) {
			return false
		}
		return schemas.URIRegexp.MatchString(str)
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
