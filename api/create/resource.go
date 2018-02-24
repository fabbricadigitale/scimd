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
// This func expects that res was populated according to the given resType.
// Commons' attributes, if present, will be ignored and overwritten
// (with the only exception of ExternalID that if populated will be used).
// Attributes whose mutability is "readOnly" will be ignored and removed.
func Resource(s storage.Storer, resType *core.ResourceType, res *resource.Resource) (ret core.ResourceTyper, err error) {

	// Make a new UUID
	ID, err := uuid.NewV4()
	if err != nil {
		return
	}

	// Setup commons
	res.ID = ID.String()

	res.Schemas = make([]string, len(resType.SchemaExtensions)+1)
	res.Schemas[0] = resType.Schema
	for i, ext := range resType.SchemaExtensions {
		res.Schemas[i+1] = ext.Schema
	}

	now := time.Now()
	res.Meta = core.Meta{
		ResourceType: resType.GetIdentifier(),
		Created:      &now,
		LastModified: &now,
		Version:      version.GenerateVersion(true, ID.String(), now.String()),
	}

	// Attributes whose mutability is "readOnly" SHALL be ignored
	ro, err := attr.Paths(resType, func(attribute *core.Attribute) bool {
		return attribute.Mutability == schemas.MutabilityReadOnly
	})
	if err != nil {
		ret = nil
		return
	}

	for _, p := range ro {
		p.Context(resType).Delete(res)
	}

	err = s.Create(res)
	if err != nil {
		ret = nil
	} else {
		ret, err = query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
	}

	return
}
