package mongo

import (
	"regexp"
	"time"

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

var _ storage.Storer = (*Adapter)(nil)
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

// Ping ...
func (a *Adapter) Ping() error {
	return a.adaptee.session.Ping()
}

// Create is ...
func (a *Adapter) Create(res *resource.Resource) error {
	dataResource := a.toDoc(res)
	return (*a.adaptee).Create(dataResource)
}

// Get is ...
func (a *Adapter) Get(resType *core.ResourceType, id, version string, fields map[attr.Path]bool) (*resource.Resource, error) {
	q, close, err := (*a.adaptee).Find(makeQuery(resType.GetIdentifier(), id, version))
	defer close()

	if err != nil {
		return nil, err
	}

	query := Query{q}
	query.Fields(fields)
	return query.one()
}

// Update is ...
func (a *Adapter) Update(resource *resource.Resource, id string, version string) error {
	dataResource := a.toDoc(resource)
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
		"$or": or,
	}

	query, close, err := (*a.adaptee).Find(_q)
	defer close()
	if err != nil {
		return nil, err
	}
	return &Query{
		q: query,
	}, nil
}

func pathToKey(p attr.Path) string {
	ep := p.Transform(keyEscape)
	if ep.Undefined() {
		return notExistingKey
	}

	if ep.URI == "" {
		return ep.String()
	}

	ns := ep.URI
	ep.URI = ""
	return ns + "." + ep.String()
}

func makeQuery(resType, id, version string) bson.M {
	q := bson.M{
		"id":                id,
		"meta.resourceType": resType,
	}

	if version != "" {
		q["meta.version"] = version
	}

	return q
}

// This method translate resource.Resource to a ready-to-store document
// Document's structure is define as following:
//  - Common attributes are placed as root keys
//  - For each schema (including base one) a key equals to the corrisponding schema's URI holds an object populated with the corrisponding complex
//  - Complex attributes are converted to mongo objects with corrisponding keys and nested fields
func (a *Adapter) toDoc(r *resource.Resource) *document {

	rt := r.ResourceType()

	d := document{
		"schemas":    r.Schemas,
		"id":         r.ID,
		"externalId": r.ExternalID,
		"meta":       fromMeta(&r.Meta),
	}

	for ns := range rt.GetSchemas() {
		if c := map[string]interface{}(*r.Values(ns)); c != nil {
			d[ns] = c
		}
	}
	return &d
}

func toResource(d *document) *resource.Resource {

	dd := (*d)

	r := &resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas: toStringSlice(dd["schemas"].([]interface{})),
			ID:      dd["id"].(string),
			Meta:    toMeta(dd["meta"].(bson.M)),
		},
	}

	rt := r.ResourceType()

	for ns, s := range rt.GetSchemas() {
		if values := dd[ns]; values != nil {
			c, err := s.Enforce(values.(bson.M))
			if err != nil {
				panic(err)
			}
			r.SetValues(ns, c)
		}
	}

	return r
}

func toStringSlice(iSlice []interface{}) []string {
	len := len(iSlice)
	slice := make([]string, len)

	for i, val := range iSlice {
		slice[i] = val.(string)
	}
	return slice
}

func fromMeta(meta *core.Meta) map[string]interface{} {
	if meta == nil {
		return nil
	}

	m := map[string]interface{}{
		"location":     meta.Location,
		"resourceType": meta.ResourceType,
		"created":      meta.Created,
		"lastModified": meta.LastModified,
	}

	// version is not a required field, it is omitted if empty
	if meta.Version != "" {
		m["version"] = meta.Version
	}

	return m
}

func toMeta(m map[string]interface{}) core.Meta {
	created := m["created"].(time.Time)
	lastMod := m["lastModified"].(time.Time)
	meta := core.Meta{
		Created:      &created,
		LastModified: &lastMod,
		Location:     m["location"].(string),
		ResourceType: m["resourceType"].(string),
	}

	// version is not a required field, it is omitted if empty
	if m["version"] != nil {
		meta.Version = m["version"].(string)
	}

	return meta
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
	m["meta.resourceType"] = resType.GetIdentifier()
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

// Represent a mongo key that's always not present
const notExistingKey = "_"

func (c *convert) relationalOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	// If any schema attribure was not found node.Value is nil.
	// For filtered attributes that are not part of a particular resource
	// type, the service provider SHALL treat the attribute as if there is
	// no attribute value, as per https://tools.ietf.org/html/rfc7644#section-3.4.2.1
	if node.Path.Undefined() {
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
	key := pathToKey(node.Path)
	value := node.Value.(string)

	switch node.Op {
	case filter.OpContains:
		return regexQueryPart(key, value, "i", "", "")
	case filter.OpStartsWith:
		return regexQueryPart(key, value, "i", "^", "")
	case filter.OpEndsWith:
		return regexQueryPart(key, value, "i", "", "$")
	default:
		return nil
	}
}

func regexQueryPart(key, value, option, prePattern, postPattern string) bson.M {
	return bson.M{
		key: bson.M{
			"$regex": bson.RegEx{
				Pattern: prePattern + regexp.QuoteMeta(value) + postPattern,
				Options: option,
			},
		},
	}
}

func comparisonOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	key := pathToKey(node.Path)
	return bson.M{
		key: bson.M{
			mapOperator[node.Op]: node.Value,
		},
	}

}

func prOperator(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {

	attrDef := node.Path.Context(resType).Attribute
	key := pathToKey(node.Path)

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
		key: bson.M{"$and": criterion},
	}
}
