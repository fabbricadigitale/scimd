package mongo

import (
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/core/resource"
)

// Adapter is the repository Adapter
type Adapter struct {
	adaptee *Driver
}

var adapter Adapter

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
func (a *Adapter) Get(id, version string) (*resource.Resource, error) {

	h := &HResource{}

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
func (a *Adapter) Update(id string, version string, resource *resource.Resource) error {
	dataResource := a.hydrateResource(resource)
	return (*a.adaptee).Update(id, version, dataResource)
}

// Delete is ...
func (a *Adapter) Delete(id, version string) error {
	return (*a.adaptee).Delete(id, version)
}

// Search is ...
func (a *Adapter) Search() error {
	return (*a.adaptee).Search()
}

// HResource is a ready-to-store format for Resource
type HResource struct {
	Data []map[string]interface{}
}

func (a *Adapter) hydrateResource(r *resource.Resource) *HResource {

	h := &HResource{}

	common := make(map[string]interface{})
	common["_urn"] = "common"
	common["schemas"] = r.Common.Schemas
	common["id"] = r.Common.ID
	common["external_id"] = r.Common.ExternalID
	common["meta"] = r.Common.Meta

	mCore := make(map[string]interface{})
	mCore["_urn"] = r.GetSchema().ID
	for key, val := range *r.GetValues(r.GetSchema().ID) {
		mCore[key] = val
	}
	h.Data = append(h.Data, common, mCore)

	for _, extSch := range r.GetSchemaExtensions() {
		mExt := make(map[string]interface{})
		if extSch != nil {
			ns := extSch.GetIdentifier()
			mExt["_urn"] = ns
			for key, val := range *r.GetValues(ns) {
				mExt[key] = val
			}
		}
		h.Data = append(h.Data, mExt)
	}

	return h
}

func (a *Adapter) toResource(h *HResource) (*resource.Resource, error) {

	r := &resource.Resource{}

	c := core.Common{}
	hCommon := h.Data[0]

	c.Schemas = hCommon["schemas"].([]string)
	c.ID = hCommon["id"].(string)
	c.ExternalID = hCommon["external_id"].(string)
	c.Meta = hCommon["meta"].(core.Meta)

	r.Common = c

	var p *core.Complex
	for i := 1; i < len(h.Data); i++ {
		ns := h.Data[i]["_urn"].(string)
		values := h.Data[i]
		delete(values, "_urn")
		(*p) = core.Complex(values)
		r.SetValues(ns, p)
	}

	return r, nil
}
