package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cenk/backoff"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/listeners"
	"github.com/fabbricadigitale/scimd/storage/mongo"
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
func Storage(endpoint, db, collection string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 1 * time.Millisecond

		err := backoff.Retry(func() error {
			var err error
			adapter, err = mongo.New(endpoint, db, collection)
			if err != nil {
				return err
			}
			listeners.AddListeners(adapter.Emitter())

			return adapter.Ping()
		}, b)

		if err != nil {
			if config.Values.Debug {
				log.Printf("error after retrying: %v", err)
			}

			e := messages.NewError(InternalServerError{
				msg: fmt.Sprintf("error after retrying: %v", err),
			})
			ctx.JSON(e.Status, e)
			ctx.Abort()
		}

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
		// (todo) > fetch http basic auth credentials from config
		return gin.BasicAuth(gin.Accounts{
			"admin": "admin", // (todo) > from config ?
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
