// +build integration

package integration

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/fabbricadigitale/scimd/schemas/core"
	v "github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

func TestServiceProviderConfigResource(t *testing.T) {
	// Non-normative of SCIM service provider configuration [https://tools.ietf.org/html/rfc7643#section-8.5]
	dat, err := ioutil.ReadFile("../../testdata/service_provider_config.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	sp := core.ServiceProviderConfig{}
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
		{false, sp.ChangePassword.Supported},
		{true, sp.Sort.Supported},
		{true, sp.Etag.Supported},

		{"ServiceProviderConfig", sp.Meta.ResourceType},
		{"https://example.com/v2/ServiceProviderConfig", sp.Meta.Location},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	assert.Contains(t, sp.AuthenticationSchemes, core.AuthenticationScheme{
		Name:             "OAuth Bearer Token",
		Description:      "Authentication scheme using the OAuth Bearer Token Standard",
		SpecURI:          "http://www.rfc-editor.org/info/rfc6750",
		DocumentationURI: "http://example.com/help/oauth.html",
		Type:             "oauthbearertoken",
		Primary:          true,
	})

	assert.Contains(t, sp.AuthenticationSchemes, core.AuthenticationScheme{
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
	res := core.NewServiceProviderConfig()
	res.ID = "User"
	res.Meta.Location = "https://example.com/v2/ResourceTypes/User"
	now := time.Now()
	res.Meta.Created = &now
	res.Meta.LastModified = &now

	errors := v.Validator.Struct(res)
	require.NotNil(t, errors)
	//require.IsType(t, (validator.ValidationErrors)(nil), errors)

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

	res.Patch.Supported = true
	res.Bulk.Supported = true
	res.Bulk.MaxOperations = 1000
	res.Bulk.MaxPayloadSize = 1048576
	res.Filter.Supported = true
	res.Filter.MaxResults = 200
	res.ChangePassword.Supported = true
	res.Sort.Supported = true
	res.Etag.Supported = true
	res.AuthenticationSchemes = []core.AuthenticationScheme{}

	// Valid URI
	res.DocumentationURI = "http://example.com/help/scim.html"
	errors = v.Validator.Struct(res)
	require.NoError(t, errors)

	// Invalid URI
	res.DocumentationURI = "NotAUri"
	errors = v.Validator.Struct(res)
	require.Error(t, errors)
}

func TestAuthenticationSchemeValidation(t *testing.T) {
	x := &core.AuthenticationScheme{}

	errors := v.Validator.Struct(x)
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)

	fields := []string{"Type", "Name", "Description"}
	failtags := []string{"required", "required", "required"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "AuthenticationScheme." + fields[e]
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	// Non matching Type for eq validation tag
	x.Name = "name"
	x.Description = "descr"
	x.Type = "xxx"

	errors = v.Validator.Struct(x)

	fields = fields[:1]
	failtags = []string{"eq=oauth|eq=oauth2|eq=oauthbearertoken|eq=httpbasic|eq=httpdigest"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "AuthenticationScheme." + fields[e]
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	// Matching Type for eq validation tag
	x.Type = "oauth2"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	x.Type = "oauth"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	x.Type = "oauthbearertoken"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	x.Type = "httpbasic"

	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	x.Type = "httpdigest"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	// Valid URI
	x.SpecURI = "http://www.rfc-editor.org/info/rfc2617"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	x.DocumentationURI = "http://example.com/help/httpBasic.html"
	errors = v.Validator.Struct(x)
	require.NoError(t, errors)

	// Invalid URI

	x.SpecURI = "NotAUri"
	errors = v.Validator.Struct(x)
	require.Error(t, errors)

	x.DocumentationURI = "NotAUri"
	errors = v.Validator.Struct(x)
	require.Error(t, errors)
}
