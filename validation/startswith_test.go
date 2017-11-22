package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestSW struct {
	Text string `validate:"startswith=b,required"`
}

func TestStartswith(t *testing.T) {
	x := TestSW{}

	// Starts with
	x.Text = "bar"

	errors := Validator.Struct(x)
	require.NoError(t, errors, "err", nil)

	// Doesn't starts with
	x.Text = "zar"

	errors = Validator.Struct(x)
	require.Error(t, errors)
}
