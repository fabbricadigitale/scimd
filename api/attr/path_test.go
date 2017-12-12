package attr

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
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

	ko := Parse(invalidUrn)
	assert.True(t, ko.Undefined())
}

func TestContext(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../../internal/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../../internal/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	rt := resTypeRepo.Get("User")

	// FQN path / case insesitive
	p := Path{
		URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
		Name: "NAME",
		Sub:  "gIvEnNaMe",
	}
	ctx := p.Context(rt)
	assert.Equal(t, rt.GetSchema(), ctx.Schema)
	assert.Equal(t, rt.GetSchema().Attributes.ByName("name"), ctx.Attribute)
	assert.Equal(t, rt.GetSchema().Attributes.ByName("name").SubAttributes.ByName("givenName"), ctx.SubAttribute)

	// Just attribute name
	p = Path{
		Name: "userName",
	}
	ctx = p.Context(rt)
	assert.Equal(t, rt.GetSchema(), ctx.Schema)
	assert.Equal(t, rt.GetSchema().Attributes.ByName("userName"), ctx.Attribute)
	assert.Nil(t, ctx.SubAttribute)

	// Common attributes
	p = Path{
		Name: "meta",
		Sub:  "resourceType",
	}
	ctx = p.Context(rt)

	assert.Nil(t, ctx.Schema)
	assert.Equal(t, core.Commons().ByName("meta"), ctx.Attribute)
	assert.Equal(t, core.Commons().ByName("meta").SubAttributes.ByName("resourceType"), ctx.SubAttribute)

	// (todo) Name and Sub without URI

	// (todo) Schema ex

}
