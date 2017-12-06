package mongo

import (
	"regexp"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"gopkg.in/mgo.v2/bson"
)

// Adapter is the repository Adapter
type Adapter struct {
	adaptee *Driver
}

// (fixme) var _ storage.Storer = (*Adapter)(nil)
// (fixme) global adapter must be avoided
var (
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

// uriKey identifies the attributes namespace into document resource
// The name stars with an underscore unlike scim properties that start with alphabetical characters
const uriKey = "_uri"
const notExistingKey = "_"

// New makes and return a new adapter of type storage.Storer using a mongo driver
func New(url, db, collection string) (storage.Storer, error) {

	adapter := &Adapter{}
	driver, err := CreateDriver(url, db, collection)
	if err != nil {
		return nil, err
	}
	adapter.adaptee = driver

	return adapter, nil
}

// Create is ...
func (a *Adapter) Create(res *resource.Resource) error {

	dataResource := a.hydrateResource(res)
	return (*a.adaptee).Create(dataResource)
}

// Get is ...
func (a *Adapter) Get(resType *core.ResourceType, id, version string, included []*attr.Path, excluded []*attr.Path) (*resource.Resource, error) {

	q, err := (*a.adaptee).Find(makeQuery(resType.GetIdentifier(), id, version))

	if err != nil {
		return nil, err
	}

	query := Query{q}
	query.Fields(included, excluded)
	return query.one()
}

// Update is ...
func (a *Adapter) Update(resource *resource.Resource, id string, version string) error {
	dataResource := a.hydrateResource(resource)
	return (*a.adaptee).Update(makeQuery(resource.ResourceType().GetIdentifier(), id, version), dataResource)
}

// Delete is ...
func (a *Adapter) Delete(resType *core.ResourceType, id, version string) error {
	return (*a.adaptee).Delete(makeQuery(resType.GetIdentifier(), id, version))
}

// Find is ...
func (a *Adapter) Find(resTypes []*core.ResourceType, filter filter.Filter) (storage.Querier, error) {

	or := make([]bson.M, len(resTypes))

	for i, resType := range resTypes {
		var err error
		or[i], err = convertToMongoQuery(resType, filter)
		if err != nil {
			return nil, err
		}
	}

	_q := bson.M{
		"data": bson.M{
			"$elemMatch": bson.M{
				"$or": or,
			},
		},
	}

	query, err := (*a.adaptee).Find(_q)
	if err != nil {
		return nil, err
	}
	return &Query{
		q: query,
	}, nil
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
	common[uriKey] = ""
	common["schemas"] = r.Schemas
	common["id"] = r.ID
	common["externalId"] = r.ExternalID
	common["meta"] = r.Meta

	rt := r.ResourceType()

	mCore := make(map[string]interface{})
	mCore[uriKey] = rt.GetSchema().ID
	for key, val := range *r.Values(rt.GetSchema().ID) {
		mCore[key] = val
	}
	h.Data = append(h.Data, common, mCore)

	for _, extSch := range rt.GetSchemaExtensions() {
		mExt := make(map[string]interface{})
		if extSch != nil {
			ns := extSch.GetIdentifier()
			mExt[uriKey] = ns
			for key, val := range *r.Values(ns) {
				mExt[key] = val
			}
		}
		h.Data = append(h.Data, mExt)
	}

	return h
}

func toResource(h *resourceDocument) *resource.Resource {

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
		ns := h.Data[i][uriKey].(string)
		values := h.Data[i]
		delete(values, uriKey)
		(*p) = datatype.Complex(values)
		r.SetValues(ns, p)
	}

	return r
}

func convertToMongoQuery(resType *core.ResourceType, ft filter.Filter) (m bson.M, err error) {

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

	var conv *convert
	m, err = conv.do(resType, ft.Normalize(resType)), nil
	m["meta.resouceType"] = resType.GetIdentifier()
	return m, err
}

type convert struct{}

func (c *convert) do(resType *core.ResourceType, f interface{}) bson.M {

	var (
		left, right bson.M
	)

	switch f.(type) {

	case *filter.Group:
		node := f.(*filter.Group)
		return c.do(resType, node.Filter)

	case *filter.And:
		node := f.(*filter.And)
		if node.Left != nil {
			left = c.do(resType, node.Left)
		}
		if node.Right != nil {
			right = c.do(resType, node.Right)
		}
		return bson.M{
			"$and": []interface{}{left, right},
		}
	case *filter.Or:
		node := f.(*filter.Or)
		if node.Left != nil {
			left = c.do(resType, node.Left)
		}
		if node.Right != nil {
			right = c.do(resType, node.Right)
		}
		return bson.M{
			"$or": []interface{}{left, right},
		}
	case *filter.Not:
		node := f.(*filter.Not)
		left = c.do(resType, node.Filter)
		return bson.M{
			"$nor": []interface{}{left},
		}
	case *filter.AttrExpr:
		node := f.(*filter.AttrExpr)
		return c.relationalOperators(resType, f, node)
	}

	return nil
}

