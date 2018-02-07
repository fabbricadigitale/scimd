package update

import (
	"time"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/version"
)

// Resource update an existing res of type resType and stores it into s.
func Resource(s storage.Storer, resType *core.ResourceType, id string, res *resource.Resource) (ret core.ResourceTyper, err error) {

	now := time.Now()
	res.Meta.LastModified = &now
	res.Meta.Version = version.GenerateVersion(true, id, now.String())

	err = s.Update(res, id, "")
	if err != nil {
		ret = nil
	} else {
		ret, err = query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
	}

	return

}
