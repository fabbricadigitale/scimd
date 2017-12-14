package mongo

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const RootKey = "data"

// Query is
type Query struct {
	q *mgo.Query
}

// Iter is
type Iter struct {
	i *mgo.Iter
}

var _ storage.Querier = (*Query)(nil)
var _ storage.Iter = (*Iter)(nil)

// Count is...
func (res *Query) Count() (n int, err error) {
	return res.q.Count()
}

// Sort is
func (res *Query) Sort(by *attr.Path, asc bool) storage.Querier {
	if asc {
		res.q = res.q.Sort(by.String())
	} else {
		res.q = res.q.Sort("-" + by.String())
	}
	return res
}

// Skip is
func (res *Query) Skip(index int) storage.Querier {
	res.q = res.q.Skip(index - 1)
	return res
}

// Limit is
func (res *Query) Limit(n int) storage.Querier {
	res.q = res.q.Limit(n)
	return res
}

// Fields is
func (res *Query) Fields(included []*attr.Path, excluded []*attr.Path) storage.Querier {

	var field bson.M
	field = make(bson.M)

	var key string
	for _, val := range included {
		key = fmt.Sprintf("%s.%s", RootKey, val.String())
		field[key] = 1
	}

	for _, val := range excluded {
		key = fmt.Sprintf("%s.%s", RootKey, val.String())
		field[key] = 0
	}

	res.q = res.q.Select(field)
	return res
}

func (res *Query) one() (*resource.Resource, error) {
	resDoc := &resourceDocument{}
	err := res.q.One(resDoc)
	if err != nil {
		return nil, err
	}
	return toResource(resDoc), nil
}

// Iter executes the query and returns an iterator capable of going over all
// the results.
func (res *Query) Iter() storage.Iter {
	return &Iter{
		i: res.q.Iter(),
	}
}

// Next retrieves the next resource from the result set, blocking if necessary.
func (it *Iter) Next() *resource.Resource {

	resDoc := &resourceDocument{}
	n := it.i.Next(resDoc)
	if !n {
		return nil
	}
	return toResource(resDoc)
}

// Done returns true only if a follow up Next call is guaranteed
// to return false.
func (it *Iter) Done() bool {
	return it.i.Done()
}

// Close kills the current iterator.
func (it *Iter) Close() {
	// TODO handle error
	it.i.Close()
}
