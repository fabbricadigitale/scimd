package mold

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type maxTestCase struct {
	val interface{} // input value
	max interface{} // maximum threshold
	pos bool        // whether this is a positive test case or not - ie., should pass or not
	res interface{} // resulting value
}

var maxTests = []maxTestCase{
	// Above the max threshold
	{float32(0.4), float32(0.25), true, float32(0.25)},

	// Less and equal the max threshold
	{float32(0.15), float32(0.25), true, float32(0.15)},
	{float32(0.25), float32(0.25), true, float32(0.25)},

	// int8 above, less and equal the max threshold
	{int8(4), int8(2), true, int8(2)},
	{int8(1), int8(2), true, int8(1)},
	{int8(8), int8(8), true, int8(8)},

	// int16 above, less and equal the max threshold
	{int16(44), int16(22), true, int16(22)},
	{int16(11), int16(22), true, int16(11)},
	{int16(16), int16(16), true, int16(16)},

	// int32 above, less and equal the max threshold
	{int32(440), int32(220), true, int32(220)},
	{int32(110), int32(220), true, int32(110)},
	{int32(320), int32(320), true, int32(320)},

	// int64 above, less and equal the max threshold
	{int64(4400), int64(2200), true, int64(2200)},
	{int64(1100), int64(2200), true, int64(1100)},
	{int64(6400), int64(6400), true, int64(6400)},

	// uint8 above, less and equal the max threshold
	{uint8(4), uint8(3), true, uint8(3)},
	{uint8(1), uint8(2), true, uint8(1)},
	{uint8(7), uint8(7), true, uint8(7)},

	// uint16 above, less and equal the max threshold
	{uint16(44), uint16(22), true, uint16(22)},
	{uint16(11), uint16(22), true, uint16(11)},
	{uint16(16), uint16(16), true, uint16(16)},

	// uint32 above, less and equal the max threshold
	{uint32(440), uint32(220), true, uint32(220)},
	{uint32(110), uint32(220), true, uint32(110)},
	{uint32(320), uint32(320), true, uint32(320)},

	// uint64 above, less and equal the max threshold
	{uint64(4400), uint64(2200), true, uint64(2200)},
	{uint64(1100), uint64(2200), true, uint64(1100)},
	{uint64(6400), uint64(6400), true, uint64(6400)},

	// unsupported type
	{"0.4", float32(0.6), false, nil},
	{float32(0.4), "notanum", false, nil},
}

func maxherror(index int, test maxTestCase) string {
	return fmt.Sprintf("Test case num. %d. Result must be %+v when input %+v is more than threshold %+v", index+1, test.res, test.val, test.max)
}

func TestMax(t *testing.T) {
	for i, test := range maxTests {
		var err error
		input := test.val
		max := test.max

		switch max.(type) {
		case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
			err = Transformer.Field(context.Background(), &input, fmt.Sprintf("max=%v", max))
		case float32, float64:
			err = Transformer.Field(context.Background(), &input, fmt.Sprintf("max=%f", max))
		case string:
			err = Transformer.Field(context.Background(), &input, fmt.Sprintf("max=%s", max))
		}

		if test.pos {
			require.Nil(t, err)
			require.Equal(t, test.res, input, maxherror(i, test))
		} else {
			switch err.(type) {
			// Our own error
			case Error:
				require.EqualError(t, err, fmt.Sprintf("Bad field type %T", test.val))
			// Other errors
			default:
				require.Error(t, err)
			}
		}
	}
}
