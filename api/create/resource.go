package create

import (
	"time"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/version"
	"github.com/satori/go.uuid"
)

// Resource creates a new res of type resType and stores it into s.
//
// This func expects that res was populated according to the given resType
// (ie. type validation is not performed here).
//
// Commons' attributes, if present, will be ignored and overwritten
// (with the only exception of ExternalID that if populated will be used).
//
// Attributes whose mutability is "readOnly" will be ignored and removed,
// other attributes can be written while creation.
//
// If a required attributes is missing an error is returned, with exception of
// "readOnly" ones whose are just ignored.
//
// Finally, this func expects that uniqueness conflict are handled by storage.Storer
func Resource(s storage.Storer, resType *core.ResourceType, res *resource.Resource) (core.ResourceTyper, error) {

	// (todo) Could client assert "schemas"?
	// If so, we must validate asserted schemas against available schema on current resType.
	// Currently, not yet supported.
	// Note also that resType cannot be asserted by client.
	res.Schemas = make([]string, len(resType.SchemaExtensions)+1)
	res.Schemas[0] = resType.Schema
	for i, ext := range resType.SchemaExtensions {
		res.Schemas[i+1] = ext.Schema
	}

	// ID, ResourceType and other stuff that can be only asserted by the service provider
	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	res.ID = ID.String()

	now := time.Now()
	res.Meta = core.Meta{
		ResourceType: resType.GetIdentifier(),
		Created:      &now,
		LastModified: &now,
		Version:      version.GenerateVersion(true, ID.String(), now.String()),
	}

	// Since the ResourceType was set, we can check required
	if err := attr.CheckRequired(res); err != nil {
		return nil, err
	}

	// Attributes whose mutability is "readOnly" SHALL be ignored
	ro, err := attr.Paths(resType, func(attribute *core.Attribute) bool {
		return attribute.Mutability == schemas.MutabilityReadOnly
	})
	if err != nil {
		return nil, err
	}

	for _, p := range ro {
		p.Context(resType).Delete(res)
	}

	// Assume s.Create will return an error in case of uniquiness conflict
	if err := s.Create(res); err != nil {
		return nil, err
	}

	return query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
}
