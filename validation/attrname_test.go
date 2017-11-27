package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testOK struct {
	Name string `validate:"attrname"`
	Type string
}

type testInvalidType struct {
	Name int `validate:"attrname"`
	Type string
}

type testMissingType struct {
	Name string `validate:"attrname"`
}

func TestAttrName(t *testing.T) {
	x := testOK{}
	y := testInvalidType{}
	z := testMissingType{}

	var err error

	testCorrectReference := testOK{
		Name: "$ref",
		Type: "reference",
	}

	testWrongReference := testOK{
		Name: "abc",
		Type: "reference",
	}

	// Match Regex
	x.Name = "bar"
	err = Validator.Var(x, "attrname")
	require.NoError(t, err)

	x.Name = "bar0"
	err = Validator.Var(x, "attrname")
	require.NoError(t, err)

	// Doesn't match Regex
	x.Name = "0bar"
	err = Validator.Var(x, "attrname")
	require.Error(t, err)

	// Invalid parent type
	require.PanicsWithValue(t, "Invalid parent type string: must be a struct", func() {
		Validator.Var(x.Name, "attrname")
	})

	// Missing Type
	z.Name = "bar"
	require.PanicsWithValue(t, "Field Type not found in the Struct", func() {
		Validator.Var(z, "attrname")
	})

	// Invalid Type
	y.Name = 123
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Var(y, "attrname")
	})

	// Reference
	err = Validator.Var(testCorrectReference, "attrname")
	require.NoError(t, err)

	err = Validator.Var(testWrongReference, "attrname")
	require.Error(t, err)
}
