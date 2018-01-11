package attr

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	rt := resTypeRepo.Get("User")

	// FQN path / case insesitive
	p := Path{
		URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
		Name: "NAME",
		Sub:  "gIvEnNaMe",
	}
	ctx := p.Context(rt)
	assert.Equal(t, rt.GetSchema(), ctx.Schema)
	assert.Equal(t, rt.GetSchema().Attributes.WithName("name"), ctx.Attribute)
	assert.Equal(t, rt.GetSchema().Attributes.WithName("name").SubAttributes.WithName("givenName"), ctx.SubAttribute)

	// Just attribute name
	p = Path{
		Name: "userName",
	}
	ctx = p.Context(rt)
	assert.Equal(t, rt.GetSchema(), ctx.Schema)
	assert.Equal(t, rt.GetSchema().Attributes.WithName("userName"), ctx.Attribute)
	assert.Nil(t, ctx.SubAttribute)

	// Common attributes
	p = Path{
		Name: "meta",
		Sub:  "resourceType",
	}
	ctx = p.Context(rt)

	assert.Nil(t, ctx.Schema)
	assert.Equal(t, core.Commons().WithName("meta"), ctx.Attribute)
	assert.Equal(t, core.Commons().WithName("meta").SubAttributes.WithName("resourceType"), ctx.SubAttribute)

	// (todo) Name and Sub without URI

	// (todo) Schema ex

}

func TestContextSet(t *testing.T) {

	// single value and no sub-attribute
	p := Path{
		Name: "userName",
	}

	newUserName := datatype.String("leogr")
	p.Context(userRes.ResourceType()).Set(newUserName, &userRes)
	values := userRes.Values(userRes.ResourceType().GetSchema().GetIdentifier())
	assert.Equal(t, (*values)["userName"], newUserName)


	// multi value with sub-attribute
	p = Path{
		Name: "emails",
		Sub: "type",
	}

	newEmailType := datatype.String("newType")
	p.Context(userRes.ResourceType()).Set(newEmailType, &userRes)
	
	emails := (*values)["emails"].([]datatype.DataTyper)
	for _, elem := range emails {
		email := elem.(datatype.Complex)
		assert.Equal(t, newEmailType, email["type"])
	}

}