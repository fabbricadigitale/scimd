package attr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	path1      = `urn:ietf:params:scim:schemas:core:2.0:User:name`
	path2      = `urn:ietf:params:scim:schemas:core:2.0:User:name.givenName`
	path3      = `urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber`
	path4      = `urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName`
	path5      = `userName`
	path6      = `name.givenName`
	path7      = `urn:ietf:params:scim:schemas:core:2.0`
	invalidUrn = `urn:urn:params:scim:schemas:core:2.0:User:name`
)

func TestPath(t *testing.T) {

	a := Parse(path1)

	assert.Equal(t, "urn:ietf:params:scim:schemas:core:2.0:User", a.URI)
	assert.Equal(t, "name", a.Name)

	assert.Equal(t, path1, a.String())

	assert.True(t, a.Valid())

	b := Parse(path2)

	assert.Equal(t, "urn:ietf:params:scim:schemas:core:2.0:User", b.URI)
	assert.Equal(t, "name", b.Name)
	assert.Equal(t, "givenName", b.Sub)

	assert.Equal(t, path2, b.String())

	c := Parse(path3)

	assert.Equal(t, "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", c.URI)
	assert.Equal(t, "employeeNumber", c.Name)

	assert.Equal(t, path3, c.String())

	d := Parse(path4)

	assert.Equal(t, "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", d.URI)
	assert.Equal(t, "manager", d.Name)
	assert.Equal(t, "displayName", d.Sub)

	assert.Equal(t, path4, d.String())

	e := Parse(path5)

	assert.Equal(t, "userName", e.Name)

	assert.Equal(t, path5, e.String())

	f := Parse(path6)

	assert.Equal(t, "name", f.Name)
	assert.Equal(t, "givenName", f.Sub)

	assert.Equal(t, path6, f.String())

	g := Parse(path7)
	assert.False(t, g.Valid())

	ko := Parse(invalidUrn)
	assert.False(t, ko.Valid())
}
