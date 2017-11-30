package mold

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	// (todo) > testing table pleaseee, complete covering all possible cases, types and so on; see min_test.go
	var err error

	var aboveFloat32 float32 = 0.4
	err = Transformer.Field(context.Background(), &aboveFloat32, "max=0.25")
	require.Nil(t, err)
	require.Equal(t, float32(0.25), aboveFloat32)

	var belowFloat32 float32 = 0.15
	err = Transformer.Field(context.Background(), &belowFloat32, "max=0.25")
	require.Nil(t, err)
	require.Equal(t, float32(0.15), belowFloat32)

	var belowInt8 int8 = 1
	err = Transformer.Field(context.Background(), &belowInt8, "max=2")
	require.Nil(t, err)
	require.Equal(t, int8(1), belowInt8)
}
