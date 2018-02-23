package integration

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResource(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	id := "2819c223-7f76-453a-919d-ab1234567891"

	res := core.GetResourceTypeRepository().Pull("User")
	require.NotNil(t, res)

	attrs := &api.Attributes{}

	r, err := query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r)
	require.NoError(t, err)

	retRes := r.(*resource.Resource)
	values := retRes.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	userName := (*values)["userName"]
	require.NotNil(t, userName)

	// Test extensions
	extensions := retRes.Values("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
	department := (*extensions)["department"]

	require.NotNil(t, department)
	require.Equal(t, datatype.String("Tour Operations"), department)

	// Test that we do not support excluding attributes that have subattributes
	attrs.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:name"}
	r2, err2 := query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r2)
	require.NoError(t, err2)

	retRes2 := r2.(*resource.Resource)
	values2 := retRes2.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	nameValue := (*values2)["name"].(*datatype.Complex)
	require.NotNil(t, nameValue) // not excluded since "name" attribute has sub attributes

	// Test that we support excluding first level attributes that does not have subattributes
	attrs.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:displayName"}
	r3, err3 := query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r3)
	require.NoError(t, err3)

	retRes3 := r3.(*resource.Resource)
	values3 := retRes3.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	dNameValue := (*values3)["displayName"]
	require.Nil(t, dNameValue) // excluded since "displayName" is a first level attribute without subattributes

	// Excluding attribute's subattribute
	attrs.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:emails.value"}

	r, err = query.Resource(adapter, res, id, attrs)
	require.NotNil(t, r)
	require.NoError(t, err)

	values = r.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	emails := (*values)["emails"]

	for _, email := range emails.([]datatype.DataTyper) {
		e := email.(*datatype.Complex)
		assert.Nil(t, (*e)["value"])
	}

	// Fail test, non existing id
	id = "wrong-id"
	r2, err = query.Resource(adapter, res, id, attrs)
	require.Nil(t, r2)
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
	search.SortOrder = api.AscendingOrder
	search.StartIndex = 1

	// search.Filter = `userName eq "tfork@example.com"`
	search.Filter = `addresses.country eq "USA"`
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

	//
}
