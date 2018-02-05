package server

import (
	"log"
	"net/http"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/create"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/api/update"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
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

	// Retrieve the storage adapter
	store, ok := c.Get("storage")
	if !ok {
		panic("Missing storage setup ...")
	}

	resArr := make([]*core.ResourceType, 0)
	resArr[0] = rs.rt

	list, err := query.Resources(store.(storage.Storer), resArr, params)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, list)
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

	// Explode the attributes
	attrs.Explode()

	// Retrieve the storage adapter
	store, ok := c.Get("storage")
	if !ok {
		panic("Missing storage setup ...")
	}
	// Retrieve the id segment
	id := c.Param("id")

	res, err := query.Resource(store.(storage.Storer), rs.rt, id, &attrs)
	if err != nil {
		log.Println("(todo) > handle error")
	}

	c.JSON(http.StatusOK, res.(*resource.Resource))
}

// Post ...
func (rs *ResourceService) Post(c *gin.Context) {
	contents := &resource.Resource{}
	if err := c.ShouldBindJSON(contents); err != nil {
		// (todo)> throw 4XX
		panic(err)
	}

	// Retrieve the storage adapter
	store, ok := c.Get("storage")
	if !ok {
		panic("Missing storage setup ...")
	}

	create.Resource(store.(storage.Storer), rs.rt, contents)

	c.JSON(http.StatusOK, nil)

}

// Put ...
func (rs *ResourceService) Put(c *gin.Context) {
	var attrs api.Attributes
	// Using the form binding engine (query)
	if err := c.ShouldBindQuery(&attrs); err != nil {
		// (todo)> throw 4XX
		panic(err)

	}

	// Explode the attributes
	attrs.Explode()
	// Retrieve the id segment
	id := c.Param("id")

	contents := &resource.Resource{}
	if err := c.ShouldBindJSON(contents); err != nil {
		// (todo)> throw 4XX
		panic(err)
	}

	// Retrieve the storage adapter
	store, ok := c.Get("storage")
	if !ok {
		panic("Missing storage setup ...")
	}
	update.Resource(store.(storage.Storer), id, contents)

	c.JSON(http.StatusOK, nil)

}

// Patch ...
func (rs *ResourceService) Patch(*gin.Context) {

}

// Delete ...
func (rs *ResourceService) Delete(*gin.Context) {

}
