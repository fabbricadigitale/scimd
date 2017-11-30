package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testPathOK struct {
	Path string `validate:"attrpath"`
}

type testPathInvalidType struct {
	Path int `validate:"attrpath"`
}

func TestAttrPath(t *testing.T) {
	x := testPathOK{}
	y := testPathInvalidType{}

	var err error

	// Valid path
	x.Path = "urn:ietf:params:scim:schemas:core:2.0:User:userName"
	err = Validator.Var(x, "attrpath")
	require.NoError(t, err)

	// Wrong path
	x.Path = "urn:ietf:params:scim:schemas:core:2.0"
	err = Validator.Var(x, "attrpath")
	require.Error(t, err)

	// Invalid type
	y.Path = 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(y, "attrpath")
	})
}
