package update

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
	uuid "github.com/satori/go.uuid"
)

// Resource update an existing res of type resType and stores it into s.
func Resource(s storage.Storer, resType *core.ResourceType, id string, res *resource.Resource) (ret core.ResourceTyper, err error) {

	// Make a new UUID
	ID, err := uuid.NewV4()
	if err != nil {
		return
	}

	// Setup commons
	res.ID = ID.String()

	now := time.Now()
	res.Meta.LastModified = &now
	res.Meta.Version = version.GenerateVersion(true, res.ID, now.String())

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

	// We need to perform mutability validation
	// 1. Attributes whose mutability is "readWrite" that are omitted from
	// the request body MAY be assumed to be not asserted by the client.

	// 2. (Immutable attributes) If one or more values are already set for the attribute,
	// the input value(s) MUST match, or HTTP status code 400 SHOULD be
	// returned with a "scimType" error code of "mutability".

	// 3. (ReadOnly) Any values provided SHALL be ignored. (performed by the client)

	/* storedResource, err := a.Get(resource.ResourceType(), id, version, nil)
	if err != nil {
		return err
	}
	*/

	// (todo) => Test immutable attributes
	ro, err = attr.Paths(resType, func(attribute *core.Attribute) bool {
		return attribute.Mutability == schemas.MutabilityImmutable

	})
	if err != nil {
		return
	}

	if len(ro) > 0 {

		attrs := &api.Attributes{}
		attrs.Attributes = make([]string, 0)
		for _, p := range ro {
			attrs.Attributes = append(attrs.Attributes, p.String())
		}

		// we need to get stored value of immutable attributes
		sr, err := query.Resource(s.(storage.Storer), resType, id, attrs)
		if err != nil {
			return nil, err
		}

		storedResource := sr.(*resource.Resource)

		for _, p := range ro {

			if p.Context(resType).Get(storedResource) != p.Context(resType).Get(res) {
				err = &api.MutabilityError{
					Path: p.String(),
				}
				return nil, err
			}
		}
	}

	err = s.Update(res, id, "")
	if err != nil {
		ret = nil
	} else {
		ret, err = query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
	}

	return

}
