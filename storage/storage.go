package storage

import (
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/event"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// Iter ...
type Iter interface {
	Next() *resource.Resource
	Done() bool
	Close()
}

// Querier ...
type Querier interface {
	Fields(map[attr.Path]bool) Querier
	Skip(int) Querier
	Limit(int) Querier
	Sort(by attr.Path, asc bool) Querier

	Count() (n int, err error)
	Iter() Iter

	Close()
}

// Storer is the target interface
type Storer interface {
	event.Dispatcher

	Ping() error

	SetIndexes(keys [][]string) error

	Create(res *resource.Resource) error

	Get(resType *core.ResourceType, id, version string, fields map[attr.Path]bool) (*resource.Resource, error)

	Update(resType *resource.Resource, id, version string) error

	Patch(resType *core.ResourceType, id string, version string, op string, path attr.Path, values interface{}) error

	Delete(resType *core.ResourceType, id, version string) error

	Find(resType []*core.ResourceType, filter filter.Filter) (Querier, error)

	Close()
}
