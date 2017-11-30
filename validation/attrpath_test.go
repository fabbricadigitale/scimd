package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttrPath(t *testing.T) {
	var err error

	// Valid path
	okPath := "urn:ietf:params:scim:schemas:core:2.0:User:userName"
	err = Validator.Var(okPath, "attrpath")
	require.NoError(t, err)

	// Wrong path
	wrongPath := "urn:ietf:params:scim:schemas:core:2.0"
	err = Validator.Var(wrongPath, "attrpath")
	require.Error(t, err)

	// Invalid type
	invalidTypePath := 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(invalidTypePath, "attrpath")
	})
}
