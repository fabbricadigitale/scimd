package query

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/validation"
)

func makeAttrs(list []string) ([]*attr.Path, error) {
	ret := make([]*attr.Path, len(list))
	for i, a := range list {
		p := attr.Parse(a)
		if !p.Undefined() {
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
	if !p.Undefined() {
		return p, nil
	}
	return nil, &api.InvalidPathError{
		Path: a,
	}
}

func Attributes(resTypes []*core.ResourceType, attrs *api.Attributes) (fields map[attr.Path]bool, err error) {
	var in, ex []*attr.Path
	fields = make(map[attr.Path]bool)

	// When specified, the default list of attributes SHALL be
	// overridden, and each resource returned MUST contain the
	// minimum set of resource attributes and any attributes or
	// sub-attributes explicitly requested by the "attributes"
	// parameter (https://tools.ietf.org/html/rfc7644#section-3.9, https://tools.ietf.org/html/rfc7644#section-3.4.2.5)
	if len(attrs.Attributes) > 0 {
		in, err = makeAttrs(attrs.Attributes)
		if err != nil {
			return
		}
	}

	// When specified, each resource returned MUST
	// contain the minimum set of resource attributes.
	// Additionally, the default set of attributes minus those
	// attributes listed in "excludedAttributes" is returned (https://tools.ietf.org/html/rfc7644#section-3.9)
	// (todo) > Specifing excludedAttribute whose schema "returned" parameter setting is "always" has no effect (https://tools.ietf.org/html/rfc7644#section-3.4.2.5)
	if len(attrs.ExcludedAttributes) > 0 {
		ex, err = makeAttrs(attrs.ExcludedAttributes)
		if err != nil {
			return
		}
	}

	// Fields projection
	for _, rt := range resTypes {
		projection, err := attr.Projection(rt, in, ex)
		if err != nil {
			return nil, err
		}
		for _, p := range projection {
			fields[*p] = true
		}
	}

	return
}

func Resource(s storage.Storer, resType *core.ResourceType, id string, attrs *api.Attributes) (res core.ResourceTyper, err error) {
	fields, err := Attributes([]*core.ResourceType{resType}, attrs)
	if err != nil {
		return
	}

	res, err = s.Get(resType, id, "", fields)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, &api.NotFoundError{Subject: id}
	}

	return res, nil
}

func Resources(s storage.Storer, resTypes []*core.ResourceType, search *api.Search) (list *messages.ListResponse, err error) {
	if err = validation.Validator.Struct(search); err != nil {
		return
	}

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
	defer q.Close()
	if err != nil {
		return
	}

	// Fields projection
	fields, err := Attributes(resTypes, &search.Attributes)
	if err != nil {
		return
	}
	if fields != nil {
		q.Fields(fields)
	}

	// Make list
	list = messages.NewListResponse()

	// Count
	list.TotalResults, err = q.Count()
	if err != nil {
		return
	}

	// Unlimited
	if search.Count == 0 {
		search.Count = list.TotalResults
		// (todo) > We need a way to LIMIT this to a MAX value (from config) - issue https://github.com/fabbricadigitale/scimd/issues/55
	}

	// Pagination
	q.Skip(search.StartIndex - 1).Limit(search.Count)
	list.StartIndex = search.StartIndex
	list.ItemsPerPage = search.Count

	// Sorting
	if search.SortBy != "" {
		var sortBy *attr.Path
		sortBy, err = makeAttr(search.SortBy)
		if err != nil {
			return
		}
		q.Sort(*sortBy, search.SortOrder != api.DescendingOrder)
	}

	// Finally, fetch resources
	var r *resource.Resource
	for iter := q.Iter(); !iter.Done(); {
		r = iter.Next()
		list.Resources = append(list.Resources, r)
	}

	return
}
