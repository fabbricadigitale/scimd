package storage

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage/mongo"
)

//Storage is the target interface
type Storage interface {
	Create(*resource.Resource) error

	Get(rType core.ResourceType, id, version string) (*resource.Resource, error)

	Count() error

	Update(rType core.ResourceType, id, version string, resource *resource.Resource) error

	Delete(rType core.ResourceType, id, version string) error

	Search(rTypes []core.ResourceType, search api.Search) error
}

// Manager is ...
type Manager struct{}

// CreateAdapter is ...
func (m *Manager) CreateAdapter(t, url, db, collection string) (Storage, error) {

	switch t {
	case "mongo":
		return mongo.GetAdapter(url, db, collection)
	default:
		return nil, nil
	}

}
