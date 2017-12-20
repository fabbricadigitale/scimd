// +build integration

package integration

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/stretchr/testify/require"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository
var adapter storage.Storer

func init() {
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.Add("../../testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.Add("../../testdata/user_schema.json")
	schemaRepo.Add("../../testdata/enterprise_user_schema.json")

	adapter, _ = mongo.New("mongodb://localhost:27017", "scimd", "resources")
}

func TestMongoCreate(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("../../testdata/enterprise_user_resource_1.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	err = adapter.Create(res)
	require.NoError(t, err)
}

func TestMongoGet(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	id := "2819c223-7f76-453a-919d-ab1234567891"
	resource, err := adapter.Get(resTypeRepo.Get("User"), id, "", nil, nil)
	require.NoError(t, err)

	require.NotNil(t, resource)
	require.Equal(t, id, resource.ID)
}

// (todo) > Test hydrateResource adapter method

// (todo) > Test toResource adapter method

// (todo) > Test Get adapter method

// (todo) > Test Delete adapter method

// (todo) > Test Update adapter method
