package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type isFieldOk struct {
	First  string
	Second int `validate:"isfield=First:hello"`
}

type isFieldEmptyParam struct {
	First  int
	Second int `validate:"isfield=:"`
}

type isFieldBadComp struct {
	First  int
	Second int `validate:"isfield=First"`
}

type isFieldInvalid struct {
	First  int
	Second int `validate:"isfield=Fir:1"`
}

func TestIsField(t *testing.T) {
	x := isFieldOk{}
	x.First = "hello"
	x.Second = 1

	err := Validator.Struct(x)
	require.NoError(t, err)

	// Empty parameters
	y := isFieldEmptyParam{}
	y.First = 1
	y.Second = 2

	require.PanicsWithValue(t, "No parameters specified", func() {
		Validator.Struct(y)
	})

	// Bad tag composition
	z := isFieldBadComp{}
	z.First = 1
	z.Second = 2

	require.PanicsWithValue(t, "Bad tag composition", func() {
		Validator.Struct(z)
	})

	// Field not found in the struct
	w := isFieldInvalid{}
	w.First = 1
	w.Second = 2

	require.PanicsWithValue(t, "Invalid field", func() {
		Validator.Struct(w)
	})

	/* TODO: test for value different from string (int and other values)

	eg:

	type example struct {
		First  int
		Second int `validate:"isfield=First:1"`
	}

	*/
}
