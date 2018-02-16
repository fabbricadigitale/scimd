package attr

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/assert"
)

var withFxNil = map[string]bool{
	"schemas":           true,
	"id":                true,
	"externalId":        true,
	"meta":              true,
	"meta.resourceType": true,
	"meta.created":      true,
	"meta.lastModified": true,
	"meta.location":     true,
	"meta.version":      true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber":      true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:costCenter":          true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:organization":        true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division":            true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:department":          true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager":             true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.value":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.$ref":        true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userName":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name":                                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.formatted":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName":                     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.givenName":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.middleName":                     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificPrefix":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificSuffix":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:displayName":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:nickName":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:profileUrl":                          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:title":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userType":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:preferredLanguage":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:locale":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:timezone":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:active":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:password":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.primary":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.value":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.display":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.type":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.primary":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims":                                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.value":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.display":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.type":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.primary":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.primary":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.formatted":                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.streetAddress":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.locality":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.region":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.postalCode":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.country":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.type":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.primary":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.$ref":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.value":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.display":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.type":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.primary":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.value":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.display":                       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.type":                          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.primary":                       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.value":              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.display":            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.type":               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.primary":            true,
}

var withReturnedAlways = map[string]bool{
	"schemas":           true,
	"id":                true,
	"meta.resourceType": true,
}

var withReturnedNever = map[string]bool{
	"urn:ietf:params:scim:schemas:core:2.0:User:password": true,
}

var withReturnedDefault = map[string]bool{
	"externalId":        true,
	"meta":              true,
	"meta.created":      true,
	"meta.lastModified": true,
	"meta.location":     true,
	"meta.version":      true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber":      true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:costCenter":          true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:organization":        true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division":            true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:department":          true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager":             true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.value":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.$ref":        true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userName":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name":                                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.formatted":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName":                     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.givenName":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.middleName":                     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificPrefix":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificSuffix":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:displayName":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:nickName":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:profileUrl":                          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:title":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userType":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:preferredLanguage":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:locale":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:timezone":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:active":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.primary":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.value":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.display":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.type":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.primary":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims":                                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.value":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.display":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.type":                            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.primary":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.primary":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.formatted":                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.streetAddress":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.locality":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.region":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.postalCode":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.country":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.type":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.primary":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.value":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.$ref":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.display":                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.type":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.value":                  true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.display":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.type":                   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.primary":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.value":                         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.display":                       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.type":                          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.primary":                       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.value":              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.display":            true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.type":               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.primary":            true,
}

func TestContext(t *testing.T) {
	rt := resTypeRepo.Pull("User")

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

func TestContexts(t *testing.T) {
	rt := resTypeRepo.Pull("User")

	var fxNilAttrs = make(map[string]bool)
	fxnil := Contexts(rt, nil)
	for _, at := range fxnil {
		fxNilAttrs[at.Path().String()] = true
	}
	require.Equal(t, withFxNil, fxNilAttrs)

	var raAttrs = make(map[string]bool)
	ra := Contexts(rt, withReturned(schemas.ReturnedAlways))
	for _, at := range ra {
		raAttrs[at.Path().String()] = true
	}
	require.Equal(t, withReturnedAlways, raAttrs)

	var rnAttrs = make(map[string]bool)
	rn := Contexts(rt, withReturned(schemas.ReturnedNever))
	for _, at := range rn {
		rnAttrs[at.Path().String()] = true
	}
	require.Equal(t, withReturnedNever, rnAttrs)

	var rdAttrs = make(map[string]bool)
	rd := Contexts(rt, withReturned(schemas.ReturnedDefault))
	for _, at := range rd {
		rdAttrs[at.Path().String()] = true
	}
	require.Equal(t, withReturnedDefault, rdAttrs)

	// TODO > What to do when rt is empty? Panics in core/resource_type/#58 on the GetIdentifier() method
	//
	// rt := &core.ResourceType{}
	// rd := Contexts(rt, withReturned(schemas.ReturnedDefault))

}
