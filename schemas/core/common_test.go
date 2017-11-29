package core

import (
	"fmt"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestMetaValidation(t *testing.T) {
	m := &Meta{}

	errors := validation.Validator.Struct(m)
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

	errors = validation.Validator.Struct(m)

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
	c := CommonAttributes{}

	fmt.Println(c)

	errors := validation.Validator.StructExcept(c, "Meta")
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

	c.Schemas = []string{"not-a-urn"}
	errors = validation.Validator.StructExcept(c, "Meta")

	c.ID = "bulkId"
	errors = validation.Validator.StructExcept(c, "Meta")
	require.NotNil(t, errors)

	c.ID = "bulkID"
	errors = validation.Validator.StructExcept(c, "Meta")
	require.Nil(t, errors)

	// (todo) > complete when urn validator will be done
	fmt.Println(errors)
}
