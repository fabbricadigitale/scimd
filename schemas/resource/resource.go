package resource

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// Valuer is the interface implemented by types that can hold resource's values
type Valuer interface {
	Values(ns string) *datatype.Complex
	SetValues(ns string, values *datatype.Complex)
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
	if r.data == nil {
		return nil
	}

	if nsdata, ok := r.data[ns]; ok {
		return nsdata
	}

	return nil
}

// UnmarshalJSON is the Resource unmarshal implementation
func (r *Resource) UnmarshalJSON(b []byte) error {
	// Unmarshal common parts
	if err := json.Unmarshal(b, &r.CommonAttributes); err != nil {
		return err
	}

	// Validate and get ResourceType
	resourceType := r.ResourceType()
	if resourceType == nil {
		return &core.ScimError{Msg: "Unsupported Resource Type"}
	}

	// Validate and get schema
	schema := resourceType.GetSchema()
	if schema == nil {
		return &core.ScimError{Msg: "Unsupported Schema"}
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

// MarshalJSON is the Resource marshal implementation
func (r *Resource) MarshalJSON() ([]byte, error) {
	var msg json.RawMessage
	var err error

	// Attach Common attribute to the map before marshal operation
	out := map[string]interface{}{
		"id":      r.ID,
		"schemas": r.Schemas,
		"meta":    r.Meta,
	}
	if r.ExternalID != "" {
		out["externalId"] = r.ExternalID
	}

	resourceType := r.ResourceType()
	if resourceType == nil {
		return nil, &core.ScimError{Msg: "Unsupported Resource Type"}
	}
	schema := resourceType.GetSchema()
	if schema == nil {
		return nil, &core.ScimError{Msg: "Unsupported Schema"}
	}

	// Marshal schema attrs to the top level
	ns := schema.GetIdentifier()
	if values := r.Values(ns); values != nil {
		if msg, err = json.Marshal(&values); err != nil {
			return nil, err
		}
		out[ns] = msg
	}

	// Marshal extensions to proper namespace key
	for _, extSch := range resourceType.GetSchemaExtensions() {
		if extSch != nil {
			ns := extSch.GetIdentifier()
			if values := r.Values(ns); values != nil {
				if msg, err = json.Marshal(&values); err != nil {
					return nil, err
				}
				out[ns] = msg
			}
		}

	}

	return json.Marshal(out)
}
