package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fabbricadigitale/scimd/storage"
	"github.com/gin-gonic/gin"
	funk "github.com/thoas/go-funk"
)

var adapter storage.Storer

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

// Status is a gin middleware forcing the abortion of a request with the given code
func Status(code int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(code)
		return
	}
}

// Storage is a middleware to
func Storage(adapter storage.Storer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Set("storage", adapter)

		ctx.Next()
	}
}

// Set is a middleware to store a value by key within the context
func Set(key string, val interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(key, val)

		ctx.Next()
	}
}

// Authentication is a gin middleware supporting multiple authentication schemes
func Authentication(authenticationType string) gin.HandlerFunc {
	switch authenticationType {
	case strings.ToLower(HTTPBasic.String()):
		return gin.BasicAuth(gin.Accounts{
			"admin": "admin", // (todo) > get http basic auth credentials from config
		})
	default:
		panic("authentication scheme not available")
	}
}

// InternalServerError is a generic server error
type InternalServerError struct {
	msg string
}

func (e InternalServerError) Error() string {
	return fmt.Sprintf("%s", e.msg)
}
