package storage

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/core/resource"
	"github.com/stretchr/testify/require"
)

func TestCreateRepository(t *testing.T) {

	var manager Manager
	_, err := manager.CreateAdapter("mongo", "mongodb://localhost:27017/test_db?maxPoolSize=100", "test_db", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestCreate(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../schemas/core/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../schemas/core/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../schemas/core/testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("../schemas/core/resource/testdata/enterprise_user_resource.json")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.NotNil(t, dat)
	require.Nil(t, err)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var manager Manager
	adapter, err := manager.CreateAdapter("mongo", "mongodb://localhost:27017", "test_db", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = adapter.Create(res)

	if err != nil {
		t.Log(err)
	}

	require.Nil(t, err)
}

// (TODO) > Test hydrateResource adapter method

// (TODO) > Test toResource adapter method

// (TODO) > Test Get adapter method

// (TODO) > Test Delete adapter method
