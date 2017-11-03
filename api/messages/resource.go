package messages

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

// Common resource attributes
type Common core.Common

// A ScimError is a description of a SCIM error.
type ScimError struct {
	msg string // description of error
}

func (e *ScimError) Error() string { return e.msg }

// Resource The data resource structure
type Resource struct {
	Common
	Attributes map[string]core.Complex `json:"-"`
}

func unmarshalAttrs(schema *core.Schema, parts map[string]json.RawMessage) (*core.Complex, error) {
	attrs := core.Complex{}
	for _, aDef := range schema.Attributes {
		if part, ok := parts[aDef.Name]; ok {
			var value interface{}
			if err := json.Unmarshal(part, &value); err != nil {
				return nil, err
			}
			attrs[aDef.Name] = value
		}
	}

	return &attrs, nil
}

func getSchema(schema string, allowedSchemas []string) *core.Schema {
	repo := schemas.Repository()
	for _, s := range allowedSchemas {
		if s == schema {
			return repo.GetSchema(s)
		}
	}

	return nil
}

func (r *Resource) UnmarshalJSON(b []byte) error {

	repo := schemas.Repository()

	// Unmarshal common parts
	if err := json.Unmarshal(b, &r.Common); err != nil {
		return err
	}

	// Validate and get ResourceType
	resourceType := repo.GetResourceType(r.Common.Meta.ResourceType)
	if resourceType == nil {
		return &ScimError{"Unsupported Resource Type"}
	}

	// Validate and get schema
	baseSchema := getSchema(resourceType.Schema, r.Common.Schemas)
	if baseSchema == nil {
		return &ScimError{"Unsupported Schema"}
	}

	// Unmarshal othe parts
	var parts map[string]json.RawMessage
	if err := json.Unmarshal(b, &parts); err != nil {
		return err
	}

	var err error
	r.Attributes = make(map[string]core.Complex)

	// Grab base schema attributes
	var baseAttrs *core.Complex
	baseAttrs, err = unmarshalAttrs(baseSchema, parts)
	if err != nil {
		return err
	}
	r.Attributes[baseSchema.ID] = *baseAttrs

	// Grab extension schemas attributes
	for _, schExt := range resourceType.SchemaExtensions {

		if extRawMsg, ok := parts[schExt.Schema]; ok {

			var extParts map[string]json.RawMessage
			if err := json.Unmarshal(extRawMsg, &extParts); err != nil {
				return err
			}

			if extSchema := getSchema(schExt.Schema, r.Common.Schemas); extSchema != nil {

				var attrs *core.Complex
				attrs, err = unmarshalAttrs(extSchema, extParts)
				r.Attributes[extSchema.ID] = *attrs
			}

		}

	}

	return nil
}
