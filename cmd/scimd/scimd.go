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

func setup() *gin.Engine {
	const (
		svcpcfgEndpoint = "/ServiceProviderConfigs"
		restypeEndpoint = "/ResourceTypes"
		schemasEndpoint = "/Schemas"
		bulkEndpoint    = "/Bulk"
		selfEndpoint    = "/Me"
		searchAction    = ".search"
	)

	spc := config()
	router := gin.Default()

	resTypeRepo := core.GetResourceTypeRepository()
	schemasRepo := core.GetSchemaRepository()

	// Setup endpoint as dictated by https://tools.ietf.org/html/rfc7644#section-3.2
	v2 := router.Group("/v2")

	//v2.Use(server.Storage(dbURL, dbName, dbCollection))

	resourceTypes := resTypeRepo.List()
	schemas := schemasRepo.List()

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

	// Alias for operations against a resource mapped to an authenticated subject
	v2.GET(selfEndpoint, selfing)
	v2.POST(selfEndpoint, selfing)
	v2.PUT(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.PATCH(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.DELETE(fmt.Sprintf("%s/:id", selfEndpoint), selfing)

	// Search from system root for one or more resource types using POST
	v2.POST(fmt.Sprintf("/%s", searchAction), searching)

	for _, rt := range resourceTypes {
		decorator.Scim2(v2, server.NewResourceService(&rt))
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

func posting(c *gin.Context) {

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
