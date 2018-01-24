package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cenk/backoff"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
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

// Storage is a middleware to
func Storage(endpoint, db, collection string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 3 * time.Minute

		err := backoff.Retry(func() error {
			var err error
			adapter, err = mongo.New(endpoint, db, collection)
			if err != nil {
				return err
			}

			return adapter.Ping()
		}, b)
		if err != nil {
			log.Fatalf("error after retrying: %v", err)
		}

		ctx.Set("storage", adapter)

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
