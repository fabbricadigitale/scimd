package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=gen_resource_type_repository.go gen "Elem=ResourceType Generic=ResourceType"
//go:generate genny -in=$GOFILE -out=gen_schema_repository.go gen "Elem=Schema Generic=Schema"

// Elem is generic
type Elem generic.Type

// Generic is generic
type Generic generic.Type

type repositoryGeneric struct {
	items map[string]Elem
	mu    sync.RWMutex
}

// GenericRepository is the ...
type GenericRepository interface {
	Get(key string) *Elem
	Add(filename string) (Elem, error)
}

func (repo *repositoryGeneric) Get(key string) *Elem {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	if item, ok := repo.items[key]; ok {
		return &item
	}
	return nil
}

func (repo *repositoryGeneric) Add(filename string) (Elem, error) {
	var data Elem

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	var id string
	if id = interface{}(data).(Identifiable).GetIdentifier(); id == "" {
		return data, errors.New("missing identifier")
	}

	repo.items[id] = data

	return data, nil
}

var (
	repoGeneric *repositoryGeneric
	onceGeneric sync.Once
)

// GetGenericRepository is a singleton repository for core schemas
func GetGenericRepository() GenericRepository {
	onceGeneric.Do(func() {
		repoGeneric = &repositoryGeneric{
			items: make(map[string]Elem),
		}
	})

	return repoGeneric
}
