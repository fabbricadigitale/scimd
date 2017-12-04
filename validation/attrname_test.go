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
	err = Validator.Struct(x)
	require.NoError(t, err)

	x.Name = "bar0"
	err = Validator.Struct(x)
	require.NoError(t, err)

	// Doesn't match Regex
	x.Name = "0bar"
	err = Validator.Struct(x)
	require.Error(t, err)

	// Invalid parent type
	invalidParentType := "bar"
	require.PanicsWithValue(t, "Invalid parent type string: must be a struct", func() {
		Validator.Var(invalidParentType, "attrname")
	})

	// Missing Type
	require.PanicsWithValue(t, "Field Type not found in the Struct", func() {
		Validator.Struct(z)
	})

	// Invalid Type
	require.PanicsWithValue(t, "Bad field type int", func() {
		Validator.Struct(y)
	})

	// Reference
	err = Validator.Struct(testCorrectReference)
	require.NoError(t, err)

	err = Validator.Struct(testWrongReference)
	require.Error(t, err)
}
