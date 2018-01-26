package messages

import (
	"encoding/json"
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

	// (todo) > assert s.Schemas exists first of all ...
	require.Equal(t, "urn:ietf:params:scim:api:messages:2.0:SearchRequest", s.Schemas[0])

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

	// (todo) > test minimum length = 1 // not empty

	// Wrong SearchRequest validation tag

	wrongURI := []string{"urn:ietf:params:scim:api:messages:2.0"}
	err = validation.Validator.Var(wrongURI, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.Error(t, err)

	// Right SearchRequest validation tag

	rightURI := []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"}
	err = validation.Validator.Var(rightURI, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest")
	require.NoError(t, err)
}
