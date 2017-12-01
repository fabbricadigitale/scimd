package query

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage"
)

func makeAttrs(list []string) ([]*attr.Path, error) {
	ret := make([]*attr.Path, len(list))
	for i, a := range list {
		p := attr.Parse(a)
		if p.Valid() {
			ret[i] = p
		} else {
			return nil, &api.InvalidPathError{
				Path: a,
			}
		}
	}
	return ret, nil
}

func makeAttr(a string) (*attr.Path, error) {
	p := attr.Parse(a)
	if p.Valid() {
		return p, nil
	}
	return nil, &api.InvalidPathError{
		Path: a,
	}
}

func Resource(s storage.Storer, resType *core.ResourceType, id string, attrs *api.Attributes) (core.ResourceTyper, error) {

	// (todo) Fields projection

	res, err := s.Get(resType, id, "")

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, &api.NotFoundError{Subject: id}
	}

	return res, nil
}

func Resources(s storage.Storer, resTypes []*core.ResourceType, search *api.Search) (list *messages.ListResponse, err error) {

	// Make filter
	var f filter.Filter
	if len(search.Filter) > 0 {
		f, err = filter.CompileString(string(search.Filter))
		if err != nil {
			return
		}
	}

	// Make query
	var q storage.Querier
	q, err = s.Find(resTypes, f)
	if err != nil {
		return
	}

	// Fields projection
	var in, ex []*attr.Path

	if len(search.Attributes.Attributes) > 0 {
		in, err = makeAttrs(search.Attributes.Attributes)
		if err != nil {
			return
		}
	}

	if len(search.Attributes.ExcludedAttributes) > 0 {
		in, err = makeAttrs(search.Attributes.ExcludedAttributes)
		if err != nil {
			return
		}
	}

	q.Fields(in, ex)

	// Count
	list.TotalResults, err = q.Count()
	if err != nil {
		return
	}

	// Make list
	list = messages.NewListResponse()

	// Pagination
	q.Skip(search.StartIndex).Limit(search.Count)
	list.StartIndex = search.StartIndex
	list.ItemsPerPage = search.Count

	// Sorting
	if search.SortBy != "" {
		var sortBy *attr.Path
		sortBy, err = makeAttr(search.SortBy)
		if err != nil {
			return
		}
		q.Sort(sortBy, search.SortOrder != api.DescendingOrder)
	}

	// Finally, fetch resources
	iter := q.Iter()
	for r := iter.Next(); r != nil; {
		list.Resources = append(list.Resources, r)
	}

	return
}
