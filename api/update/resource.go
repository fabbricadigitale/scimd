package update

import (
	"time"

	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/version"
)

// Resource update an existing res of type resType and stores it into s.
func Resource(s storage.Storer, id string, res *resource.Resource) (err error) {

	now := time.Now()
	res.Meta.LastModified = &now
	res.Meta.Version = version.GenerateVersion(true, id, now.String())

	err = s.Update(res, id, "")
	// (todo) Reload res?
	return

}
