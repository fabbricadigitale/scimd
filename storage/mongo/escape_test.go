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
