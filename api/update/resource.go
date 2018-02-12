package update

import (
	"time"

	"github.com/fabbricadigitale/scimd/required"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/fabbricadigitale/scimd/version"
	uuid "github.com/satori/go.uuid"
)

// Resource update an existing res of type resType and stores it into s.
func Resource(s storage.Storer, resType *core.ResourceType, id string, res *resource.Resource) (ret core.ResourceTyper, err error) {

	err = required.ValidateRequired(res)
	if err != nil {
		return
	}

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

	err = s.Update(res, id, "")
	if err != nil {
		switch err.(type) {
		case mongo.ResourceNotFoundError:
			err = messages.NewError(&api.ResourceNotFoundError{
				Detail: id,
			})
		}
		ret = nil
	} else {
		ret, err = query.Resource(s.(storage.Storer), resType, res.ID, &api.Attributes{})
	}

	return

}
