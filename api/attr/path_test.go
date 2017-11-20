package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	path1 = `urn:ietf:params:scim:schemas:core:2.0:User:name.givenName`
)

func TestPath(t *testing.T) {

	a := Parse(path1)

	assert.Equal(t, "urn:ietf:params:scim:schemas:core:2.0:User", a.URI)
	assert.Equal(t, "name", a.Name)
	assert.Equal(t, "givenName", a.Sub)

	assert.Equal(t, path1, a.String())
}
