package messages

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
)

func TestListResponseResource(t *testing.T) {

	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.PushFromFile("../../internal/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.PushFromFile("../../internal/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.PushFromFile("../../internal/testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Unmarshal
	dat, err := ioutil.ReadFile("testdata/list.json")
	require.NotNil(t, dat)
	require.Nil(t, err)

	list := ListResponse{}
	err = json.Unmarshal(dat, &list)
	require.Equal(t, "urn:ietf:params:scim:api:messages:2.0:ListResponse", list.Schemas[0])
	require.Equal(t, 2, list.TotalResults)
	require.Equal(t, 3, list.ItemsPerPage)
	require.Equal(t, 1, list.StartIndex)

	res := resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
			ID:         "123",
			ExternalID: "456",
			Meta: core.Meta{
				Location:     "https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646",
				ResourceType: "User",
				Version:      "Wa330bc54f0671c9",
			},
		},
	}

	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"userName": datatype.String("Ale"),
	})

	require.Len(t, list.Resources, 1)
	require.Equal(t, res, *list.Resources[0])

	// Marshal
	list2 := ListResponse{}
	list2.Schemas = []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	list2.TotalResults = 6
	list2.ItemsPerPage = 5
	list2.StartIndex = 4
	res2 := resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
			ID:         "123",
			ExternalID: "456",
			Meta: core.Meta{
				Location:     "https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646",
				ResourceType: "User",
				Version:      "Wa330bc54f0671c9",
			},
		},
	}
	res2.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"userName": datatype.String("Sam"),
	})
	res2.SetValues("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", &datatype.Complex{
		"userName": datatype.String("Ale"),
	})
	list2.Resources = []*resource.Resource{
		&res2,
	}

	b, err2 := json.Marshal(&list2)

	require.NotNil(t, b)
	require.Nil(t, err2)

	require.JSONEq(t, `{
		"schemas": [
			"urn:ietf:params:scim:api:messages:2.0:ListResponse"
		],
		"totalResults": 6,
		"itemsPerPage": 5,
		"startIndex": 4,
		"Resources": [
			{
				"externalId": "456",
				"id": "123",
				"meta": {
					"location": "https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646",
					"resourceType": "User",
					"version": "Wa330bc54f0671c9"
				},
				"schemas": [
					"urn:ietf:params:scim:schemas:core:2.0:User"
				],
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User": {
					"userName": "Ale"
				},
				"userName": "Sam"
			}
		]
	}`, string(b))
}

func TestListResponseValid(t *testing.T) {
	var err error

	// Right ListResponse validation tags
	// Schema
	schema := []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	err = validation.Validator.Var(schema, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:ListResponse")
	require.NoError(t, err)

	// Totalresults
	results := 2
	err = validation.Validator.Var(results, "required")
	require.NoError(t, err)

	// ItemsPerPage
	items := 3
	err = validation.Validator.Var(items, "required")

	// StartIndex
	index := 1
	err = validation.Validator.Var(index, "gt=0,required")
	require.NoError(t, err)

	// Struct ListResponses
	testOk := ListResponse{}
	testOk.Schemas = []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	testOk.TotalResults = 2
	testOk.ItemsPerPage = 3
	testOk.StartIndex = 1
	testOk.Resources = []*resource.Resource{}

	err = validation.Validator.Struct(testOk)
	require.NoError(t, err)

	// Wrong ListResponse validation tags
	// Schema empty
	schema = []string{}
	err = validation.Validator.Var(schema, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:ListResponse")
	require.Error(t, err)
	// Random schema
	schema = []string{"randomurn"}
	err = validation.Validator.Var(schema, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:ListResponse")
	require.Error(t, err)

	// TotalResults empty
	results = 0
	err = validation.Validator.Var(results, "required")
	require.Error(t, err)

	// ItemsPerPage empty
	items = 0
	err = validation.Validator.Var(items, "required")
	require.Error(t, err)

	// StartIndex not greater than 0
	index = 0
	err = validation.Validator.Var(index, "gt=0,required")
	require.Error(t, err)

	// Bad struct ListResponse
	testWrong := ListResponse{}
	testWrong.Schemas = []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"}
	testWrong.TotalResults = 1
	testWrong.ItemsPerPage = 3
	testWrong.StartIndex = 2

	err = validation.Validator.Struct(testWrong)
	require.Error(t, err)
}
