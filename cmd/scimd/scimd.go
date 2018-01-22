package main

import (
	"fmt"
	"net/http"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/gin-gonic/gin"
)

func main() {
	setup().Run(":8787")
}

func setup() *gin.Engine {
	const (
		svcpcfgEndpoint = "/ServiceProviderConfig"
		restypeEndpoint = "/ResourceTypes"
		schemasEndpoint = "/Schemas"
		bulkEndpoint    = "/Bulk"
		selfEndpoint    = "/Me"
		searchAction    = ".search"
	)

	spc := config()
	router := gin.Default()

	// Setup endpoint as dictated by https://tools.ietf.org/html/rfc7644#section-3.2
	v2 := router.Group("/v2")

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
	v2.GET(restypeEndpoint, listing)
	// Retrieve one or more supported schemas
	v2.GET(schemasEndpoint, listing)
	v2.GET(fmt.Sprintf("%s/:id", schemasEndpoint), getting)

	// Bulk updates to one or more supported schemas
	if spc.Bulk.Supported {
		v2.POST(bulkEndpoint, bulking)
	}

	// Alias for operations against a resource mapped to an authenticated subject
	v2.GET(selfEndpoint, selfing)
	v2.POST(selfEndpoint, selfing)
	v2.PUT(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.PATCH(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.DELETE(fmt.Sprintf("%s/:id", selfEndpoint), selfing)

	// Search from system root for one or more resource types using POST
	v2.POST(fmt.Sprintf("/%s", searchAction), getting)

	for _, rt := range core.GetResourceTypeRepository().List() {
		// (todo) > verify whether RFC specifies endpoint to retrieve resource type by identifier, or not
		// v2.GET(fmt.Sprintf("%s/%s", restypeEndpoint, rt.GetIdentifier()), getting)

		// List all resources
		v2.GET(rt.Endpoint, listing)

		// Create new resource
		v2.POST(rt.Endpoint, posting)

		// Search within a resource endpoint for one or more resource types using POST
		v2.POST(fmt.Sprintf("%s/%s", rt.Endpoint, searchAction), getting)

		// Retrieve, add, modify, or delete resource
		v2.GET(fmt.Sprintf("%s/:id", rt.Endpoint), getting)
		v2.PUT(fmt.Sprintf("%s/:id", rt.Endpoint), putting)
		v2.PATCH(fmt.Sprintf("%s/:id", rt.Endpoint), patching)
		v2.DELETE(fmt.Sprintf("%s/:id", rt.Endpoint), deleting)
	}

	return router
}

// query string via decorators?

func listing(c *gin.Context) {

}

func posting(c *gin.Context) {

}

func getting(c *gin.Context) {

}

func putting(c *gin.Context) {

}

func deleting(c *gin.Context) {

}

func selfing(c *gin.Context) {

}

func patching(c *gin.Context) {

}

func bulking(c *gin.Context) {

}

func resourcing(c *gin.Context) {

}
