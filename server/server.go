package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// MethodNotImplemented ...
func MethodNotImplemented(methods []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Abort incoming request if its method is not one of the supported
		if contains(methods, ctx.Request.Method) {
			ctx.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		// Otherwise go ahead
		ctx.Next()
	}
}
