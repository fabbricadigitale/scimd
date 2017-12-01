package storage

import (
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

type Iter interface {
	Next() *resource.Resource
	Done() bool
	Close()
}

type Querier interface {
	Fields(included []*attr.Path, excluded []*attr.Path) Querier
	Skip(int) Querier
	Limit(int) Querier
	Sort(by *attr.Path, asc bool) Querier

	Count() (n int, err error)
	Iter() Iter
}

// Storer is the target interface
type Storer interface {
	Create(res *resource.Resource) error

	Get(resType *core.ResourceType, id, version string) (*resource.Resource, error)

	Update(resType *resource.Resource, id, version string) error

	Delete(resType *core.ResourceType, id, version string) error

	Find(resType []*core.ResourceType, filter filter.Filter) (Querier, error)
}
