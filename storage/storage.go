package storage

import (
	"github.com/fabbricadigitale/scimd/schemas/core/resource"
	"github.com/fabbricadigitale/scimd/storage/mongo"
)

//Storage is the target interface
type Storage interface {
	Create(*resource.Resource) error

	Get(id, version string) (*resource.Resource, error)

	Count() error

	Update(id, version string, resource *resource.Resource) error

	Delete(id, version string) error

	Search() error
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
