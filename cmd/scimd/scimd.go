package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/fabbricadigitale/scimd/server/decorator"
	"github.com/gin-gonic/gin"
)

func main() {
	setup().Run(":8787")
}

// Setup endpoints as dictated by https://tools.ietf.org/html/rfc7644#section-3.2
func setup() *gin.Engine {
	const (
		svcpcfgEndpoint = "/ServiceProviderConfigs"
		restypeEndpoint = "/ResourceTypes"
		schemasEndpoint = "/Schemas"
		bulkEndpoint    = "/Bulk"
		selfEndpoint    = "/Me"
		searchAction    = ".search"
	)

	// Initialize configurations
	spc := config()

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

	v2.Use(server.Storage(dbURL, dbName, dbCollection))

	unsupportedMethods := []string{}
	if !spc.Patch.Supported {
		unsupportedMethods = append(unsupportedMethods, http.MethodPatch)
	}
	v2.Use(server.MethodNotImplemented(unsupportedMethods))

	for _, authScheme := range spc.AuthenticationSchemes {
		v2.Use(server.Authentication(authScheme.Type))
	}

	// Retrieve service provider config
	v2.GET(svcpcfgEndpoint, getting)

	// Retrieve supported resource types
	decorator.Scim2(v2, server.NewStaticResourceService(restypeEndpoint, resourceTypes))

	// Retrieve one or more supported schemas
	decorator.Scim2(v2, server.NewStaticResourceService(schemasEndpoint, schemas))

	// Bulk updates to one or more supported schemas
	if spc.Bulk.Supported {
		v2.POST(bulkEndpoint, bulking)
	}

	// Search from system root for one or more resource types using POST
	v2.POST(fmt.Sprintf("/%s", searchAction), searching)

	// Create endpoints for all resource types
	for _, rt := range resourceTypes {
		decorator.Scim2(v2, server.NewResourceService(&rt))
	}

	// Alias for operations against a resource mapped to an authenticated subject
	const mountSelf = false
	me := v2.Group(selfEndpoint)
	if !mountSelf {
		// RFC 7644 - Section 3.11 - 1st bullet
		me.Use(server.Status(http.StatusNotImplemented))
	}
	if self := resTypeRepo.Get("User"); self != nil {
		self.Endpoint = ""
		decorator.Scim2(me, server.NewResourceService(self))
	}

	return router
}

func listing(c *gin.Context) {
	params := api.NewSearch()
	// Using the form binding engine (query)
	if err := c.ShouldBindQuery(params); err != nil {
		// (todo)> throw 4XX
		panic(err)
	}

	// Go ahead ...
	params.Attributes.Explode()
	log.Printf("%+v\n", params)
}

func searching(c *gin.Context) {
	contents := &messages.SearchRequest{}
	if err := c.ShouldBindJSON(contents); err != nil {
		// (todo)> throw 4XX
		panic(err)
	}

	// Go ahead ...
	log.Printf("%+v\n", contents)
}

func getting(c *gin.Context) {
	var attrs api.Attributes
	// Using the form binding engine (query)
	if err := c.ShouldBindQuery(&attrs); err != nil {
		// (todo)> throw 4XX
		panic(err)

	}

	// Go ahead ...
	attrs.Explode()
	log.Printf("%+v\n", attrs)

}

func bulking(c *gin.Context) {

}
