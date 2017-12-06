package integration

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	v "github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestMetaValidation(t *testing.T) {
	m := &core.Meta{}

	errors := v.Validator.Struct(m)
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	require.Len(t, errors, 4)

	fields := []string{"Location", "ResourceType", "Created", "LastModified"}
	failtags := []string{"uri", "required", "required", "required"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "Meta." + fields[e]
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	m.Location = "schema://uri"

	errors = v.Validator.Struct(m)

	fields = fields[1:]
	failtags = failtags[1:]

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "Meta." + fields[e]
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}
}

func TestCommonValidation(t *testing.T) {
	c := core.CommonAttributes{}

	errors := v.Validator.StructExcept(c, "Meta")
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	require.Len(t, errors, 2)

	fields := []string{"Schemas", "ID"}
	failtags := []string{"gt", "required", "excludes"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "CommonAttributes." + fields[e]
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	// Empty Schema fail on gt validation tag, ID fails on excludes validation tag
	c.Schemas = []string{}
	c.ID = "bulkId"
	errors = v.Validator.StructExcept(c, "Meta")
	require.Len(t, errors, 2)

	for _, err := range errors.(validator.ValidationErrors) {
		require.NotNil(t, err)
		require.IsType(t, (validator.ValidationErrors)(nil), errors)
	}

	// Wrong URN fail on urn validation tag, valid ID
	c.Schemas = []string{"not-a-urn"}
	c.ID = "test"
	errors = v.Validator.StructExcept(c, "Meta")
	require.Error(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	// Both schemas contents (ie., URNs) and ID valid
	c.Schemas = []string{"urn:ietf:params:scim:schemas:core:2.0:User"}
	c.ID = "test"
	errors = v.Validator.StructExcept(c, "Meta")
	require.NoError(t, errors)
}
