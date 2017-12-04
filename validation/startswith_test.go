package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartswith(t *testing.T) {
	var err error

	// Starts with
	okString := "bar"
	err = Validator.Var(okString, "startswith=b")
	require.NoError(t, err)

	// Doesn't starts with
	wrongString := "zar"
	err = Validator.Var(wrongString, "startswith=b")
	require.Error(t, err)

	// Invalid type
	invalidType := 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(invalidType, "startswith")
	})

}
