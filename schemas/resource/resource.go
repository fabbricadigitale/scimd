package resource

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// Resource The data resource structure
type Resource struct {
	core.Common
	data map[string]*datatype.Complex
}

// SetValues is the method to set Resource attributes
func (r *Resource) SetValues(ns string, values *datatype.Complex) {
	r.data[ns] = values
}

// GetValues is the method to access the attributes
func (r *Resource) GetValues(ns string) *datatype.Complex {
	return r.data[ns]
}

// UnmarshalJSON is the Resource Marshal implementation
func (r *Resource) UnmarshalJSON(b []byte) error {
	// Unmarshal common parts
	if err := json.Unmarshal(b, &r.Common); err != nil {
		return err
	}

	// Validate and get ResourceType
	resourceType := r.GetResourceType()
	if resourceType == nil {
		return &core.ScimError{"Unsupported Resource Type"}
	}

	// Validate and get schema
	schema := r.GetSchema()
	if schema == nil {
		return &core.ScimError{"Unsupported Schema"}
	}

	// Unmarshal other parts
	var parts map[string]json.RawMessage
	if err := json.Unmarshal(b, &parts); err != nil {
		return err
	}

	var err error
	r.data = make(map[string]*datatype.Complex)

	// Get schema attributes' values
	var values *datatype.Complex
	if values, err = schema.Unmarshal(parts); err != nil {
		return err
	}
	r.SetValues(schema.GetIdentifier(), values)

	exts := r.GetSchemaExtensions()

	for _, schExt := range exts {
		if extRawMsg, ok := parts[schExt.GetIdentifier()]; ok && schExt != nil {
			var extParts map[string]json.RawMessage
			if err := json.Unmarshal(extRawMsg, &extParts); err != nil {
				return err
			}
			var values *datatype.Complex
			if values, err = schExt.Unmarshal(extParts); err != nil {
				return err
			}
			r.SetValues(schExt.GetIdentifier(), values)
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

	// Validate and get ResourceType
	resourceType := r.GetResourceType()
	if resourceType == nil {
		return nil, &core.ScimError{"Unsupported Resource Type"}
	}
	// Validate and get schema
	schema := r.GetSchema()
	if schema == nil {
		return nil, &core.ScimError{"Unsupported Schema"}
	}
	// ****

	// Marshal schema attrs to the top level
	for key, values := range *r.GetValues(schema.GetIdentifier()) {
		if msg, err = json.Marshal(values); err != nil {
			return nil, err
		}
		out[key] = msg
	}

	// Marshal extensions to proper namespace key
	for _, extSch := range r.GetSchemaExtensions() {
		if extSch != nil {
			ns := extSch.GetIdentifier()
			values := *r.GetValues(ns)
			if msg, err = json.Marshal(values); err != nil {
				return nil, err
			}
			out[ns] = msg
		}

	}

	return json.Marshal(out)
}
