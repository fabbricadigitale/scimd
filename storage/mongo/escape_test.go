package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyEscaping(t *testing.T) {

	tests := []string{
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		"$ref",
	}

	for _, key := range tests {
		require.True(t, keyRegexp.MatchString(key)) // ensure test cases are valid

		escapedKey := keyEscape(key)
		assert.False(t, keyRegexp.MatchString(escapedKey))
		assert.NotEqual(t, key, escapedKey)

		unescapedKey := keyUnescape(escapedKey)
		assert.Equal(t, key, unescapedKey)
	}

}

type testCase struct {
	key      string
	expected string
}

func TestEscapeAttribute(t *testing.T) {
	testCases := []testCase{
		{
			key:      "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:userName",
			expected: "urn:ietf:params:scim:schemas:extension:enterprise:2°0:User.userName",
		},
		{
			key:      "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:name.familyName",
			expected: "urn:ietf:params:scim:schemas:extension:enterprise:2°0:User.name.familyName",
		},
		{
			key:      "$ref",
			expected: "§ref",
		},
	}

	for _, tc := range testCases {
		require.True(t, keyRegexp.MatchString(tc.key)) // ensure test cases are valid

		escapedKey := escapeAttribute(tc.key)
		assert.NotEqual(t, tc.key, escapedKey)
		assert.Equal(t, tc.expected, escapedKey)
	}
}
