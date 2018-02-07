package integration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/olebedev/emitter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

	var called bool
	em := adapter.Emitter()
	em.On("create", func(evt *emitter.Event) {
		called = true
		assert.Equal(t, len(evt.Args), 1)
		assert.IsType(t, (*resource.Resource)(nil), evt.Args[0])
	})

	err = adapter.Create(res)

	assert.True(t, called)

	require.NoError(t, err)
}

func TestMongoGet(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	id := "2819c223-7f76-453a-919d-ab1234567891"
	resource, err := adapter.Get(resTypeRepo.Get("User"), id, "", nil)
	require.NoError(t, err)

	require.NotNil(t, resource)
	require.Equal(t, id, resource.ID)

	// Excluding externalId
	id = "2819c223-7f76-453a-919d-ab1234567891"
	m := make(map[attr.Path]bool)
	m[attr.Path{
		Name: "userName",
	}] = true
	m[attr.Path{
		Name: "schemas",
	}] = true
	m[attr.Path{
		Name: "id",
	}] = true
	m[attr.Path{
		Name: "meta",
	}] = true
	resource, err = adapter.Get(resTypeRepo.Get("User"), id, "", m)
	require.NoError(t, err)

	require.NotNil(t, resource)
	require.Equal(t, "", resource.ExternalID)

	// Non-existing ID
	id = "2819c223-7f76-453a-919d-ab1234567898"
	resource, err = adapter.Get(resTypeRepo.Get("User"), id, "", nil)
	require.Nil(t, resource)
	require.EqualError(t, err, "not found")

	// Empty ID
	resource, err = adapter.Get(resTypeRepo.Get("User"), "", "", nil)
	require.Nil(t, resource)
	require.EqualError(t, err, "not found")

}

func TestMongoUpdate(t *testing.T) {
	log.Println("TestMongoUpdate")
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	notExistingID := "zzzzzzzzzzzzzzzzzzzzzz"

	id := "2819c223-7f76-453a-919d-ab1234567891"
	dat, err := ioutil.ReadFile("../../testdata/user_to_update.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	err = adapter.Update(res, notExistingID, "")
	require.Error(t, err)

	err = adapter.Update(res, id, "")
	require.NoError(t, err)

}

func TestMongoFind(t *testing.T) {
	log.Println("TestMongoFind")
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	res := []*core.ResourceType{
		&core.ResourceType{
			CommonAttributes: core.CommonAttributes{
				Schemas: []string{"urn:ietf:params:scim:schemas:core:2.0:ResourceType"},
				ID:      "User",
				Meta: core.Meta{
					Location:     "https://example.com/v2/ResourceTypes/User",
					ResourceType: "ResourceType",
				},
			},
			Name:        "User",
			Endpoint:    "/User",
			Description: "User Account",
			Schema:      "urn:ietf:params:scim:schemas:core:2.0:User",
			SchemaExtensions: []core.SchemaExtension{
				core.SchemaExtension{
					Schema:   "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
					Required: true,
				},
			},
		},
	}

	var f filter.Filter = filter.AttrExpr{
		Path: attr.Path{
			URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
			Name: "userName",
			Sub:  "",
		},
		Op:    filter.OpEqual,
		Value: "tfork@example.com",
	}

	q, err := adapter.Find(res, f)
	count, err := q.Count()
	require.Nil(t, err)
	require.Equal(t, 1, count)

	if q != nil {
		q.Close()
	}

	// Invalid schema urn

	res[0].Schema = "invalid-urn"
	q, err = adapter.Find(res, f)
	require.Nil(t, q)

	if q != nil {
		q.Close()
	}
}

func TestMongoDelete(t *testing.T) {
	log.Println("TestMongoDelete")
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	// Delete object with specified id
	id := "2819c223-7f76-453a-919d-ab1234567891"
	err := adapter.Delete(resTypeRepo.Get("User"), id, "")
	require.NoError(t, err)

	id = "2819c223-7f76-453a-919d-111111111111"
	err = adapter.Delete(resTypeRepo.Get("User"), id, "")
	require.Error(t, err)
}
