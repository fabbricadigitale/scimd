package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
