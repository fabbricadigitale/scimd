package resource

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

// Common resource attributes
type Common core.Common

// Resource The data resource structure
type Resource struct {
	Common
	Attributes map[string]core.Complex `json:"-"`
}

func getSchema(schema string, allowedSchemas []string) *core.Schema {
	repo := schemas.GetSchemaRepository()
	for _, s := range allowedSchemas {
		if s == schema {
			return repo.Get(s)
		}
	}

	return nil
}

func (r *Resource) UnmarshalJSON(b []byte) error {
	repo := schemas.GetResourceTypeRepository()

	// Unmarshal common parts
	if err := json.Unmarshal(b, &r.Common); err != nil {
		return err
	}

	// Validate and get ResourceType
	resourceType := repo.Get(r.Common.Meta.ResourceType)
	if resourceType == nil {
		return &core.ScimError{"Unsupported Resource Type"}
	}

	// Validate and get schema
	baseSchema := getSchema(resourceType.Schema, r.Common.Schemas)
	if baseSchema == nil {
		return &core.ScimError{"Unsupported Schema"}
	}

	// Unmarshal other parts
	var parts map[string]json.RawMessage
	if err := json.Unmarshal(b, &parts); err != nil {
		return err
	}

	var err error
	r.Attributes = make(map[string]core.Complex)

	// Grab base schema attributes
	var baseAttrs *core.Complex
	baseAttrs, err = baseSchema.Attributes.Unmarshal(parts)

	if err != nil {
		return err
	}
	r.Attributes[baseSchema.GetIdentifier()] = *baseAttrs

	// Grab extension schemas attributes
	for _, schExt := range resourceType.SchemaExtensions {

		if extRawMsg, ok := parts[schExt.Schema]; ok {

			var extParts map[string]json.RawMessage
			if err := json.Unmarshal(extRawMsg, &extParts); err != nil {
				return err
			}

			if extSchema := getSchema(schExt.Schema, r.Common.Schemas); extSchema != nil {

				var attrs *core.Complex
				attrs, err = extSchema.Attributes.Unmarshal(extParts)
				if err != nil {
					return err
				}
				r.Attributes[extSchema.GetIdentifier()] = *attrs
			}

		}

	}

	return nil
}

// MarshalJSON is the Resource Marshal implementation
func (r *Resource) MarshalJSON() ([]byte, error) {

	var msg json.RawMessage
	var err error

	// Attach Common attribute to the map before marshal operation
	// TODO: implement "omitempty" check
	out := map[string]interface{}{
		"id":         r.Common.ID,
		"externalId": r.Common.ExternalID,
		"schemas":    r.Common.Schemas,
		"meta":       r.Common.Meta,
	}

	// Get BaseSchema to encode core attributes
	// TODO: Generalize this code block
	// ****
	repo := schemas.GetResourceTypeRepository()

	// Validate and get ResourceType
	resourceType := repo.Get(r.Common.Meta.ResourceType)
	if resourceType == nil {
		return nil, &core.ScimError{"Unsupported Resource Type"}
	}
	// Validate and get schema
	baseSchema := getSchema(resourceType.Schema, r.Common.Schemas)
	// ****

	// Bring to the above level core attributes
	for key, value := range r.Attributes[baseSchema.GetIdentifier()] {

		if msg, err = json.Marshal(value); err != nil {
			return nil, err
		}
		out[key] = msg
	}

	for schema, attrs := range r.Attributes {
		if schema == baseSchema.GetIdentifier() {
			continue
		}
		if msg, err = json.Marshal(attrs); err != nil {
			return nil, err
		}
		out[schema] = msg
	}

	return json.Marshal(out)
}
