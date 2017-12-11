package core

import "github.com/fabbricadigitale/scimd/schemas/datatype"
import "fmt"
import "reflect"

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

func enforceMulti(attribute *Attribute, data []interface{}) ([]datatype.DataTyper, error) {
	ret := make([]datatype.DataTyper, len(data))

	for i, p := range data {
		value, err := enforceSingle(attribute, p)
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
		if mv, ok := data.([]interface{}); ok {
			return enforceMulti(attribute, mv)
		}
		return nil, &IncompatibleTypeError{
			expected: "[]" + attribute.Type,
			actual:   reflect.TypeOf(data).String(),
		}
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
