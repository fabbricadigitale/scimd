package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func insensitiveContains(s []string, e string) bool {
	e = strings.ToLower(e)
	for _, a := range s {
		a = strings.ToLower(a)
		if a == e {
			return true
		}
	}
	return false
}

// MethodNotImplemented ...
func MethodNotImplemented(notSupportedMethods []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m := ctx.Request.Method
		// Abort incoming request if its method is not supported
		if insensitiveContains(notSupportedMethods, m) {
			ctx.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		// Otherwise go ahead
		ctx.Next()
	}
}