func (c *convert) relationalOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {

	// If any schema attribure was not found node.Value is nil.
	// For filtered attributes that are not part of a particular resource
	// type, the service provider SHALL treat the attribute as if there is
	// no attribute value, as per https://tools.ietf.org/html/rfc7644#section-3.4.2.1
	if !node.Path.Valid() {
		return bson.M{
			notExistingKey: bson.M{
				mapOperator[node.Op]: node.Value,
			},
		}
	}

	// The 'co', 'sw' and 'ew' operators can only be used if the attribute type is string
	if node.Op == filter.OpContains || node.Op == filter.OpStartsWith || node.Op == filter.OpEndsWith {
		return stringOperators(resType, f, node)
	} else if node.Op == filter.OpPresent {
		return prOperator(resType, f, node)
	} else {
		return comparisonOperators(resType, f, node)
	}
}

func newInvalidFilterError(detail, filter string) *api.InvalidFilterError {
	var e *api.InvalidFilterError
	e = &api.InvalidFilterError{
		Filter: filter,
		Detail: detail,
	}
	return e
}

func stringOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {

	attrDef := node.Path.FindAttribute(resType)

	var path *attr.Path
	path = &node.Path

	uri, key := convertKey(path)
	value := node.Value.(string)

	if attrDef.MultiValued {

		switch node.Op {

		case filter.OpContains:
			return multiValuedQueryPart(uri, key, value, "i", "", "")
		case filter.OpStartsWith:
			return multiValuedQueryPart(uri, key, value, "i", "^", "")
		case filter.OpEndsWith:
			return multiValuedQueryPart(uri, key, value, "i", "", "$")
		default:
			return nil
		}
	} else {

		switch node.Op {

		case filter.OpContains:
			return singleValueQueryPart(uri, key, value, "i", "", "")
		case filter.OpStartsWith:
			return singleValueQueryPart(uri, key, value, "i", "^", "")
		case filter.OpEndsWith:
			return singleValueQueryPart(uri, key, value, "i", "", "$")
		default:
			return nil
		}
	}
}

func multiValuedQueryPart(uri, key, value, option, prePattern, postPattern string) bson.M {
	return bson.M{
		"$elemMatch": bson.M{
			"$and": []interface{}{
				bson.M{
					key: bson.M{
						"$regex": bson.RegEx{
							Pattern: prePattern + regexp.QuoteMeta(value) + postPattern,
							Options: option,
						},
					},
				},
				bson.M{
					uriKey: bson.M{
						"$eq": uri,
					},
				},
			},
		},
	}
}

func singleValueQueryPart(uri, key, value, option, prePattern, postPattern string) bson.M {
	return bson.M{
		"$and": []interface{}{
			bson.M{
				key: bson.M{
					"$regex": bson.RegEx{
						Pattern: prePattern + regexp.QuoteMeta(value) + postPattern,
						Options: option,
					},
				},
			},
			bson.M{
				uriKey: bson.M{
					"$eq": uri,
				},
			},
		},
	}
}

func convertKey(p *attr.Path) (urn, key string) {
	urn = p.URI
	if p.Valid() {
		key = p.Name
		if p.Sub != "" {
			key += "." + p.Sub
		}
	} else {
		key = notExistingKey
	}
	return
}

func comparisonOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	attrDef := node.Path.FindAttribute(resType)

	var path *attr.Path
	path = &node.Path

	uri, key := convertKey(path)

	if attrDef.MultiValued {
		return bson.M{
			"$elemMatch": bson.M{
				"$and": []interface{}{
					bson.M{
						key: bson.M{
							mapOperator[node.Op]: node.Value,
						},
					},
					bson.M{
						uriKey: bson.M{
							"$eq": uri,
						},
					},
				},
			},
		}

	}
	return bson.M{
		"$and": []interface{}{
			bson.M{
				key: bson.M{
					mapOperator[node.Op]: node.Value,
				},
			},
			bson.M{
				uriKey: bson.M{
					"$eq": uri,
				},
			},
		},
	}

}

func prOperator(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	attrDef := node.Path.FindAttribute(resType)

	var path *attr.Path
	path = &node.Path

	uri, key := convertKey(path)

	existsCriteria := bson.M{key: bson.M{"$exists": true}}
	nullCriteria := bson.M{key: bson.M{"$ne": nil}}
	emptyStringCriteria := bson.M{key: bson.M{"$ne": ""}}
	emptyArrayCriteria := bson.M{key: bson.M{"$not": bson.M{"$size": 0}}}
	emptyObjectCriteria := bson.M{key: bson.M{"$ne": bson.M{}}}

	criterion := make([]interface{}, 0)
	criterion = append(criterion, existsCriteria, nullCriteria)
	if attrDef.MultiValued {
		criterion = append(criterion, emptyArrayCriteria)
	} else {
		switch attrDef.Type {
		case datatype.StringType:
			criterion = append(criterion, emptyStringCriteria)
		case datatype.ComplexType:
			criterion = append(criterion, emptyObjectCriteria)
		}
	}
	return bson.M{
		"$and": []interface{}{
			bson.M{"$and": criterion},
			bson.M{
				uriKey: bson.M{
					"$eq": uri,
				},
			},
		},
	}
}
