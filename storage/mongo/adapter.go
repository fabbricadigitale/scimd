package mongo

import (
	"time"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/event"
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

// SetIndexes is ...
func (a *Adapter) SetIndexes(keys [][]string) error {

	ret := make([][]string, 0)

	for _, composedKeys := range keys {
		escapedKey := make([]string, 0)
		for _, singleKey := range composedKeys {
			escapedKey = append(escapedKey, escapeAttribute(singleKey))
		}
		ret = append(ret, escapedKey)
	}
	return a.adaptee.SetIndexes(ret)
}

// Ping ...
func (a *Adapter) Ping() error {
	return a.adaptee.session.Ping()
}

// Create is ...
func (a *Adapter) Create(res *resource.Resource) error {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("create", res, a)

	dataResource := a.toDoc(res)
	return (*a.adaptee).Create(dataResource)

}

// Close is the method to explicitly call to close the session
func (a *Adapter) Close() {
	a.adaptee.Close()
}

// Get is ...
func (a *Adapter) Get(resType *core.ResourceType, id, version string, fields map[attr.Path]bool) (*resource.Resource, error) {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("get", resType, id, version, fields)

	return a.DoGet(resType, id, version, fields)
}

// DoGet is ... Does not emit any event.
// Is used internally to fetch referenced documents.
func (a *Adapter) DoGet(resType *core.ResourceType, id, version string, fields map[attr.Path]bool) (*resource.Resource, error) {

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
	a.Emitter().Emit("update", resource, a)

	return a.DoUpdate(resource, id, version)
}

// DoUpdate is the method that permforms update operations. Does not emit any event.
// Is used internally to perform update to referenced documents.
func (a *Adapter) DoUpdate(resource *resource.Resource, id, version string) error {
	dataResource := a.toDoc(resource)
	return (*a.adaptee).Update(makeQuery(resource.ResourceType().GetIdentifier(), id, version), dataResource)
}

// Patch is ...
func (a *Adapter) Patch(resType *core.ResourceType, id string, version string, op string, f interface{}, value interface{}) error {

	p := &storage.PContainer{
		Value: value,
	}

	c := convert{}
	path := c.extractPath(f)

	var mResource *resource.Resource // resource before the update
	if path.Name == "password" {
		a.Emitter().Emit("patchPassword", p)
	}

	mResource, _ = a.Get(resType, id, version, nil)

	q, err := makeComplexQuery(resType, id, version, f)
	if err != nil {
		return err
	}

	v, err := convertChangeValue(resType, op, path, (*p).Value)
	if err != nil {
		return err
	}

	err = (*a.adaptee).Patch(id, q, v)

	if err == nil {
		// Get the current resource
		// Persist the update in related resources' readOnly attributes
		resource, err := a.Get(resType, id, version, nil)
		if err != nil {
			return err
		}
		a.Emitter().Emit("patch", mResource, resource, a)

	}

	return err
}

// Delete is ...
func (a *Adapter) Delete(resType *core.ResourceType, id, version string) error {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("delete", resType, id, version, a)

	return (*a.adaptee).Delete(makeQuery(resType.GetIdentifier(), id, version))
}

// Find is ...
func (a *Adapter) Find(resTypes []*core.ResourceType, filter filter.Filter) (storage.Querier, error) {
	// Emit an event and wait it has been sent successfully
	a.Emitter().Emit("find", resTypes, filter)

	or := make([]bson.M, len(resTypes))

	for i, resType := range resTypes {
		var err error
		or[i], err = convertToMongoQuery(resType, filter, "meta.resourceType", resType.GetIdentifier())
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
