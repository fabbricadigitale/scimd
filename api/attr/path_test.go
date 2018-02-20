package attr

import (
	"testing"

	"github.com/thoas/go-funk"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type attributeTest struct {
	attribute []string
	expected  []string
}

var attributesTest = []attributeTest{
	// Existing attribute
	{
		[]string{
			"userName",
		},
		[]string{
			"urn:ietf:params:scim:schemas:core:2.0:User:userName",
		},
	},
	{
		[]string{
			"emails",
		},
		[]string{
			"urn:ietf:params:scim:schemas:core:2.0:User:emails",
		},
	},
	// Non-existing attribute
	{
		[]string{
			"username",
		},
		[]string{},
	},
	// Existing Common Attributes (do not have a schema)
	{
		[]string{
			"id",
		},
		[]string{
			"id",
		},
	},
	{
		[]string{
			"externalId",
		},
		[]string{
			"externalId",
		},
	},
	// Attribute and one of its subattributes
	{
		[]string{
			"name",
			"familyName", // 2nd level
		},
		[]string{
			"urn:ietf:params:scim:schemas:core:2.0:User:name",
			"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName", // 2nd level
		},
	},
	{
		[]string{
			"name",
			"familyName", // 2nd level
			"userName",
		},
		[]string{
			"urn:ietf:params:scim:schemas:core:2.0:User:userName",
			"urn:ietf:params:scim:schemas:core:2.0:User:name",
			"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName", // 2nd level
		},
	},
}

const (
	path1      = `urn:ietf:params:scim:schemas:core:2.0:User:name`
	path2      = `urn:ietf:params:scim:schemas:core:2.0:User:name.givenName`
	path3      = `urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber`
	path4      = `urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName`
	path5      = `userName`
	path6      = `name.givenName`
	path7      = `urn:ietf:params:scim:schemas:core:2.0`
	path8      = `urn:ietf:params:scim:schemas:core:2.0:User:groups`
	path9      = `urn:ietf:params:scim:schemas:core:2.0:User:groups.display`
	path10     = `urn:ietf:params:scim:schemas:core:2.0:User:groups.$ref`
	invalidUrn = `urn:urn:params:scim:schemas:core:2.0:User:name`
	invalidRef = `urn:urn:params:scim:schemas:core:2.0:User:$ref`
)

func TestPath(t *testing.T) {
	a := Parse(path1)

	assert.Equal(t, "urn:ietf:params:scim:schemas:core:2.0:User", a.URI)
	assert.Equal(t, "name", a.Name)

	assert.Equal(t, path1, a.String())

	assert.False(t, a.Undefined())

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
	assert.True(t, g.Undefined())

	h := Parse(path8)
	assert.Equal(t, "groups", h.Name)

	i := Parse(path9)
	assert.Equal(t, "groups", i.Name)
	assert.Equal(t, "display", i.Sub)

	l := Parse(path10)
	assert.Equal(t, "groups", l.Name)
	assert.Equal(t, "$ref", l.Sub)

	ko1 := Parse(invalidUrn)
	assert.True(t, ko1.Undefined())

	ko2 := Parse(invalidRef)
	assert.True(t, ko2.Undefined())
}

func TestPaths(t *testing.T) {
	rt := resTypeRepo.Pull("User")

	for _, tt := range attributesTest {
		attr, err := Paths(rt, func(attribute *core.Attribute) bool {
			return funk.ContainsString(tt.attribute, attribute.Name)
		})
		require.NoError(t, err)

		results := make([]string, len(attr))
		for i, r := range attr {
			results[i] = r.String()
		}
		require.Equal(t, tt.expected, results)
	}

	// Paths(rt, nil) with fx = nil returns all attributes (ignoring their returned characteristic)
	var attrs = make(map[string]bool)
	p, err := Paths(rt, nil)
	require.NoError(t, err)
	for _, at := range p {
		attrs[at.String()] = true
	}
	require.Equal(t, withFxNil, attrs)
}
