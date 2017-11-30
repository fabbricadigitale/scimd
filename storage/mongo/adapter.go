package mongo

import (
	"fmt"
	"reflect"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"gopkg.in/mgo.v2/bson"
)

// Adapter is the repository Adapter
type Adapter struct {
	adaptee *Driver
}

// (fixme) var _ storage.Storer = (*Adapter)(nil)
// (fixme) global adapter must be avoided
var (
	adapter Adapter
	// OpEqual              = "eq"
	// OpNotEqual           = "ne"
	// OpContains           = "co"
	// OpStartsWith         = "sw"
	// OpEndsWith           = "ew"
	// OpGreaterThan        = "gt"
	// OpLessThan           = "lt"
	// OpGreaterOrEqualThan = "ge"
	// OpLessOrEqualThan    = "le"
	// OpPresent            = "pr"
	mapOperator = map[string]string{
		filter.OpEqual:              "$eq",
		filter.OpNotEqual:           "$ne",
		filter.OpGreaterThan:        "$gt",
		filter.OpLessThan:           "$lt",
		filter.OpGreaterOrEqualThan: "$gte",
		filter.OpLessOrEqualThan:    "$lte",
	}
)

// urnKey identifies the attributes namespace into document resource
// The name stars with an underscore unlike scim properties that start with alphabetical characters
const urnKey = "_urn"

// GetAdapter ...
func GetAdapter(url, db, collection string) (*Adapter, error) {

	if (Adapter{}) == adapter {
		driver, err := CreateDriver(url, db, collection)
		if err != nil {
			return nil, err
		}
		adapter.adaptee = driver
	}
	return &adapter, nil
}

// Create is ...
func (a *Adapter) Create(res *resource.Resource) error {

	dataResource := a.hydrateResource(res)
	return (*a.adaptee).Create(dataResource)
}

// Get is ...
func (a *Adapter) Get(resType *core.ResourceType, id, version string) (*resource.Resource, error) {

	h := &resourceDocument{}

	h, err := (*a.adaptee).Get(id, version)

	if err != nil {
		return nil, err
	}

	return a.toResource(h)
}

// Count ...
func (a *Adapter) Count() error {
	return (*a.adaptee).Count()
}

// Update is ...
func (a *Adapter) Update(resource *resource.Resource, id string, version string) error {
	dataResource := a.hydrateResource(resource)
	return (*a.adaptee).Update(id, version, dataResource)
}

// Delete is ...
func (a *Adapter) Delete(resType *core.ResourceType, id, version string) error {
	return (*a.adaptee).Delete(id, version)
}

// Search is ...
func (a *Adapter) Search(resTypes []*core.ResourceType, search *api.Search) error {
	q, _ := convertToMongoQuery(search)

	return (*a.adaptee).Search(q)
}

// resourceDocument is a ready-to-store format for Resource
type resourceDocument struct {
	Data []map[string]interface{}
}

// This method translate Resource to a ready-to-store document
// The document has a Data property, array of []map[string]interface{},  with a fixed order:
// index = 0 -> common attributes
// index = 1 -> core attributes
// index > 1 -> extensions attributes
func (a *Adapter) hydrateResource(r *resource.Resource) *resourceDocument {

	h := &resourceDocument{}

	common := make(map[string]interface{})
	common["schemas"] = r.Schemas
	common["id"] = r.ID
	common["external_id"] = r.ExternalID
	common["meta"] = r.Meta

	rt := r.ResourceType()

	mCore := make(map[string]interface{})
	mCore[urnKey] = rt.GetSchema().ID
	for key, val := range *r.Values(rt.GetSchema().ID) {
		mCore[key] = val
	}
	h.Data = append(h.Data, common, mCore)

	for _, extSch := range rt.GetSchemaExtensions() {
		mExt := make(map[string]interface{})
		if extSch != nil {
			ns := extSch.GetIdentifier()
			mExt[urnKey] = ns
			for key, val := range *r.Values(ns) {
				mExt[key] = val
			}
		}
		h.Data = append(h.Data, mExt)
	}

	return h
}

func (a *Adapter) toResource(h *resourceDocument) (*resource.Resource, error) {

	hCommon := h.Data[0]
	r := &resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    hCommon["schemas"].([]string),
			ID:         hCommon["id"].(string),
			ExternalID: hCommon["external_id"].(string),
			Meta:       hCommon["meta"].(core.Meta),
		},
	}

	var p *datatype.Complex
	for i := 1; i < len(h.Data); i++ {
		ns := h.Data[i][urnKey].(string)
		values := h.Data[i]
		delete(values, urnKey)
		(*p) = datatype.Complex(values)
		r.SetValues(ns, p)
	}

	return r, nil
}

func convertToMongoQuery(query *api.Search) (m bson.M, err error) {

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				err = r.(error)
			default:
				err = &api.InternalServerError{
					Detail: r.(string),
				}
			}
		}
	}()

	f, err := filter.CompileString(string(query.Filter))
	if err != nil {
		return nil, err
	}
	var conv *convert
	m, err = conv.do(f), nil
	return m, err
}

type convert struct{}

func (c *convert) do(f filter.Filter) bson.M {

	var (
		left, right bson.M
	)

	switch f.(type) {

	case *filter.And:
		node := f.(*filter.And)
		if node.Left != nil {
			left = c.do(node.Left)
		}
		if node.Right != nil {
			right = c.do(node.Right)
		}
		return bson.M{
			"$and": []interface{}{left, right},
		}
	case *filter.Or:
		node := f.(*filter.Or)
		if node.Left != nil {
			left = c.do(node.Left)
		}
		if node.Right != nil {
			right = c.do(node.Right)
		}
		return bson.M{
			"$or": []interface{}{left, right},
		}
	case *filter.Not:
		node := f.(*filter.Not)
		left = c.do(node.Filter)
		return bson.M{
			"$nor": []interface{}{left},
		}
	case *filter.AttrExpr:
		node := f.(*filter.AttrExpr)

		// The 'co', 'sw' and ew operators can only be used if the attribute type is string
		if node.Op == filter.OpContains || node.Op == filter.OpStartsWith || node.Op == filter.OpEndsWith {
			// (TODO) > checks attribute type (refs https://github.com/fabbricadigitale/scimd/issues/32)
			if reflect.ValueOf(node.Value).Kind() != reflect.String {
				detail := fmt.Sprintf("Cannot use %s operator with non-string value", node.Op)

				var e *api.InvalidFilterError
				e = &api.InvalidFilterError{
					Filter: f.String(),
					Detail: detail,
				}
				panic(e)
			}

			switch node.Op {
			case filter.OpContains:
				return bson.M{
					node.Path.String(): bson.M{
						"$regex": bson.RegEx{
							Pattern: node.Value.(string),
							Options: "i",
						},
					},
				}
			case filter.OpStartsWith:
				return bson.M{
					node.Path.String(): bson.M{
						"$regex": bson.RegEx{
							Pattern: "^" + node.Value.(string),
							Options: "i",
						},
					},
				}
			case filter.OpEndsWith:
				return bson.M{
					node.Path.String(): bson.M{
						"$regex": bson.RegEx{
							Pattern: node.Value.(string) + "$",
							Options: "i",
						},
					},
				}
			}

		} else {
			return bson.M{
				node.Path.String(): bson.M{
					mapOperator[node.Op]: node.Value,
				},
			}
		}

	}

	return nil
}
