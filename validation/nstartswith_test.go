package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotStartswith(t *testing.T) {
	var err error

	// Not starts with
	okString := "rab"
	err = Validator.Var(okString, "nstartswith=b")
	require.NoError(t, err)

	// Starts with
	wrongString := "bar"
	err = Validator.Var(wrongString, "nstartswith=b")
	require.Error(t, err)

	// Invalid type
	invalidType := 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(invalidType, "nstartswith")
	})

}
