package query

import (
	"reflect"

	"github.com/fabbricadigitale/scimd/schemas/datatype"

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

func Resource(s storage.Storer, resType *core.ResourceType, id string, attrs *api.Attributes) (core.ResourceTyper, error) {

	// (todo) Fields projection

	res, err := s.Get(resType, id, "", nil, nil)

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

	// When specified, the default list of attributes SHALL be
	// overridden, and each resource returned MUST contain the
	// minimum set of resource attributes and any attributes or
	// sub-attributes explicitly requested by the "attributes"
	// parameter (https://tools.ietf.org/html/rfc7644#section-3.9, https://tools.ietf.org/html/rfc7644#section-3.4.2.5)
	if len(search.Attributes.Attributes) > 0 {
		in, err = makeAttrs(search.Attributes.Attributes)
		if err != nil {
			return
		}
	}

	// When specified, each resource returned MUST
	// contain the minimum set of resource attributes.
	// Additionally, the default set of attributes minus those
	// attributes listed in "excludedAttributes" is returned (https://tools.ietf.org/html/rfc7644#section-3.9)
	// (todo) > Specifing excludedAttribute whose schema "returned" parameter setting is "always" has no effect (https://tools.ietf.org/html/rfc7644#section-3.4.2.5)
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

func checkAttributeByProperty(attribute *core.Attribute, property, value string) bool {
	a := reflect.ValueOf(attribute)
	v := reflect.Indirect(a).FieldByName(property)
	if v.Interface() == value {
		return true
	}
	return false
}

func getSchemasAttributes(resType *core.ResourceType, property, value string) []*attr.Path {
	var as []*attr.Path
	as = make([]*attr.Path, 0)
	schemas := resType.GetSchemas()
	for _, schema := range schemas {
		for _, attribute := range schema.Attributes {
			if attribute.Type == datatype.ComplexType {
				for _, subAttribute := range attribute.SubAttributes {
					if checkAttributeByProperty(subAttribute, property, value) {
						as = append(as, newContext(schema, attribute, subAttribute).Path())
					}
				}
			} else {
				if checkAttributeByProperty(attribute, property, value) {
					as = append(as, newContext(schema, attribute, nil).Path())
				}
			}

		}

	}
	return as
}

func newContext(schema *core.Schema, attribute *core.Attribute, subAttribute *core.Attribute) *attr.Context {
	ctx := attr.Context{}

	if schema != nil {
		ctx.Schema = schema
	}

	if attribute != nil {
		ctx.Attribute = attribute
	}

	if subAttribute != nil {
		ctx.SubAttribute = subAttribute
	}

	return &ctx
}
