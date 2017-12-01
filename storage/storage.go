package storage

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

//Storage is the target interface
type Storage interface {
	Create(res *resource.Resource) error

	Get(resType *core.ResourceType, id, version string) (*resource.Resource, error)

	Count() error // (todo)

	Update(resType *resource.Resource, id, version string) error

	Delete(resType *core.ResourceType, id, version string) error

	Search(resTypes []*core.ResourceType, search *api.Search) error
}

// Manager is ...
type Manager struct{}

// CreateAdapter is ...
func (m *Manager) CreateAdapter(t, url, db, collection string) (Storage, error) {

	switch t {
	// (fixme) Do NOT import child packages
	// case "mongo":
	// 	return mongo.GetAdapter(url, db, collection)
	default:
		return nil, nil
	}

}
