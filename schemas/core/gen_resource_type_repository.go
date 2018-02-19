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

type repositoryResourceType struct {
	items map[string]ResourceType
	mu    sync.RWMutex
}

// ResourceTypeRepository is the ...
type ResourceTypeRepository interface {
	Pull(key string) *ResourceType // (fixme) > evaluate whether make senses to do not return a pointer ...
	Push(elem ResourceType) (ResourceType, error)
	PushFromFile(filename string) (ResourceType, error)
	PushFromData(data []byte) (ResourceType, error)
	List() []ResourceType
}

// List returns all elements
func (repo *repositoryResourceType) List() []ResourceType {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	res := make([]ResourceType, len(repo.items))
	i := 0
	for _, elem := range repo.items {
		res[i] = elem
		i++
	}

	return res
}

// Pull provides the element for a given key, or nil if it does not exist within the repository.
func (repo *repositoryResourceType) Pull(key string) *ResourceType {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	if item, ok := repo.items[key]; ok {
		return &item
	}
	return nil
}

// PushFromData allows to store an elem within its repository
func (repo *repositoryResourceType) Push(elem ResourceType) (ResourceType, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var id string
	if id = interface{}(elem).(Identifiable).GetIdentifier(); id == "" {
		return elem, errors.New("missing identifier")
	}

	repo.items[id] = elem

	return elem, nil
}

// PushFromData allows to load an element from bytes and to store it within its repository
func (repo *repositoryResourceType) PushFromData(data []byte) (ResourceType, error) {
	var elem ResourceType

	err := json.Unmarshal(data, &elem)
	if err != nil {
		return elem, err
	}

	return repo.Push(elem)
}

// PushFromFile allows to load an element from file system and to store it within its repository
func (repo *repositoryResourceType) PushFromFile(filename string) (ResourceType, error) {
	var elem ResourceType

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return elem, err
	}
	return repo.PushFromData(bytes)
}

var (
	repoResourceType *repositoryResourceType
	onceResourceType sync.Once
)

// GetResourceTypeRepository is a singleton repository for core schemas
func GetResourceTypeRepository() ResourceTypeRepository {
	onceResourceType.Do(func() {
		repoResourceType = &repositoryResourceType{
			items: make(map[string]ResourceType),
		}
	})

	return repoResourceType
}
