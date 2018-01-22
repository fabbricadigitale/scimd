package validation

import (
	"strings"
	"reflect"
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

// The uniqueField validator checks if a field in a slice has unique values, in a case-insensitive way, and works with strings only.

var uniqueAttr = func(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.Slice, reflect.Array:

		innerType := field.Type().Elem()
		if innerType.Kind() == reflect.Ptr {
			innerType = innerType.Elem()
		}
		_, ok := innerType.FieldByName(param)
		if !ok {
			fmt.Printf("'%v' it's not an existing attribute in the struct.\n", param)
			return false
		}

		m := make(map[string]struct{})
		for i := 0; i < field.Len(); i++ {
			iV := field.Index(i)
			
			if iV.Kind() == reflect.Ptr {
				iV = iV.Elem()
			}

			if iV.FieldByName(param).Kind() != reflect.String {
				fmt.Printf("Field is %v, validator can be used on strings only attributes.\n", iV.FieldByName(param).Kind())
				return false
			}

			m[strings.ToLower(iV.FieldByName(param).String())] = struct{}{}
		}
		return field.Len() == len(m)
	default:
		panic(fmt.Sprintf("Can't be used on %T, only on a Slice", field.Interface()))
	}
}
