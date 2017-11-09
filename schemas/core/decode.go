package core

import (
	"encoding/json"
)

// TODO Improve func and case for add empty array
func isNullMsg(data json.RawMessage) bool {
	return len(data) == 4 && string(data[0]) == "n" && string(data[1]) == "u" && string(data[2]) == "l" && string(data[3]) == "l"
}

// Unmarshal a SCIM a complex value by attributes definition
func (attributes *Attributes) Unmarshal(data map[string]json.RawMessage) (*Complex, error) {
	ret := Complex{}
	for _, aDef := range *attributes {
		if part, ok := data[aDef.Name]; ok {
			value, err := aDef.Unmarshal(part)
			if err != nil {
				return nil, err
			}
			ret[aDef.Name] = value
		}
	}

	return &ret, nil
}

// Unmarshal a SCIM simple value by attribute definition
func (attribute *Attribute) Unmarshal(data json.RawMessage) (interface{}, error) {

	if isNullMsg(data) {
		return nil, nil
	}

	if attribute.MultiValued {
		return unmarshalMulti(attribute, data)
	}
	return unmarshalSingular(attribute, data)
}

func unmarshalSingular(attr *Attribute, data json.RawMessage) (Value, error) {

	var err error

	if attr.Type == ComplexType {
		var subParts map[string]json.RawMessage
		err = json.Unmarshal(data, &subParts)
		if err != nil {
			return nil, err
		}
		c, err := attr.SubAttributes.Unmarshal(subParts)
		return *c, err
	}

	var p DataType
	if p, err = NewDataType(attr.Type); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, p); err != nil {
		return nil, err
	}
	return p.Indirect(), nil
}

func unmarshalMulti(attr *Attribute, data json.RawMessage) (MultiValue, error) {
	var parts []json.RawMessage
	if err := json.Unmarshal(data, &parts); err != nil {
		return nil, err
	}

	ret := make(MultiValue, len(parts))

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
