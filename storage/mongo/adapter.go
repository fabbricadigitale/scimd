package mongo

import (
	"fmt"
	"time"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/event"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/globalsign/mgo/bson"
	"github.com/olebedev/emitter"
)

// Adapter is the repository Adapter
type Adapter struct {
	adaptee *Driver
	event.Dispatcher
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
	adapter.Dispatcher = event.NewDispatcher(0)
	adapter.Emitter().Use("*", emitter.Void)

	return adapter, nil
}

// Ping ...
func (a *Adapter) Ping() error {
	return a.adaptee.session.Ping()
}

// Create is ...
func (a *Adapter) Create(res *resource.Resource) error {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("create", res)

	dataResource := a.toDoc(res)
	return (*a.adaptee).Create(dataResource)
}

// Get is ...
func (a *Adapter) Get(resType *core.ResourceType, id, version string, fields map[attr.Path]bool) (*resource.Resource, error) {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("get", resType, id, version, fields)

	// Setup query
	q, close, err := (*a.adaptee).Find(makeQuery(resType.GetIdentifier(), id, version))
	defer close()
	if err != nil {
		return nil, err
	}

	// Set projection
	query := Query{q: q}
	query.Fields(fields)

	// Make new document
	d := document{}

	// Query
	err = q.One(&d)
	if err != nil {
		return nil, err
	}

	return toResource(&d), nil
}

// Update is ...
func (a *Adapter) Update(resource *resource.Resource, id string, version string) error {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("update", resource, id, version)

	// We need to perform mutability validation
	// 1. Attributes whose mutability is "readWrite" that are omitted from
	// the request body MAY be assumed to be not asserted by the client.

	// 2. (Immutable attributes) If one or more values are already set for the attribute,
	// the input value(s) MUST match, or HTTP status code 400 SHOULD be
	// returned with a "scimType" error code of "mutability".

	// 3. (ReadOnly) Any values provided SHALL be ignored. (performed by the client)

	/* storedResource, err := a.Get(resource.ResourceType(), id, version, nil)
	if err != nil {
		return err
	}
	*/

	// (todo) => Test immutable attributes
	ro, err := attr.Paths(resource.ResourceType(), func(attribute *core.Attribute) bool {
		return attribute.Mutability == schemas.MutabilityImmutable

	})
	if err != nil {
		return err
	}

	if len(ro) > 0 {
		// we need to get stored value of immutable attributes
		sr, err := a.Get(resource.ResourceType(), id, version, nil)
		if err != nil {
			return err
		}

		for _, p := range ro {
			if p.Context(sr.ResourceType()).Get(sr) != p.Context(resource.ResourceType()).Get(resource) {
				return MutabilityError{
					msg: fmt.Sprintf("%s is immutable", p.String()),
				}
			}
		}
	}

	dataResource := a.toDoc(resource)
	return (*a.adaptee).Update(makeQuery(resource.ResourceType().GetIdentifier(), id, version), dataResource)
}

// Delete is ...
func (a *Adapter) Delete(resType *core.ResourceType, id, version string) error {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("delete", resType, id, version)

	return (*a.adaptee).Delete(makeQuery(resType.GetIdentifier(), id, version))
}

// Find is ...
func (a *Adapter) Find(resTypes []*core.ResourceType, filter filter.Filter) (storage.Querier, error) {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("find", resTypes, filter)

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
	if err != nil {
		close()
		return nil, err
	}
	return &Query{
		q: query,
		c: close,
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
		"schemas": r.Schemas,
		"id":      r.ID,
		"meta":    fromMeta(&r.Meta),
	}
	if r.ExternalID != "" {
		d["externalId"] = r.ExternalID
	}

	schemas, err := rt.GetSchemas()
	if err != nil {
		panic(err)
	}
	for ns := range schemas {
		if c := map[string]interface{}(*r.Values(ns)); c != nil {
			d[ns] = c
		}
	}
	return &d
}

func toResource(d *document) *resource.Resource {
	dd := (*d)

	// We are assuming here schemas, id, and meta will always be present here
	r := &resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas: toStringSlice(dd["schemas"].([]interface{})),
			ID:      dd["id"].(string),
			Meta:    toMeta(dd["meta"].(bson.M)),
		},
	}
	if dd["externalId"] != nil {
		r.CommonAttributes.ExternalID = dd["externalId"].(string)
	}

	rt := r.ResourceType()

	schemas, err := rt.GetSchemas()
	if err != nil {
		panic(err)
	}
	for ns, s := range schemas {
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
	meta := core.Meta{
		ResourceType: m["resourceType"].(string),
	}

	// version is not a required field, it is omitted if empty
	if m["version"] != nil {
		meta.Version = m["version"].(string)
	}

	if m["created"] != nil {
		created := m["created"].(time.Time)
		meta.Created = &created
	}

	if m["lastModified"] != nil {
		lastMod := m["lastModified"].(time.Time)
		meta.LastModified = &lastMod
	}

	if m["location"] != nil {
		meta.Location = m["location"].(string)
	}

	return meta
}

// MutabilityError is ...
// (note) => I have not yet figured out where to define the errors
// raised in a specific context that they will be wrapped in the API to avoid import cycling.
type MutabilityError struct {
	msg string
}

func (e MutabilityError) Error() string {
	return e.msg
}
