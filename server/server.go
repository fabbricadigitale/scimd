package server

import (
	"fmt"
	"net/http"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage/listeners"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/gin-gonic/gin"
)

// Get setups endpoints as dictated by RFC 7644
//
// Details at https://tools.ietf.org/html/rfc7644#section-3.2
func Get(spc *core.ServiceProviderConfig) *gin.Engine {
	const (
		svcpcfgEndpoint = "/ServiceProviderConfigs"
		restypeEndpoint = "/ResourceTypes"
		schemasEndpoint = "/Schemas"
		bulkEndpoint    = "/Bulk"
		selfEndpoint    = "/Me"
		searchAction    = ".search"
	)

	if !config.Values.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Obtain an engine
	router := gin.Default()

	resTypeRepo := core.GetResourceTypeRepository()
	schemasRepo := core.GetSchemaRepository()

	// Root group
	v2 := router.Group("/v2")

	// Retrieve list of resource types
	resourceTypes := resTypeRepo.List()
	// Retrieve list of schemas
	schemas := schemasRepo.List()

	// (todo) > switch endpoint by config.Values.Storage.Type
	endpoint := fmt.Sprintf("%s:%d", config.Values.Storage.Host, config.Values.Storage.Port)
	adapter, err := mongo.New(endpoint, config.Values.Storage.Name, config.Values.Storage.Coll)
	if err != nil {
		panic(err)
	}
	listeners.AddListeners(adapter.Emitter())

	v2.Use(Storage(adapter))

	unsupportedMethods := []string{}
	if !spc.Patch.Supported {
		unsupportedMethods = append(unsupportedMethods, http.MethodPatch)
	}
	v2.Use(MethodNotImplemented(unsupportedMethods))

	for _, authScheme := range spc.AuthenticationSchemes {
		v2.Use(Authentication(authScheme.Type))
	}

	// Retrieve supported resource types
	Scim2(v2, NewStaticResourceService(restypeEndpoint, resourceTypes))

	// Retrieve one or more supported schemas
	Scim2(v2, NewStaticResourceService(schemasEndpoint, schemas))

	// Retrieve the ServiceProviderConfig
	Scim2(v2, NewNotIdentifiableStaticResourceService(svcpcfgEndpoint, config.ServiceProviderConfig()))

	// Bulk updates to one or more supported schemas
	if spc.Bulk.Supported {
		v2.POST(bulkEndpoint, bulking)
	}

	// (todo) > search from system root for one or more resource types using POST
	// v2.POST(fmt.Sprintf("/%s", searchAction), searching)

	// Create endpoints for all resource types
	for _, rt := range resourceTypes {
		Scim2(v2, NewResourceService(rt))
	}

	// Alias for operations against a resource mapped to an authenticated subject
	me := v2.Group(selfEndpoint)
	if !config.Values.Enable.Self {
		// RFC 7644 - Section 3.11 - 1st bullet
		me.Use(Status(http.StatusNotImplemented))
	}
	if self := resTypeRepo.Pull("User"); self != nil {
		self.Endpoint = ""
		Scim2(me, NewResourceService(*self))
	}

	return router
}

// (fixme)
func bulking(c *gin.Context) {

}
