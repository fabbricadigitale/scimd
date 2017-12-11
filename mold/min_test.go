package mold

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type minTestCase struct {
	val interface{} // input value
	min interface{} // minimum threshold
	pos bool        // whether this is a positive test case or not - ie., should pass or not
	res interface{} // resulting value
}

var minTests = []minTestCase{
	// less than the minimum threshold
	{float32(0.15), float64(0.25), true, float32(0.25)},

	// greather (or equal) than the minimum threshold
	{float32(0.40), float64(0.25), true, float32(0.40)},
	{float32(0.25), float64(0.25), true, float32(0.25)},

	// unsupported type
	{"0.4", float64(0.4), false, nil},
	{float64(0.4), "notanum", false, nil},
}

func herror(index int, test minTestCase) string {
	return fmt.Sprintf("Test case num. %d. Result must be %+v when input %+v is less than threshold %+v", index+1, test.res, test.val, test.min)
}

func TestMin(t *testing.T) {
	for i, test := range minTests {
		input := test.val
		err := Transformer.Field(context.Background(), &input, fmt.Sprintf("min=%f", test.min))
		if test.pos {
			require.Nil(t, err)
			require.Equal(t, test.res, input, herror(i, test))
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
