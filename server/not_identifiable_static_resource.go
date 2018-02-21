package server

import (
	"net/http"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/gin-gonic/gin"
)

// NotIdentifiableStaticResource describers ...
type NotIdentifiableStaticResource struct {
	resource interface{}
	endpoint string
	Service
	Lister
}

// NewNotIdentifiableStaticResourceService creates a new `NotIdentifiableStaticResource` given a path and a resource
func NewNotIdentifiableStaticResourceService(endpoint string, resource interface{}) *NotIdentifiableStaticResource {
	_, identifiable := resource.(core.Identifiable)
	if identifiable {
		panic("can not create a not identifiable service for an identifiable resource")
	}

	return &NotIdentifiableStaticResource{
		endpoint: endpoint,
		resource: resource,
	}
}

// Path returns the endpoint of the `NotIdentifiableStaticResource`
func (rs *NotIdentifiableStaticResource) Path() string {
	return rs.endpoint
}

// List ...
func (rs *NotIdentifiableStaticResource) List(c *gin.Context) {
	// (note) > this is anomalous list, we do not want a ListResponse in response here
	c.JSON(http.StatusOK, rs.resource)
}
