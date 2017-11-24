package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
)

type testAN struct {
	Name    string `validate:"attrname"`
	Type    string `validate:"attrname"`
	Integer int    `validate:"attrname"`
}

func TestAttrName(t *testing.T) {
	x := testAN{}

	fields := []string{"Name", "Type", "Integer"}
	failtags := []string{"attrname", "attrname", "attrname"}

	defer func() {
		r := recover()
		require.NotNil(t, r)
		require.Equal(t, "Bad field type int", r)
	}()

	// Match Regex
	x.Name = "bar"

	errors := Validator.Var(x, "attrname")
	require.NoError(t, errors)

	x.Name = "bar0"

	errors = Validator.Var(x, "attrname")
	require.NoError(t, errors)

	// Doesn't match Regex
	x.Name = "0bar"

	errors = Validator.Var(x, "attrname")
	require.Error(t, errors)

	for e, err := range errors.(validator.ValidationErrors) {
		require.Equal(t, "testAN."+fields[e], err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}
}
