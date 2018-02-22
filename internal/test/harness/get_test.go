package harness

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/server"
	"github.com/stretchr/testify/require"
)

var aaa = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("admin:admin")))

func TestGetSchema(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)

	for _, schema := range core.GetSchemaRepository().List() {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/v2/Schemas/%s", schema.GetIdentifier()), nil)
		req.Header.Add("Authorization", aaa)
		srv.ServeHTTP(rec, req)

		var exp []byte
		if exp, err = json.Marshal(schema); err != nil {
			t.Fatalf("%s", err)
		}
		require.Equal(t, string(exp), rec.Body.String())
	}
}

func TestListSchemas(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v2/Schemas", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	fmt.Println(rec.Body.String()) // (todo) > test returns a list response containing them
}

func TestGetResourceType(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)

	for _, rt := range core.GetResourceTypeRepository().List() {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/v2/ResourceTypes/%s", rt.GetIdentifier()), nil)
		req.Header.Add("Authorization", aaa)
		srv.ServeHTTP(rec, req)

		var exp []byte
		if exp, err = json.Marshal(rt); err != nil {
			t.Fatalf("%s", err)
		}
		require.Equal(t, string(exp), rec.Body.String())
	}
}

func TestListResourceTypes(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v2/ResourceTypes", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	fmt.Println(rec.Body.String()) // (todo) > test returns a list response containing them
}
