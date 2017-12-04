package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndswith(t *testing.T) {
	var err error

	// Ends with
	okString := "bar"
	err = Validator.Var(okString, "endswith=r")
	require.NoError(t, err)

	// Doesn't ends with
	wrongString := "barz"
	err = Validator.Var(wrongString, "endswith=r")
	require.Error(t, err)

	// Invalid type
	invalidType := 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(invalidType, "endswith")
	})
}
