package mold

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type normurnTestCase struct {
	val interface{} // input urn
	pos bool        // whether this is a positive test case or not - ie., should pass or not
	res string      // resulting value
}

var normurnTests = []normurnTestCase{

	// Valid URN
	{"urn:A:b", true, "urn:a:b"},
	{"URN:A:b", true, "urn:a:b"},
	{"URN:A:B", true, "urn:a:B"},
	{"urn:foo:A123%C456", true, "urn:foo:A123%c456"},
	{"URN:FOO:a123%2C456", true, "urn:foo:a123%2c456"},

	// Not a valid URN
	{"URN:a:", false, ""},

	// Bad field type
	{123, false, ""},
}

func TestNormUrn(t *testing.T) {
	for _, test := range normurnTests {
		input := test.val
		err := Transformer.Field(context.Background(), &input, fmt.Sprintf("normurn=%s", test.val))
		fmt.Println(err)
		if test.pos {
			require.Nil(t, err)
			require.Equal(t, test.res, input)
		} else {
			switch err.(type) {
			case *InvalidUrnError:
				require.EqualError(t, err, fmt.Sprintf("Not a valid URN %s", test.val))
			case *BadTypeError:
				require.EqualError(t, err, fmt.Sprintf("Bad field type %T", test.val))
			default:
				require.Error(t, err)
			}
		}
	}
}
