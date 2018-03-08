package query

import (
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/mold"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/hasher"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
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

// Attributes retrieves a map of the fields to realize the projection
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

// Resource retrieves a resource filtering by id and resourceType.
// The 'attrs' parameter allows a projection of the attributes of the resource that is returned to the client.
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

// Resources retrieves a list of resources filtering by the search.Filter and the resourceTypes' array.
func Resources(s storage.Storer, resTypes []*core.ResourceType, search *api.Search) (list *messages.ListResponse, err error) {

	// (TODO) > wrap errors here
	if err = validation.Validator.Var(resTypes, "gt=0"); err != nil {
		return
	}

	if err = validation.Validator.Struct(search); err != nil {
		return
	}

	if err = mold.Transformer.Struct(nil, search); err != nil {
		return
	}

	filterString := string(search.Filter)

	hasher := hasher.NewBCrypt()
	var exclude []string
	exclude = make([]string, 0)

	// Make filter
	var f filter.Filter
	if len(search.Filter) > 0 {
		f, err = filter.CompileString(filterString)
		if err != nil {
			return
		}
	}

	var passwords map[string]passwordInfo
	var cs customizer
	pwd := &passwordInfo{
		not: false,
	}

	// If there is a filter with a password, next lines change the filter
	// The new filter will be more inclusive.
	if len(search.Filter) > 0 {
		passwords = make(map[string]passwordInfo)
		for _, resType := range resTypes {
			cs.customize(resType, &f, pwd)

			if f.String() != filterString {
				passwords[resType.ID] = *pwd
			}
		}
	}

	// Make query
	var q storage.Querier
	q, err = s.Find(resTypes, f)
	defer q.Close()
	if err != nil {
		return
	}

	// compare plain password with the stored hashed password
	// and exclude item that do not match the logic of the filter query.
	if len(passwords) > 0 {

		var res *resource.Resource
		for it := q.Iter(); !it.Done(); {
			res = it.Next()

			resourceType := res.Meta.ResourceType

			// (FIXME) => make this urn configurable
			values := res.Values("urn:ietf:params:scim:schemas:core:2.0:User")

			hashedPassword := (*values)["password"].(datatype.String)

			b := hasher.Compare([]byte(hashedPassword), []byte(passwords[resourceType].value))

			if passwords[resourceType].operator == "eq" && !passwords[resourceType].not && !b {
				exclude = append(exclude, res.ID)
			}

			if passwords[resourceType].operator == "eq" && passwords[resourceType].not && b {
				exclude = append(exclude, res.ID)
			}

			if passwords[resourceType].operator == "ne" && !passwords[resourceType].not && b {
				exclude = append(exclude, res.ID)
			}

			if passwords[resourceType].operator == "ne" && passwords[resourceType].not && !b {
				exclude = append(exclude, res.ID)
			}
		}
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
	// Remove excluded
	list.TotalResults -= len(exclude)

	if search.Count > config.Values.PageSize {
		search.Count = config.Values.PageSize
	}

	if search.Count == 0 || search.Count > list.TotalResults {
		search.Count = list.TotalResults
	}

	// Pagination
	q.Skip(search.StartIndex - 1).Limit(search.Count - (search.StartIndex - 1))
	list.StartIndex = search.StartIndex

	if search.Count-(list.StartIndex-1) >= 0 {
		list.ItemsPerPage = search.Count - (list.StartIndex - 1)
	} else {
		list.ItemsPerPage = 0
	}

	// Sorting
	if search.SortBy != "" {
		var sortBy *attr.Path
		sortBy, err = makeAttr(search.SortBy)
		if err != nil {
			return
		}
		q.Sort(*sortBy, search.SortOrder != api.DescendingOrder)
	}

	list.Resources = make([]interface{}, 0)
	// Finally, fetch resources
	var r *resource.Resource
	for iter := q.Iter(); !iter.Done(); {
		r = iter.Next()
		if !contains(exclude, r.ID) {
			list.Resources = append(list.Resources, r)
		}
	}

	return
}

func contains(slice []string, ID string) bool {
	for _, value := range slice {
		if value == ID {
			return true
		}
	}
	return false
}

type customizer struct{}
type passwordInfo struct {
	value    string
	operator string
	not      bool
}

func (c *customizer) customize(resType *core.ResourceType, ft *filter.Filter, pwd *passwordInfo) {

	switch (*ft).(type) {

	case filter.Group:
		node := (*ft).(filter.Group)
		var filter filter.Group
		c.customize(resType, &node.Filter, pwd)
		filter.Filter = node.Filter
		*ft = filter

	case filter.And:
		node := (*ft).(filter.And)
		var filter filter.And
		if node.Left != nil {
			c.customize(resType, &node.Left, pwd)
			filter.Left = node.Left
		}
		if node.Right != nil {
			c.customize(resType, &node.Right, pwd)
			filter.Right = node.Right
		}
		(*ft) = filter

	case filter.Or:
		node := (*ft).(filter.Or)
		var filter filter.Or
		if node.Left != nil {
			c.customize(resType, &node.Left, pwd)
			filter.Left = node.Left
		}
		if node.Right != nil {
			c.customize(resType, &node.Right, pwd)
			filter.Right = node.Right
		}
		(*ft) = filter

	case filter.Not:
		node := (*ft).(filter.Not)
		var ftn filter.Not

		switch node.Filter.(type) {
		case *filter.AttrExpr:
			pwd.not = true
		}

		c.customize(resType, &node.Filter, pwd)
		ftn.Filter = node.Filter

		(*ft) = ftn

	case *filter.AttrExpr:
		filter := (*ft).Normalize(resType).(*filter.AttrExpr)
		if filter.Path.URI == "urn:ietf:params:scim:schemas:core:2.0:User" && filter.Path.Name == "password" && (filter.Op == "eq" || filter.Op == "ne") {

			pwd.value = filter.Value.(string)
			pwd.operator = filter.Op

			// Change the filter with a more inclusive one
			// and we delegate this search logic to the next step
			// when we can iterate trough query resultset where there is also password attribute
			// so we can compare hashed password and plain password

			// NOTE: Ensure the custom search logic implemented for the 'eq' and 'ne' operators is enough to represent all compare password operations
			filter.Op = (map[bool]string{false: "ne", true: "eq"})[pwd.not]
			filter.Value = ""
			*ft = filter
		}
	}
}
