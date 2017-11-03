package schemas

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

type repository struct {
	serviceProviderConfig core.ServiceProviderConfig
	schemas               map[string]core.Schema
	resourceTypes         map[string]core.ResourceType
	mu                    sync.RWMutex
}

func loadData(filename string, data interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &data)
}

func (r *repository) GetServiceProviderConfig() *core.ServiceProviderConfig {
	return &r.serviceProviderConfig
}

func (r *repository) LoadServiceProviderConfig(filename string) error {

	var config core.ServiceProviderConfig
	if err := loadData(filename, &config); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.serviceProviderConfig = config

	return nil
}

func (r *repository) GetSchema(ID string) *core.Schema {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if schema, ok := r.schemas[ID]; ok {
		return &schema
	}
	return nil
}

func (r *repository) LoadSchema(filename string) error {

	var schema core.Schema
	if err := loadData(filename, &schema); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.schemas[schema.ID] = schema

	return nil
}

func (r *repository) GetResourceType(name string) *core.ResourceType {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if resourceType, ok := r.resourceTypes[name]; ok {
		return &resourceType
	}
	return nil
}

func (r *repository) LoadResourceType(filename string) error {

	var resourceType core.ResourceType
	if err := loadData(filename, &resourceType); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.resourceTypes[resourceType.Name] = resourceType

	return nil
}

func createRepository() *repository {
	return &repository{
		serviceProviderConfig: core.ServiceProviderConfig{},
		schemas:               make(map[string]core.Schema),
		resourceTypes:         make(map[string]core.ResourceType),
	}
}

var (
	r    *repository
	once sync.Once
)

// Singleton repo for core schemas
func Repository() *repository {
	once.Do(func() {
		r = createRepository()
	})

	return r
}
