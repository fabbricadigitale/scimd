package decorator

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/server"
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
func Scim2(group *gin.RouterGroup, resource server.Service) {
	endpoint := resource.Path()

	if res, ok := resource.(server.Poster); ok {
		group.POST(endpoint, res.Post)
	}

	if res, ok := resource.(server.Lister); ok {
		group.GET(endpoint, res.List)
	}

	if res, ok := resource.(server.Searcher); ok {
		group.POST(fmt.Sprintf("%s/%s", endpoint, ".search"), res.Search)
	}

	if res, ok := resource.(server.Getter); ok {
		group.GET(fmt.Sprintf("%s/:id", endpoint), res.Get)
	}

	if res, ok := resource.(server.Patcher); ok {
		group.PATCH(fmt.Sprintf("%s/:id", endpoint), res.Patch)
	}

	if res, ok := resource.(server.Putter); ok {
		group.PUT(fmt.Sprintf("%s/:id", endpoint), res.Put)
	}

	if res, ok := resource.(server.Deleter); ok {
		group.DELETE(fmt.Sprintf("%s/:id", endpoint), res.Delete)
	}
}
