package resource

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

// Common resource attributes
type Common core.Common

// Resource The data resource structure
type Resource struct {
	Common
	data map[string]*core.Complex
}

// SetValues is the method to set Resource attributes
func (r *Resource) SetValues(schema string, values *core.Complex) {
	r.data[schema] = values
}

// GetValues is the method to access the attributes
func (r *Resource) GetValues(schema string) *core.Complex {
	return r.data[schema]
}

func getSchema(schema string, allowedSchemas []string) *core.Schema {
	repo := core.GetSchemaRepository()
	for _, s := range allowedSchemas {
		if s == schema {
			return repo.Get(s)
		}
	}

	return nil
}

// UnmarshalJSON is the Resource Marshal implementation
func (r *Resource) UnmarshalJSON(b []byte) error {
	repo := core.GetResourceTypeRepository()

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
	r.data = make(map[string]*core.Complex)

	// Grab base schema attributes
	var baseAttrs *core.Complex
	baseAttrs, err = baseSchema.Attributes.Unmarshal(parts)

	if err != nil {
		return err
	}
	r.SetValues(baseSchema.GetIdentifier(), baseAttrs)

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
				r.SetValues(extSchema.GetIdentifier(), attrs)
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
	repo := core.GetResourceTypeRepository()

	// Validate and get ResourceType
	resourceType := repo.Get(r.Common.Meta.ResourceType)
	if resourceType == nil {
		return nil, &core.ScimError{"Unsupported Resource Type"}
	}
	// Validate and get schema
	baseSchema := getSchema(resourceType.Schema, r.Common.Schemas)
	// ****

	// Bring to the above level core attributes
	for key, value := range *r.GetValues(baseSchema.GetIdentifier()) {

		if msg, err = json.Marshal(value); err != nil {
			return nil, err
		}
		out[key] = msg
	}

	for _, extSch := range r.Common.Schemas {

		if extSch == baseSchema.GetIdentifier() {
			continue
		}

		attrs := *r.GetValues(extSch)
		if msg, err = json.Marshal(attrs); err != nil {
			return nil, err
		}
		out[extSch] = msg
	}

	return json.Marshal(out)
}
