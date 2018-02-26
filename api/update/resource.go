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

	err = s.Update(res, id, "")
	if err != nil {

		ret = nil
	} else {
		ret, err = query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
	}

	return

}
