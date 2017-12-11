package resource

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// Valuer is the interface implemented by types that can hold resource's values
type Valuer interface {
	Values(ns string) *datatype.Complex
}

// Resource represents a mapped resource. It implements both core.ResourceTyper and Valuer
type Resource struct {
	core.CommonAttributes
	data map[string]*datatype.Complex
}

var _ core.ResourceTyper = (*Resource)(nil)
var _ Valuer = (*Resource)(nil)

// SetValues is the method to set Resource attributes by schema namespace
func (r *Resource) SetValues(ns string, values *datatype.Complex) {
	if r.data == nil {
		r.data = make(map[string]*datatype.Complex)
	}

	if values == nil {
		delete(r.data, ns)
	} else {
		r.data[ns] = values
	}
}

// Values is the method to access the attributes by schema namespace
func (r *Resource) Values(ns string) *datatype.Complex {
	return r.data[ns]
}

// UnmarshalJSON is the Resource Marshal implementation
func (r *Resource) UnmarshalJSON(b []byte) error {
	// Unmarshal common parts
	if err := json.Unmarshal(b, &r.CommonAttributes); err != nil {
		return err
	}

	// Validate and get ResourceType
	resourceType := r.ResourceType()
	if resourceType == nil {
		return &core.ScimError{"Unsupported Resource Type"}
	}

	// Validate and get schema
	schema := resourceType.GetSchema()
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

	exts := resourceType.GetSchemaExtensions()

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
		"id":         r.ID,
		"externalId": r.ExternalID,
		"schemas":    r.Schemas,
		"meta":       r.Meta,
	}

	// Get BaseSchema to encode core attributes
	// TODO: Generalize this code block
	// ****

	// Validate and get ResourceType
	resourceType := r.ResourceType()
	if resourceType == nil {
		return nil, &core.ScimError{"Unsupported Resource Type"}
	}
	// Validate and get schema
	schema := resourceType.GetSchema()
	if schema == nil {
		return nil, &core.ScimError{"Unsupported Schema"}
	}
	// ****

	// Marshal schema attrs to the top level
	for key, values := range *r.Values(schema.GetIdentifier()) {
		if msg, err = json.Marshal(values); err != nil {
			return nil, err
		}
		out[key] = msg
	}

	// Marshal extensions to proper namespace key
	for _, extSch := range resourceType.GetSchemaExtensions() {
		if extSch != nil {
			ns := extSch.GetIdentifier()
			values := *r.Values(ns)
			if msg, err = json.Marshal(values); err != nil {
				return nil, err
			}
			out[ns] = msg
		}

	}

	return json.Marshal(out)
}
