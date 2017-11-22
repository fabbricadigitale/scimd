package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

type TestSW struct {
	Text    string `validate:"startswith=b"`
	Integer int    `validate:"startswith=b"`
}

func TestStartswith(t *testing.T) {
	x := TestSW{}

	fields := []string{"Text", "Integer"}
	failtags := []string{"startswith", "startswith"}

	defer func() {
		r := recover()
		require.NotNil(t, r)
		require.Equal(t, "Bad field type int", r)
	}()

	// Starts with
	x.Text = "bar"

	errors := Validator.Struct(x)
	require.NoError(t, errors)

	// Doesn't starts with
	x.Text = "zar"

	errors = Validator.Struct(x)
	require.Error(t, errors)

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "TestSW."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

}
