package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURN(t *testing.T) {

	var err error

	// Valid URN RFC examples
	validURN := "URN:foo:a123,456"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	validURN = "URN:FOO:a123%2c456"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	// Valid URN Scim v2
	validURN = "urn:ietf:params:scim:schemas:core:2.0:User"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	validURN = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	// Valid URN - NSS can contain special characters
	validURN = "urn:ciao:!(xyz)+a,b.*@g=$_'"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	// Valid URN - ID can contain an hypen
	validURN = "URN:abcd-abcd:x"
	err = Validator.Var(validURN, "urn")
	require.NoError(t, err)

	// Invalid URN - ID cannot contain an hypen in first position
	invalidURN := "URN:-abcd:x"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	// Invalid URN - ID cannot contain spaces
	invalidURN = "urn:white space:NSS"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	invalidURN = "urn:concat:no spaces"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	// Invalid URN - Incomplete URN
	invalidURN = "urn:"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	invalidURN = "urn::"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	invalidURN = "urn:a"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	invalidURN = "urn:a:"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	// Invalid URN composition
	invalidURN = "urn:urn:params:scim:schemas:core:2.0:User"
	err = Validator.Var(invalidURN, "urn")
	require.Error(t, err)

	// Invalid type
	invalidURNType := 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(invalidURNType, "urn")
	})
}
