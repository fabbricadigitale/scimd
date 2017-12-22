package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ParentRight struct {
	Age int `validate:"topstructname=GranpaRight"`
}

type GranpaRight struct {
	Parent ParentRight
}

type ParentWrong struct {
	Age int `validate:"topstructname=Parent"`
}

type GranpaWrong struct {
	Parent ParentWrong
}

type InvalidParam struct {
	Param int `validate:"topstructname="`
}

func TestIsGranpa(t *testing.T) {
	// Right top struct name
	x := GranpaRight{}
	x.Parent.Age = 2

	err := Validator.Struct(x)
	require.NoError(t, err)

	// Wrong top struct name
	y := GranpaWrong{}
	y.Parent.Age = 3

	err = Validator.Struct(y)
	require.Error(t, err)

	// Invalid parent type
	z := 4
	require.PanicsWithValue(t, "Invalid parent type int: must be a struct", func() {
		Validator.Var(z, "topstructname=Granpa")
	})

	// Invalid parameter
	w := InvalidParam{}
	require.PanicsWithValue(t, "Invalid param value", func() {
		Validator.Struct(w)
	})
}
