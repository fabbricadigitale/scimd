package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

//var validate *validator.Validate

func TestResourceTypeResource(t *testing.T) {
	// Non-normative of SCIM user resource type [https: //tools.ietf.org/html/rfc7643#section-8.6]
	dat, err := ioutil.ReadFile("testdata/user.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	res := ResourceType{}
	json.Unmarshal(dat, &res)

	assert.Contains(t, res.Schemas, "urn:ietf:params:scim:schemas:core:2.0:ResourceType")

	equalities := []struct {
		value string
		field interface{}
	}{
		{"User", res.ID},
		{"User", res.Name},
		{"/Users", res.Endpoint},
		{"User Account", res.Description},
		{"urn:ietf:params:scim:schemas:core:2.0:User", res.Schema},
		{"https://example.com/v2/ResourceTypes/User", res.Meta.Location},
		{"ResourceType", res.Meta.ResourceType},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	assert.Contains(t, res.SchemaExtensions, SchemaExtension{
		Required: true,
		Schema:   "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
	})
}

func TestResourceTypeValidation(t *testing.T) {
	res := NewResourceType("urn:ietf:params:scim:schemas:core:2.0:User", "ResourceType")
	res.Name = ""
	res.Endpoint = ""
	res.Schema = ""
	res.Meta.Location = "https://example.com/v2/ResourceTypes/User"
	now := time.Now()
	res.Meta.Created = &now
	res.Meta.LastModified = &now

	errors := validation.Validator.Struct(res)
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	require.Len(t, errors, 3)

	fields := []string{"Name", "Endpoint", "Schema"}
	failtags := []string{"required", "startswith", "urn", "required"}

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "ResourceType."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	//
	res.Name = "User"

	errors = validation.Validator.Struct(res)

	require.Len(t, errors, 2)

	fields = fields[1:]
	failtags = failtags[1:]

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "ResourceType."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	//
	res.Endpoint = "WrongEndpoint"

	errors = validation.Validator.Struct(res)
	require.Len(t, errors, 2)

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "ResourceType."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	res.Endpoint = "/Users"
	errors = validation.Validator.Struct(res)
	require.Len(t, errors, 1)

	fields = fields[1:]
	failtags = failtags[1:]

	// non urn on schema
	res.Schema = "urn:a:"
	errors = validation.Validator.Struct(res)
	require.Error(t, errors)

	// urn on schema
	res.Schema = "urn:ietf:params:scim:schemas:core:2.0:User"
	errors = validation.Validator.Struct(res)
	require.NoError(t, errors)

	// nested struct schemaext
	// Valid URN
	res.SchemaExtensions = []SchemaExtension{
		SchemaExtension{
			Schema: "urn:ietf:params:scim:schemas:core:2.0:User",
		},
	}
	errors = validation.Validator.Struct(res)
	require.NoError(t, errors)

	// Invalid URN
	res.SchemaExtensions = []SchemaExtension{
		SchemaExtension{
			Schema: "notValidUrn",
		},
	}
	errors = validation.Validator.Struct(res)
	require.Error(t, errors)
}
