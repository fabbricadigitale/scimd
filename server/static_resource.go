package server

import (
	"log"
	"net/http"
	"reflect"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/gin-gonic/gin"
)

// StaticResourceService describers ...
type StaticResourceService struct {
	resources map[string]interface{}
	endpoint  string
	Service
	Lister
	Getter
}

// NewStaticResourceService creates a new `StaticResourceService` given a path and a list of `core.ResourceTyper`
func NewStaticResourceService(endpoint string, resources interface{}) *StaticResourceService {
	m := make(map[string]interface{})
	rv := reflect.ValueOf(resources)

	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			res := rv.Index(i).Interface()
			m[res.(core.Identifiable).GetIdentifier()] = res
		}
	default:
		panic("not available...")
	}

	return &StaticResourceService{
		endpoint:  endpoint,
		resources: m,
	}
}

// Path returns the endpoint of the `StaticResourceService`
func (rs *StaticResourceService) Path() string {
	return rs.endpoint
}

// List ...
func (rs *StaticResourceService) List(c *gin.Context) {
	log.Println(rs.resources)
	// (fixme) > how to return a list response message (which contains resource.Resource(s)) ?
}

// Get ...
func (rs *StaticResourceService) Get(c *gin.Context) {
	id := c.Param("id")
	message := rs.resources[id]
	c.JSON(http.StatusOK, message)
}
