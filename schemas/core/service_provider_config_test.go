package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{ChangePassword{Supported: false}, sp.ChangePassword},
		{Sort{Supported: true}, sp.Sort},
		{Etag{Supported: true}, sp.Etag},

		{"ServiceProviderConfig", sp.Meta.ResourceType},
		{"https://example.com/v2/ServiceProviderConfig", sp.Meta.Location},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	assert.Contains(t, sp.AuthenticationSchemas, AuthenticationSchema{
		Name:             "OAuth Bearer Token",
		Description:      "Authentication scheme using the OAuth Bearer Token Standard",
		SpecURI:          "http://www.rfc-editor.org/info/rfc6750",
		DocumentationURI: "http://example.com/help/oauth.html",
		Type:             "oauthbearertoken",
		Primary:          true,
	})

	assert.Contains(t, sp.AuthenticationSchemas, AuthenticationSchema{

		Name:             "HTTP Basic",
		Description:      "Authentication scheme using the HTTP Basic Standard",
		SpecURI:          "http://www.rfc-editor.org/info/rfc2617",
		DocumentationURI: "http://example.com/help/httpBasic.html",
		Type:             "httpbasic",
	})

}
