package core

import (
	"fmt"
	"reflect"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// A IncompatibleTypeError is a SCIM error describing when a value is not compatible with attribute type.
type IncompatibleTypeError struct {
	expected string
	actual   string
}

func (e *IncompatibleTypeError) Error() string {
	return fmt.Sprintf("Expected type '%s', but found '%s'", e.expected, e.actual)
}

func enforceSingle(attribute *Attribute, data interface{}) (datatype.DataTyper, error) {

	v, err := datatype.Cast(data, attribute.Type)
	if v == nil || err != nil {
		return nil, err
	}

	if v.Type() == datatype.ComplexType {
		return attribute.SubAttributes.Enforce(v.(datatype.Complex))
	}

	return v, nil

}

func enforceMulti(attribute *Attribute, data interface{}) ([]datatype.DataTyper, error) {
	var l int
	var v reflect.Value

	t := reflect.TypeOf(data)

	switch t.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		v = reflect.ValueOf(data)
		l = v.Len()
		// go ahead
	default:
		return nil, &IncompatibleTypeError{
			expected: "[]" + attribute.Type,
			actual:   t.String(),
		}
	}

	ret := make([]datatype.DataTyper, l)
	for i := 0; i < l; i++ {
		value, err := enforceSingle(attribute, v.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		ret[i] = value
	}
	return ret, nil
}

// Enforce a SCIM simple value by attribute definition
func (attribute *Attribute) Enforce(data interface{}) (interface{}, error) {

	if attribute.MultiValued {
		return enforceMulti(attribute, data)
	}

	return enforceSingle(attribute, data)

}

// Enforce SCIM Types on data, using attributes definition.
func (attributes *Attributes) Enforce(data map[string]interface{}) (*datatype.Complex, error) {

	// nil is a Null value
	if data == nil {
		return nil, nil
	}

	ret := datatype.Complex{}
	for _, a := range *attributes {
		k := a.Name
		if v, ok := data[k]; ok {
			ret[k], _ = a.Enforce(v)
		}
		// keys of a map not set are Unassigned
	}

	return &ret, nil
}

// Enforce SCIM Types on data, using schema definition.
func (schema *Schema) Enforce(data map[string]interface{}) (*datatype.Complex, error) {
	return schema.Attributes.Enforce(data)
}
