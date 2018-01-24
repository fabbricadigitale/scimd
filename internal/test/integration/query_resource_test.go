package integration

import (
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/api"
	"encoding/json"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"io/ioutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResource(t *testing.T) {

	
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	id := "2819c223-7f76-453a-919d-ab1234567891"

	dat, err := ioutil.ReadFile("../../testdata/user.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &core.ResourceType{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	attrs := &api.Attributes{}

	r, err := query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r)
	require.NoError(t, err)

	retRes := r.(*resource.Resource)
	values := retRes.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	userName := (*values)["userName"]

	require.NotNil(t, userName)

	// Excluding attribute
	// attrs.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:name"}
	// r2, err2 := query.Resource(adapter, res, id, attrs)
	// require.NotNil(t, r2)
	// require.NoError(t, err2)

	// retRes2 := r2.(*resource.Resource)
	// values2 := retRes2.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	// nameValue := (*values2)["name"].(*datatype.Complex)
	// fmt.Println(nameValue)
	// require.Nil(t, nameValue)

	// Excluding attribute's subattribute
	attrs.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:emails.value"}

	r, err = query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r)
	require.NoError(t, err)

	values = r.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	emails := (*values)["emails"]
	for _, email := range emails.([]datatype.DataTyper) {
		e := email.(*datatype.Complex)
		require.Nil(t, (*e)["value"])
	}
	


	return

	// Fail test, non existing id
	id = "wrong-id"
	r, err = query.Resource(adapter, res, id, attrs)
	require.Nil(t, r)
	require.Error(t, err)
}

func TestResources(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	dat, err := ioutil.ReadFile("../../testdata/user.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	resType := &core.ResourceType{}
	err = json.Unmarshal(dat, resType)
	require.NoError(t, err)

	resTypes := []*core.ResourceType{resType}
	require.NotNil(t, resTypes)
	
	// Filtering by attribute
	search := &api.Search{}
	search.Filter = `userName eq "tfork@example.com"`

	r, err := query.Resources(adapter, resTypes, search)
	require.NoError(t, err)
	require.NotEmpty(t, r.TotalResults)
 
	// Filtering by attribute's subattribute
	search.Filter = `name.middleName eq "Geraldine"`
	r, err = query.Resources(adapter, resTypes, search)
	require.NoError(t, err)
	require.NotEmpty(t, r.TotalResults)

	// Filtering by non existing attribute name
	search.Filter = `nonexistingattr eq "Geraldine"`
	r, err = query.Resources(adapter, resTypes, search)
	require.NoError(t, err)
	require.Empty(t, r.TotalResults)

	// Filtering by non existing attribute value
	search.Filter = `name.middleName eq "non-existing"`
	r, err = query.Resources(adapter, resTypes, search)
	require.NoError(t, err)
	require.Empty(t, r.TotalResults)
}
