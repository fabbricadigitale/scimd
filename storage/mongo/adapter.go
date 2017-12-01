package mongo

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// Adapter is the repository Adapter
type Adapter struct {
	adaptee *Driver
}

// (fixme) var _ storage.Storer = (*Adapter)(nil)
// (fixme) global adapter must be avoided
var adapter Adapter

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
	return (*a.adaptee).Search()
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
