package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

type TestEW struct {
	Text    string `json:"text" validate:"endswith=r"`
	Integer int    `json:"integer" validate:"endswith=r"`
}

func TestEndswith(t *testing.T) {
	x := TestEW{}

	fields := []string{"Text", "Integer"}
	failtags := []string{"endswith", "endswith"}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered: ", r)
		}
	}()

	// Ends with
	x.Text = "bar"

	errors := Validator.Struct(x)
	require.NoError(t, errors)

	// Doesn't ends with
	x.Text = "bars"

	errors = Validator.Struct(x)
	require.Error(t, errors)

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "TestEW."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}
}
