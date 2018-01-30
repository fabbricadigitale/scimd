package mongo

import (
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Query is
type Query struct {
	q *mgo.Query
	c func()
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
func (res *Query) Sort(by attr.Path, asc bool) storage.Querier {
	if asc {
		res.q = res.q.Sort(pathToKey(by))
	} else {
		res.q = res.q.Sort("-" + pathToKey(by))
	}
	return res
}

// Skip is
func (res *Query) Skip(index int) storage.Querier {
	if index < 0 {
		index = 0
	}
	res.q = res.q.Skip(index)
	return res
}

// Limit is
func (res *Query) Limit(n int) storage.Querier {
	// (note) > negative n affects batch size (https://docs.mongodb.com/manual/reference/method/cursor.limit/#negative-values)
	res.q = res.q.Limit(n)
	return res
}

// Fields is
func (res *Query) Fields(fields map[attr.Path]bool) storage.Querier {

	var selector bson.M
	selector = make(bson.M)

	if fields != nil {
		for p, on := range fields {
			var s int
			if on {
				s = 1
			}
			selector[pathToKey(p)] = s
		}
	} // else an empty selector is set

	res.q = res.q.Select(selector)
	return res
}

func (res *Query) Close() {
	res.c()
}

func (res *Query) one() (*resource.Resource, error) {
	defer res.Close()
	d := &document{}
	err := res.q.One(d)
	if err != nil {
		return nil, err
	}

	return toResource(d), nil
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
	d := &document{}
	n := it.i.Next(d)
	if !n {
		return nil
	}
	return toResource(d)
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
