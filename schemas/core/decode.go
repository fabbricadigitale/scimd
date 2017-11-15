package core

import (
	"encoding/json"
	"strings"
)

// Unmarshal a SCIM a complex value by attributes definition
func (attributes *Attributes) Unmarshal(data map[string]json.RawMessage) (*Complex, error) {
	ret := Complex{}
	for _, aDef := range *attributes {
		if part, ok := data[aDef.Name]; ok {
			value, err := aDef.Unmarshal(part)
			if err != nil {
				return &ret, err
			}
			ret[aDef.Name] = value
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

func unmarshalSingular(attr *Attribute, data json.RawMessage) (DataType, error) {

	if len(data) == 4 && strings.ToLower(string(data)) == "null" {
		return nil, nil
	}

	var err error

	if attr.Type == ComplexType {
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

	var p DataType
	if p, err = NewDataType(attr.Type); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, p); err != nil {
		return nil, err
	}
	return p.Value(), nil
}

func unmarshalMulti(attr *Attribute, data json.RawMessage) ([]DataType, error) {
	var parts []json.RawMessage
	if err := json.Unmarshal(data, &parts); err != nil {
		return nil, err
	}

	ret := make([]DataType, len(parts))

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
func (schema *Schema) Unmarshal(data map[string]json.RawMessage) (*Complex, error) {
	return schema.Attributes.Unmarshal(data)
}
