package harness

import (
	"github.com/fabbricadigitale/scimd/api/messages"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var list messages.ListResponse
	json.Unmarshal([]byte(rec.Body.String()), &list)

	listResponseURN := []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	require.Equal(t, 2, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)
	require.Equal(t, 1, list.StartIndex)
	require.Equal(t, 2, len(list.Resources))
	require.Equal(t, listResponseURN, list.Schemas)

	respUsrSchema := list.Resources[0]
	actSchema, _ := json.Marshal(respUsrSchema)

	respGroupSchema := list.Resources[1]
	actGroup, _ := json.Marshal(respGroupSchema)

	expUserSchema := core.GetSchemaRepository().Pull("urn:ietf:params:scim:schemas:core:2.0:User")
	expSchema, _ := json.Marshal(expUserSchema)

	expGroupSchema := core.GetSchemaRepository().Pull("urn:ietf:params:scim:schemas:core:2.0:Group")
	expGroup, _ := json.Marshal(expGroupSchema)

	require.JSONEq(t, string(expSchema), string(actSchema))
	require.JSONEq(t, string(expGroup), string(actGroup))
}

func TestGetResourceType(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)

	for _, rt := range core.GetResourceTypeRepository().List() {

		fmt.Printf("%T %+v", rt, rt)

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

	var list messages.ListResponse
	json.Unmarshal([]byte(rec.Body.String()), &list)

	listResponseURN := []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	require.Equal(t, 2, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)
	require.Equal(t, 1, list.StartIndex)
	require.Equal(t, 2, len(list.Resources))
	require.Equal(t, listResponseURN, list.Schemas)

	respUserResType := list.Resources[0]
	actUserResType, _ := json.Marshal(respUserResType)

	respGroupResType := list.Resources[1]
	actGroupResType, _ := json.Marshal(respGroupResType)

	expectedUserResType := core.GetResourceTypeRepository().Pull("User")
	expUserResType, _ := json.Marshal(expectedUserResType)

	expectedGroupResType := core.GetResourceTypeRepository().Pull("Group")
	expGroupResType, _ := json.Marshal(expectedGroupResType)

	require.JSONEq(t, string(expUserResType), string(actUserResType))
	require.JSONEq(t, string(expGroupResType), string(actGroupResType))
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

	var list messages.ListResponse
	json.Unmarshal([]byte(rec.Body.String()), &list)

	listResponseURN := []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	require.Equal(t, 2, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)
	require.Equal(t, 1, list.StartIndex)
	require.Equal(t, 2, len(list.Resources))
	require.Equal(t, listResponseURN, list.Schemas)

	respUser1 := list.Resources[0]
	actUser1, _ := json.Marshal(respUser1)	
	expUser1, _ := ioutil.ReadFile("../../testdata/resp_user_full_attributes.json")
	require.JSONEq(t, string(expUser1), string(actUser1))

	respUser2 := list.Resources[1]
	actUser2, _ := json.Marshal(respUser2)	
	expUser2, _ := ioutil.ReadFile("../../testdata/user_resource_2.json")
	require.JSONEq(t, string(expUser2), string(actUser2))
}

// PAGINATION
func TestListUsersWithPagination(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "2")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	list := &messages.ListResponse{}
	json.Unmarshal([]byte(rec.Body.String()), list)

	exp, _ := ioutil.ReadFile("../../testdata/user_resource_2.json")

	require.Equal(t, 1, len(list.Resources))
	require.Equal(t, 2, list.StartIndex)
	require.Equal(t, 1, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	res := list.Resources[0]
	act, _ := json.Marshal(res)

	require.JSONEq(t, string(exp), string(act))
}

func TestListUsersWithStartIndexAndCount(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "1")
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	list := &messages.ListResponse{}
	json.Unmarshal([]byte(rec.Body.String()), list)

	exp, _ := ioutil.ReadFile("../../testdata/resp_user_full_attributes.json")

	require.Equal(t, 1, len(list.Resources))
	require.Equal(t, 1, list.StartIndex)
	require.Equal(t, 1, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	res := list.Resources[0]
	act, _ := json.Marshal(res)

	require.JSONEq(t, string(exp), string(act))
}

