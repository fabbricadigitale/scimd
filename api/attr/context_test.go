package attr

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	rt := resTypeRepo.Get("User")

	// Unknown attribute
	s := Parse("unknown")
	ctx := s.Context(rt)
	assert.Nil(t, ctx)

	// Name and Sub without URI
	// Unknown since not present on base schema
	p := Path{
		Name: "MANAGER",
		Sub:  "value",
	}
	ctx = p.Context(rt)
	assert.Nil(t, ctx)

	// FQN path / case insensitive
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
		Name: "NAME",
		Sub:  "gIvEnNaMe",
	}
	ctx = p.Context(rt)
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

	// Name and Sub without URI
	p = Path{
		Name: "name",
		Sub:  "familyName",
	}
	ctx = p.Context(rt)

	assert.Equal(t, rt.GetSchema(), ctx.Schema)
	assert.Equal(t, rt.GetSchema().Attributes.WithName("name"), ctx.Attribute)
	assert.Equal(t, rt.GetSchema().Attributes.WithName("name").SubAttributes.WithName("familyName"), ctx.SubAttribute)

	// Schema extension
	// FQN path / case insesitive
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name: "MANAGER",
		Sub:  "DisplayNaMe",
	}
	ctx = p.Context(rt)
	assert.Equal(t, rt.GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"], ctx.Schema)
	assert.Equal(t, rt.GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].Attributes.WithName("manager"), ctx.Attribute)
	assert.Equal(t, rt.GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].Attributes.WithName("manager").SubAttributes.WithName("displayName"), ctx.SubAttribute)

	// URI and Name
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name: "emPLOyeeNumber",
	}
	ctx = p.Context(rt)
	assert.Equal(t, rt.GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"], ctx.Schema)
	assert.Equal(t, rt.GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].Attributes.WithName("employeeNumber"), ctx.Attribute)
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
		Sub:  "type",
	}

	newEmailType := datatype.String("newType")
	p.Context(userRes.ResourceType()).Set(newEmailType, &userRes)

	emails := (*values)["emails"].([]datatype.DataTyper)
	for _, elem := range emails {
		email := elem.(datatype.Complex)
		assert.Equal(t, newEmailType, email["type"])
	}

	// extension
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name: "MANAGER",
		Sub:  "DisplayNaMe",
	}

	newDisplayName := datatype.String("ldcxaqqa")
	p.Context(userRes.ResourceType()).Set(newDisplayName, &userRes)
	values = userRes.Values(userRes.ResourceType().GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].GetIdentifier())
	assert.Equal(t, newDisplayName, (*values)["manager"].(datatype.Complex)["displayName"])

	// accomodation of singlevalue in a multivalued attribute
	p = Path{
		Name: "emails",
	}

	newEmail := datatype.Complex{}
	newEmail["value"] = "me@mail.com"
	newEmail["type"] = "work"

	p.Context(userRes.ResourceType()).Set(newEmail, &userRes)
	values = userRes.Values(userRes.ResourceType().GetSchema().GetIdentifier())
	em := (*values)["emails"]

	assert.IsType(t, datatype.Complex{}, em)

	// accomodation of a multivalue in a singlevalued attribute
	p = Path{
		Name: "Name",
	}

	newName := make([]datatype.Complex, 2)
	newName[0] = datatype.Complex{
		"givenName":  "Bob",
		"familyName": "abc",
	}
	newName[1] = datatype.Complex{
		"givenName":  "Alice",
		"familyName": "def",
	}

	p.Context(userRes.ResourceType()).Set(newName, &userRes)
	name := (*values)["name"]
	assert.IsType(t, []datatype.Complex{}, name)

	// accomodation of a complex attribute in a string attribute
	p = Path{
		Name: "userName",
	}

	newUserNameComplex := datatype.Complex{
		"userName": "leogr",
	}

	p.Context(userRes.ResourceType()).Set(newUserNameComplex, &userRes)
	userName := (*values)["userName"]
	assert.IsType(t, datatype.Complex{}, userName)

	newUserNameInteger := datatype.Integer(123)
	p.Context(userRes.ResourceType()).Set(newUserNameInteger, &userRes)
	userName = (*values)["userName"]
	assert.IsType(t, datatype.Integer(123), userName)
}

func TestContextGet(t *testing.T) {

	// single value and no sub-attribute
	p := Path{
		Name: "userName",
	}

	userName := p.Context(userRes.ResourceType()).Get(&userRes)
	assert.Equal(t, datatype.String("tfork@example.com"), userName)

	// multi value
	p = Path{
		Name: "emails",
	}

	emails := p.Context(userRes.ResourceType()).Get(&userRes)
	assert.IsType(t, []datatype.DataTyper{}, emails)

	// multi value with sub-attribute
	p = Path{
		Name: "emails",
		Sub:  "type",
	}

	emails = p.Context(userRes.ResourceType()).Get(&userRes)
	assert.IsType(t, []interface{}{}, emails)

	// extension
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name: "MANAGER",
		Sub:  "DisplayNaMe",
	}

	sub := p.Context(userRes.ResourceType()).Get(&userRes)
	assert.Equal(t, datatype.String("John Smith"), sub)
}

func TestContextDelete(t *testing.T) {
	p := Path{
		Name: "userName",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)
	values := userRes.Values(userRes.ResourceType().GetSchema().GetIdentifier())
	userName := (*values)["userName"]
	assert.Nil(t, userName)

	// multi value
	p = Path{
		Name: "emails",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)
	emails := (*values)["emails"]
	assert.Nil(t, emails)

	// multi value with sub-attribute
	p = Path{
		Name: "ims",
		Sub:  "type",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)
	ims := (*values)["ims"].([]datatype.DataTyper)
	for _, elem := range ims {
		email := elem.(datatype.Complex)
		assert.Nil(t, email["type"])
	}

	// extension
	p = Path{
		URI:  "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
		Name: "MANAGER",
		Sub:  "DisplayNaMe",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)
	values = userRes.Values(userRes.ResourceType().GetSchemaExtensions()["urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"].GetIdentifier())
	displayName := (*values)["manager"].(datatype.Complex)["displayName"]
	assert.Nil(t, displayName)
}

// (todo) > Test contexts() with filter about attributes sub attributes characteristic
