package server

import (
	"log"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/gin-gonic/gin"
)

// ResourceService describes ...
type ResourceService struct {
	rt *core.ResourceType
	Service
	Lister
	Getter
	Putter
	Patcher
	Deleter
	Searcher
}

// NewResourceService creates a new `ResourceService` for the given `core.ResourceTyper`
func NewResourceService(rt *core.ResourceType) *ResourceService {
	return &ResourceService{
		rt: rt,
	}
}

// Path returns the endpoint of the `ResourceService`
func (rs *ResourceService) Path() string {
	return rs.rt.Endpoint
}

// List ...
func (rs *ResourceService) List(c *gin.Context) {
	params := api.NewSearch()
	// Using the form binding engine (query)
	if err := c.ShouldBindQuery(params); err != nil {
		// (todo) > throw 4XX
		panic(err)
	}

	// Go ahead ...
	params.Attributes.Explode()
	log.Printf("%+v\n", params)
}

// Search ...
func (rs *ResourceService) Search(c *gin.Context) {
	contents := &messages.SearchRequest{}
	if err := c.ShouldBindJSON(contents); err != nil {
		// (todo)> throw 4XX
		panic(err)
	}

	// Go ahead ...
	log.Printf("%+v\n", contents)
}

// Get ...
func (rs *ResourceService) Get(c *gin.Context) {
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

// Post ...
func (rs *ResourceService) Post(*gin.Context) {

}

// Put ...
func (rs *ResourceService) Put(*gin.Context) {

}

// Patch ...
func (rs *ResourceService) Patch(*gin.Context) {

}

// Delete ...
func (rs *ResourceService) Delete(*gin.Context) {

}
