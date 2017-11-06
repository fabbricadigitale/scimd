package core

import (
	"encoding/json"
	"time"
)

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
	if attribute.MultiValued {
		return unmarshalMulti(attribute, data)
	}
	return unmarshalSingular(attribute, data)
}

func unmarshalSingular(attr *Attribute, data json.RawMessage) (interface{}, error) {

	var err error
	var ret interface{}
	switch attr.Type {
	case "string", "reference":
		var s string
		err = json.Unmarshal(data, &s)
		ret = s
	case "boolean":
		var b bool
		err = json.Unmarshal(data, &b)
		ret = b
	case "decimal":
		var f float64
		err = json.Unmarshal(data, &f)
		ret = f
	case "integer":
		var i int64
		err = json.Unmarshal(data, &i)
		ret = i
	case "dateTime":
		var t time.Time
		err = json.Unmarshal(data, &t)
		ret = t
	case "binary":
		var b string
		err = json.Unmarshal(data, &b)
		ret = []byte(b)

	case "complex":
		c := Complex{}
		var subParts map[string]json.RawMessage
		err = json.Unmarshal(data, &subParts)
		if err != nil {
			return nil, err
		}
		for _, subDef := range attr.SubAttributes {
			if data, ok := subParts[subDef.Name]; ok {
				value, err := subDef.Unmarshal(data)
				if err != nil {
					return nil, err
				}
				c[subDef.Name] = value
			}
		}
		ret = c

	default:
		return nil, &SchemaError{"Invalid type"}

	}

	if err != nil {
		return nil, err
	}

	return ret, nil

}

func unmarshalMulti(attr *Attribute, data json.RawMessage) ([]interface{}, error) {
	var parts []json.RawMessage
	if err := json.Unmarshal(data, &parts); err != nil {
		return nil, err
	}

	ret := make([]interface{}, len(parts))

	for i, p := range parts {
		value, err := unmarshalSingular(attr, p)
		if err != nil {
			return nil, err
		}
		ret[i] = value
	}

	return ret, nil
}
