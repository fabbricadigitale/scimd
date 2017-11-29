package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testPathOK struct {
	Path string `validate:"attrpath"`
}

func TestAttrPath(t *testing.T) {
	x := testPathOK{}

	var err error

	// Valid path
	x.Path = "urn:ietf:params:scim:schemas:core:2.0:User:userName"
	err = Validator.Var(x, "attrpath")
	require.NoError(t, err)

	// Wrong path
	x.Path = "urn:ietf:params:scim:schemas:core:2.0"
	err = Validator.Var(x, "attrpath")
	require.Error(t, err)
}
