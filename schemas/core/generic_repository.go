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
	Pull(key string) *Elem // (fixme) > evaluate whether make senses to do not return a pointer ...
	Push(elem Elem) (Elem, error)
	PushFromFile(filename string) (Elem, error)
	PushFromData(data []byte) (Elem, error)
	List() []Elem
	Clean()
}

// List returns all elements
func (repo *repositoryGeneric) List() []Elem {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	res := make([]Elem, len(repo.items))
	i := 0
	for _, elem := range repo.items {
		res[i] = elem
		i++
	}

	return res
}

// Pull provides the element for a given key, or nil if it does not exist within the repository.
func (repo *repositoryGeneric) Pull(key string) *Elem {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	if item, ok := repo.items[key]; ok {
		return &item
	}
	return nil
}

// PushFromData allows to store an elem within its repository
func (repo *repositoryGeneric) Push(elem Elem) (Elem, error) {
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
func (repo *repositoryGeneric) PushFromData(data []byte) (Elem, error) {
	var elem Elem

	err := json.Unmarshal(data, &elem)
	if err != nil {
		return elem, err
	}

	return repo.Push(elem)
}

// PushFromFile allows to load an element from file system and to store it within its repository
func (repo *repositoryGeneric) PushFromFile(filename string) (Elem, error) {
	var elem Elem

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return elem, err
	}
	return repo.PushFromData(bytes)
}

// Clean empties the repository
func (repo *repositoryGeneric) Clean() {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for k := range repo.items {
		delete(repo.items, k)
	}
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
