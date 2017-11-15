package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestServiceProviderConfigResource(t *testing.T) {
	// Non-normative of SCIM service provider configuration [https://tools.ietf.org/html/rfc7643#section-8.5]
	dat, err := ioutil.ReadFile("testdata/service_provider_config.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	sp := ServiceProviderConfig{}
	json.Unmarshal(dat, &sp)

	assert.Contains(t, sp.Schemas, "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig")

	equalities := []struct {
		value interface{}
		field interface{}
	}{
		{"http://example.com/help/scim.html", sp.DocumentationURI},
		{true, sp.Patch.Supported},
		{true, sp.Bulk.Supported},
		{1000, sp.Bulk.MaxOperations},
		{1048576, sp.Bulk.MaxPayloadSize},
		{true, sp.Filter.Supported},
		{200, sp.Filter.MaxResults},
		{changePassword{Supported: false}, sp.ChangePassword},
		{sort{Supported: true}, sp.Sort},
		{etag{Supported: true}, sp.Etag},

		{"ServiceProviderConfig", sp.Meta.ResourceType},
		{"https://example.com/v2/ServiceProviderConfig", sp.Meta.Location},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	assert.Contains(t, sp.AuthenticationSchemes, authenticationScheme{
		Name:             "OAuth Bearer Token",
		Description:      "Authentication scheme using the OAuth Bearer Token Standard",
		SpecURI:          "http://www.rfc-editor.org/info/rfc6750",
		DocumentationURI: "http://example.com/help/oauth.html",
		Type:             "oauthbearertoken",
		Primary:          true,
	})

	assert.Contains(t, sp.AuthenticationSchemes, authenticationScheme{
		Name:             "HTTP Basic",
		Description:      "Authentication scheme using the HTTP Basic Standard",
		SpecURI:          "http://www.rfc-editor.org/info/rfc2617",
		DocumentationURI: "http://example.com/help/httpBasic.html",
		Type:             "httpbasic",
		Primary:          false,
	})

	assert.Len(t, sp.AuthenticationSchemes, 2)
}

func TestServiceProviderConfigValidation(t *testing.T) {
	res := &ServiceProviderConfig{}

	errors := validation.Validator.Struct(res)
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	require.Len(t, errors, 10)

	structs := []string{"Patch", "Bulk", "Bulk", "Bulk", "Filter", "Filter", "ChangePassword", "Sort", "Etag", "AuthenticationSchemes"}
	fields := []string{"Supported", "Supported", "MaxOperations", "MaxPayloadSize", "Supported", "MaxResults", "Supported", "Supported", "Supported", ""}
	failtags := []string{"required", "required", "required", "required", "required", "required", "required", "required", "required", "required"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "ServiceProviderConfig." + structs[e]
		if len(fields[e]) > 0 {
			exp += "." + fields[e]
		} else {
			fields[e] = structs[e]
		}
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	// (todo) > complete test with positive and negative cases when (and if) struct'll have other validations other than required
}
