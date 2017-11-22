package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestEW struct {
	Text string `validate:"endswith=r,required"`
}

func TestEndswith(t *testing.T) {
	x := TestEW{}

	// Ends with
	x.Text = "bar"

	errors := Validator.Struct(x)
	require.NoError(t, errors, "err", nil)

	// Doesn't ends with
	x.Text = "bars"

	errors = Validator.Struct(x)
	require.Error(t, errors)
}
