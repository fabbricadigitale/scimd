package harness

import (

	"io/ioutil"
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

func TestGetServiceProviderConfig(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v2/ServiceProviderConfigs", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	var exp []byte
	if exp, err = json.Marshal(spc); err != nil {
		t.Fatalf("%s", err)
	}
	require.Equal(t, string(exp), rec.Body.String())
}

func TestGetWithoutInclusions(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users/2819c223-7f76-453a-919d-ab1234567891", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)
	exp, _ := ioutil.ReadFile("../../testdata/resp_user_full_attributes.json")

	require.JSONEq(t, string(exp), rec.Body.String())
}

func TestGetWithExistingAttributes(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users/2819c223-7f76-453a-919d-ab1234567891?attributes=displayname,name.givenname", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	resp := rec.Body.String()

	usr, _ := ioutil.ReadFile("../../testdata/resp_user_existing_attributes.json")

	require.JSONEq(t, string(usr), resp)
}


func TestListUsers(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	// (todo) > Add expected
	// exp := ...
	// require.Equal(t, , rec.Body.String())

	// require.Equal(t, 2, len()) // (todo) > check response have 2 users
	fmt.Println(rec.Body.String()) // (todo) > test returns a list response containing them
}

// PAGINATION
