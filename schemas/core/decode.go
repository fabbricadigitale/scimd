package core

import (
	"encoding/json"
	"strings"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// ByName returns the *Attribute with the given name, performing a insensitive match. It returns nil if no attribute was found.
func (attributes Attributes) ByName(name string) *Attribute {
	name = strings.ToLower(name)
	for _, a := range attributes {
		if name == strings.ToLower(a.Name) {
			return a
		}
	}
	return nil
}

func byKeyInsensitive(key string, data map[string]json.RawMessage) *json.RawMessage {
	if part, ok := data[key]; ok {
		return &part
	}
	key = strings.ToLower(key)
	for k, v := range data {
		if key == strings.ToLower(k) {
			return &v
		}
	}
	return nil
}

// Unmarshal a SCIM a complex value by attributes definition
func (attributes *Attributes) Unmarshal(data map[string]json.RawMessage) (*datatype.Complex, error) {
	ret := datatype.Complex{}
	for _, a := range *attributes {
		if part := byKeyInsensitive(a.Name, data); part != nil {
			value, err := a.Unmarshal(*part)
			if err != nil {
				return &ret, err
			}
			ret[a.Name] = value
		}
	}

	return &ret, nil
}

// Unmarshal a SCIM simple value by attribute definition
func (attribute *Attribute) Unmarshal(data json.RawMessage) (interface{}, error) {

	if attribute.MultiValued {
		return unmarshalMulti(attribute, data)
	}

	return unmarshalSingular(attribute, data)
}

func unmarshalSingular(attr *Attribute, data json.RawMessage) (datatype.DataTyper, error) {

	if len(data) == 4 && strings.ToLower(string(data)) == "null" {
		return nil, nil
	}

	var err error

	if attr.Type == datatype.ComplexType {
		var subParts map[string]json.RawMessage
		if err = json.Unmarshal(data, &subParts); err != nil {
			return nil, err
		}
		c, err := attr.SubAttributes.Unmarshal(subParts)
		if c != nil {
			return *c, err
		}
		return nil, err
	}

	var p datatype.DataTyper
	if p, err = datatype.New(attr.Type); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, p); err != nil {
		return nil, err
	}
	return p.Value(), nil
}

func unmarshalMulti(attr *Attribute, data json.RawMessage) ([]datatype.DataTyper, error) {
	var parts []json.RawMessage
	if err := json.Unmarshal(data, &parts); err != nil {
		return nil, err
	}

	ret := make([]datatype.DataTyper, len(parts))

	for i, p := range parts {
		value, err := unmarshalSingular(attr, p)
		if err != nil {
			return nil, err
		}
		ret[i] = value
	}

	return ret, nil
}

// Unmarshal SCIM values by schema definition
func (schema *Schema) Unmarshal(data map[string]json.RawMessage) (*datatype.Complex, error) {
	return schema.Attributes.Unmarshal(data)
}
