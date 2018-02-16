package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Scim2 defines the routes as per RFC 7644
//
//   POST 	/endpoint
//   GET  	/endpoint
//   POST 	/endpoint/.search
// 	 GET  	/endpoint/:id
//   PUT  	/endpoint/:id
//   PATCH  /endpoint/:id
//   DELETE	/endpoint/:id
func Scim2(group *gin.RouterGroup, resource Service) {
	endpoint := resource.Path()

	if res, ok := resource.(Poster); ok {
		group.POST(endpoint, res.Post)
	}

	if res, ok := resource.(Lister); ok {
		group.GET(endpoint, res.List)
	}

	if res, ok := resource.(Searcher); ok {
		group.POST(fmt.Sprintf("%s/%s", endpoint, ".search"), res.Search)
	}

	if res, ok := resource.(Getter); ok {
		group.GET(fmt.Sprintf("%s/:id", endpoint), res.Get)
	}

	if res, ok := resource.(Patcher); ok {
		group.PATCH(fmt.Sprintf("%s/:id", endpoint), res.Patch)
	}

	if res, ok := resource.(Putter); ok {
		group.PUT(fmt.Sprintf("%s/:id", endpoint), res.Put)
	}

	if res, ok := resource.(Deleter); ok {
		group.DELETE(fmt.Sprintf("%s/:id", endpoint), res.Delete)
	}
}
