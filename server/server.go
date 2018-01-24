package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

// MethodNotImplemented is a gin middleware responsible to abort requests which method is not supported.
func MethodNotImplemented(notSupportedMethods []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		m := strings.ToLower(ctx.Request.Method)
		var methods []string
		funk.Map(notSupportedMethods, func(x string) string {
			return strings.ToLower(x)
		})

		// Abort incoming request if its method is not supported
		if funk.ContainsString(methods, m) {
			ctx.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		// Otherwise go ahead
		ctx.Next()
	}
}

// Authentication is a gin middleware supporting multiple authentication schemes
func Authentication(authenticationType string) gin.HandlerFunc {
	switch authenticationType {
	case strings.ToLower(HTTPBasic.String()):
		// (todo) > fetch http basic auth credentials from config
		return gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		})
	default:
		panic("authentication scheme not available")
	}
}
