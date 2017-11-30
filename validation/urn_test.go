package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testURN struct {
	URN string `validate:"urn"`
}

type testInvalidURNType struct {
	URN int `validate:"urn"`
}

func TestURN(t *testing.T) {
	x := testURN{}
	z := testInvalidURNType{}

	var err error

	// Valid URN RFC examples
	x.URN = "URN:foo:a123,456"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	x.URN = "URN:FOO:a123%2c456"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	// Valid URN Scim v2
	x.URN = "urn:ietf:params:scim:schemas:core:2.0:User"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	x.URN = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	// Valid URN - NSS can contain special characters
	x.URN = "urn:ciao:#?!#(xyz)+a,b.*@g=$_'"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	// Valid URN - ID can contain an hypen
	x.URN = "URN:abcd-abcd:x"
	err = Validator.Var(x, "urn")
	require.NoError(t, err)

	// Invalid URN - ID cannot contain an hypen in first position
	x.URN = "URN:-abcd:x"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	// Invalid URN - ID cannot contain spaces
	x.URN = "urn:white space:NSS"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	x.URN = "urn:concat:no spaces"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	// Invalid URN - Incomplete URN
	x.URN = "urn:"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	x.URN = "urn::"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	x.URN = "urn:a"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	x.URN = "urn:a:"
	err = Validator.Var(x, "urn")
	require.Error(t, err)

	// Invalid URN composition
	x.URN = "urn:urn:params:scim:schemas:core:2.0:User"
	require.PanicsWithValue(t, "Invalid URN composition: urn:urn:params:scim:schemas:core:2.0:User", func() {
		Validator.Var(x, "urn")
	})

	// Invalid type
	z.URN = 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(z, "urn")
	})
}