func TestListUsersWrongPagination(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "-1")
	q.Add("count", "-1")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 500, rec.Code)

	// TODO => Why this doesn't work?
	/* byt := []byte(`{
		"schemas": [
			"urn:ietf:params:scim:api:messages:2.0:Error"
		],
		"status": 500,
		"detail": "Key: 'Search.Pagination.StartIndex' Error:Field validation for 'StartIndex' failed on the 'gt' tag"
	}`)
	require.JSONEq(t, string(byt), rec.Body.String()) */
}

func TestListUsersWithABiggerLimit(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "2")
	q.Add("count", "33")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	list := &messages.ListResponse{}
	json.Unmarshal([]byte(rec.Body.String()), list)

	exp, _ := ioutil.ReadFile("../../testdata/user_resource_3.json")

	require.Equal(t, 1, len(list.Resources))
	require.Equal(t, 2, list.StartIndex)
	require.Equal(t, 1, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	res := list.Resources[0]
	act, _ := json.Marshal(res)

	require.JSONEq(t, string(exp), string(act))
}

func TestListUsersWithABiggerStartIndex(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "33")
	q.Add("count", "2")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	list := &messages.ListResponse{}
	json.Unmarshal([]byte(rec.Body.String()), list)

	require.Equal(t, 0, len(list.Resources))
	require.Equal(t, 33, list.StartIndex)
	require.Equal(t, 0, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	res := list.Resources

	require.Empty(t, res)
}

func TestListSchemasWithPagination(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Schemas", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "2")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	var list messages.ListResponse
	json.Unmarshal([]byte(rec.Body.String()), &list)

	require.Equal(t, 1, len(list.Resources))
	require.Equal(t, 2, list.StartIndex)
	require.Equal(t, 1, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	repoGroupSchema := list.Resources[0]
	act, _ := json.Marshal(repoGroupSchema)

	expGroupSchema := core.GetSchemaRepository().Pull("urn:ietf:params:scim:schemas:core:2.0:Group")
	expGroup, _ := json.Marshal(expGroupSchema)

	require.JSONEq(t, string(expGroup), string(act))
}

func TestListSchemasWithStartIndexAndCount(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Schemas", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "1")
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 200, rec.Code)

	var list messages.ListResponse
	json.Unmarshal([]byte(rec.Body.String()), &list)

	require.Equal(t, 1, len(list.Resources))
	require.Equal(t, 1, list.StartIndex)
	require.Equal(t, 1, list.ItemsPerPage)
	require.Equal(t, 2, list.TotalResults)

	repoUserSchema := list.Resources[0]
	act, _ := json.Marshal(repoUserSchema)

	expUserSchema := core.GetSchemaRepository().Pull("urn:ietf:params:scim:schemas:core:2.0:User")
	expUser, _ := json.Marshal(expUserSchema)

	require.JSONEq(t, string(expUser), string(act))
}

func TestListSchemasWrongPagination(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v2/Schemas", nil)
	req.Header.Add("Authorization", aaa)

	q := req.URL.Query()
	q.Add("startIndex", "-1")
	q.Add("count", "-1")
	req.URL.RawQuery = q.Encode()

	srv.ServeHTTP(rec, req)

	require.Equal(t, 500, rec.Code)

	// TODO => Why this doesn't work?
	/* byt := []byte(`{
		"schemas": [
			"urn:ietf:params:scim:api:messages:2.0:Error"
		],
		"status": 500,
		"detail": "Key: 'Search.Pagination.StartIndex' Error:Field validation for 'StartIndex' failed on the 'gt' tag"
	}`)
	require.JSONEq(t, string(byt), rec.Body.String()) */
}
