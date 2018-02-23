package messages

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
)

func TestSearchRequestResource(t *testing.T) {

	// Unmarshal
	dat, err := ioutil.ReadFile("testdata/search.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	s := SearchRequest{}
	json.Unmarshal(dat, &s)

	// s.Schemas exists
	assert.NotEmpty(t, s.Schemas)
	require.Contains(t, s.Schemas[0], "urn:ietf:params:scim:api:messages:2.0:SearchRequest")

	// Marshal
	r := SearchRequest{}
	r.Schemas = []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"}

	b, err2 := json.Marshal(r)

	require.NotNil(t, b)
	require.Nil(t, err2)
	require.JSONEq(t, `{
		"schemas":["urn:ietf:params:scim:api:messages:2.0:SearchRequest"]
		}`, string(b))

}

func TestSearchRequestValid(t *testing.T) {
	var err error

	// Right SearchRequest validation tag, Schemas not empty, there's only one URI inside of Schemas
	rightURI := []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"}
	err = validation.Validator.Var(rightURI, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.NoError(t, err)
	require.NotEmpty(t, rightURI)
	require.Len(t, rightURI, 1)

	// SearchRequest's Schemas cannot be empty
	noURI := []string{}
	err = validation.Validator.Var(noURI, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.Error(t, err)

	// Wrong SearchRequest validation tag
	wrongURI := []string{"urn:ietf:params:scim:api:messages:2.0"}
	err = validation.Validator.Var(wrongURI, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.Error(t, err)

	// Can't be two URIs inside SearchRequest's Schemas
	twoURIs := []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest", "urn:ietf:params:scim:api:messages:2.0:SearchRequest"}
	err = validation.Validator.Var(twoURIs, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.Error(t, err)
}
