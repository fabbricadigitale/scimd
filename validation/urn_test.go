package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testURNOK struct {
	URN string `validate:"urn"`
}

type testInvalidURN struct {
	URN string `validate:"urn"`
}

type testInvalidURNType struct {
	URN int `validate:"urn"`
}

func TestURN(t *testing.T) {
	x := testURNOK{}
	y := testInvalidURN{}
	z := testInvalidURNType{}

	var err error

	// Valid URN
	x.URN = "urn:ietf:params:scim:schemas:core:2.0:User:name"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	// Invalid URN composition
	y.URN = "urn:urn:params:scim:schemas:core:2.0:User:name"
	require.PanicsWithValue(t, "Invalid URN composition: urn:urn:params:scim:schemas:core:2.0:User:name", func() {
		Validator.Var(y, "urn")
	})

	// Invalid type
	z.URN = 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(z, "urn")
	})
}
