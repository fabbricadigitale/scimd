package delete

import (
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage"
)

// Resource update an existing res of type resType and stores it into s.
func Resource(s storage.Storer, resType *core.ResourceType, id string) (err error) {

	err = s.Delete(resType, id, "")
	if err != nil {
		return err
	}
	return
}
