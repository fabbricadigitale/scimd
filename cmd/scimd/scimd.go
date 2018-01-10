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

	v2.GET(svcpcfgEndpoint, getting)
	v2.GET(restypeEndpoint, listing)
	v2.GET(schemasEndpoint, listing)
	v2.POST(bulkEndpoint, bulking)
	v2.GET(selfEndpoint, selfing)
	v2.POST(selfEndpoint, selfing)
	v2.PUT(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.PATCH(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.DELETE(fmt.Sprintf("%s/:id", selfEndpoint), selfing)
	v2.POST(fmt.Sprintf("/%s", searchAction), getting)

	for _, rt := range core.GetResourceTypeRepository().List() {
		v2.GET(fmt.Sprintf("%s/%s", restypeEndpoint, rt.GetIdentifier()), getting)
		v2.GET(rt.Endpoint, listing)
		v2.POST(rt.Endpoint, posting)
		v2.POST(fmt.Sprintf("%s/%s", rt.Endpoint, searchAction), getting)

		v2.GET(fmt.Sprintf("%s/:id", rt.Endpoint), getting)
		v2.PUT(fmt.Sprintf("%s/:id", rt.Endpoint), putting)
		v2.PATCH(fmt.Sprintf("%s/:id", rt.Endpoint), patching)
		v2.DELETE(fmt.Sprintf("%s/:id", rt.Endpoint), deleting)
	}

	// (fixme) > schema IDs must be escaped
	//	for _, sc := range core.GetSchemaRepository().List() {
	//		v2.GET(fmt.Sprintf("%s/%s", schemasEndpoint, sc.GetIdentifier()), getting)
	//	}

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
