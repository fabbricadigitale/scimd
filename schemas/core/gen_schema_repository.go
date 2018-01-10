// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
)

type repositorySchema struct {
	items map[string]Schema
	mu    sync.RWMutex
}

// SchemaRepository is the ...
type SchemaRepository interface {
	Get(key string) *Schema // (fixme) > evaluate whether make senses to do not return a pointer ...
	Add(filename string) (Schema, error)
	List() []Schema
}

// List returns all elements
func (repo *repositorySchema) List() []Schema {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	res := make([]Schema, len(repo.items))
	i := 0
	for _, elem := range repo.items {
		res[i] = elem
		i++
	}

	return res
}

// Get provides the element for a given key, or nil if it does not exist within the repository.
func (repo *repositorySchema) Get(key string) *Schema {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	if item, ok := repo.items[key]; ok {
		return &item
	}
	return nil
}

// Add allows to load an element and to store it within this repository
func (repo *repositorySchema) Add(filename string) (Schema, error) {
	var data Schema

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
	repoSchema *repositorySchema
	onceSchema sync.Once
)

// GetSchemaRepository is a singleton repository for core schemas
func GetSchemaRepository() SchemaRepository {
	onceSchema.Do(func() {
		repoSchema = &repositorySchema{
			items: make(map[string]Schema),
		}
	})

	return repoSchema
}
