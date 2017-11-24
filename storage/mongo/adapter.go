package mongo

import (
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

	dataResource := a.toData(res)

	return (*a.adaptee).Create(&dataResource)
}

// Get is ...
func (a *Adapter) Get(id, version string) error {
	return (*a.adaptee).Get(id, version)
}

// Count ...
func (a *Adapter) Count() error {
	return (*a.adaptee).Count()
}

// Update is ...
func (a *Adapter) Update() error {
	return (*a.adaptee).Update()
}

// Delete is ...
func (a *Adapter) Delete(id, version string) error {
	return (*a.adaptee).Delete(id, version)
}

// Search is ...
func (a *Adapter) Search() error {
	return (*a.adaptee).Search()
}

type Data map[string]interface{}

func (a *Adapter) toData(r *resource.Resource) Data {

	m := Data{}

	m["schemas"] = r.Common.Schemas
	m["id"] = r.Common.ID
	m["external_id"] = r.Common.ExternalID
	m["meta"] = r.Common.Meta

	for key, val := range *r.GetValues(r.GetSchema().ID) {
		m[key] = val
	}

	for _, extSch := range r.GetSchemaExtensions() {
		if extSch != nil {
			ns := extSch.GetIdentifier()
			values := *r.GetValues(ns)
			m[ns] = values
		}

	}

	return m
}
