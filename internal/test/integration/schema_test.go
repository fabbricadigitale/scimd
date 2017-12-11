// +build integration

package integration

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	v "github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaResource(t *testing.T) {
	// Non-normative of SCIM user schama representation [https://tools.ietf.org/html/rfc7643#section-8.7.1]
	dat, err := ioutil.ReadFile("../../testdata/user_schema.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	sch := core.Schema{}
	json.Unmarshal(dat, &sch)

	equalities := []struct {
		value string
		field interface{}
	}{
		{"urn:ietf:params:scim:schemas:core:2.0:User", sch.ID},
		{"urn:ietf:params:scim:schemas:core:2.0:User", sch.CommonAttributes.ID},
		{"Schema", sch.Meta.ResourceType},
		{"Schema", sch.CommonAttributes.Meta.ResourceType},
		{"/v2/Schemas/urn:ietf:params:scim:schemas:core:2.0:User", sch.Meta.Location},
		{"/v2/Schemas/urn:ietf:params:scim:schemas:core:2.0:User", sch.CommonAttributes.Meta.Location},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	assert.Contains(t, sch.Attributes, &core.Attribute{
		Name:        "userName",
		Description: "Unique identifier for the User, typically used by the user to directly authenticate to the service provider. Each User MUST include a non-empty userName value.  This identifier MUST be unique across the service provider's entire set of Users. REQUIRED.",
		Type:        "string",
		MultiValued: false,
		Required:    true,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "server",
	})

	sa := []*core.Attribute{
		{
			Name:        "formatted",
			Description: "The full name, including all middle names, titles, and suffixes as appropriate, formatted for display (e.g., 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
		{
			Name:        "familyName",
			Description: "The family name of the User, or last name in most Western languages (e.g., 'Jensen' given the full name 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
		{
			Name:        "givenName",
			Description: "The given name of the User, or first name in most Western languages (e.g., 'Barbara' given the full name 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
		{
			Name:        "middleName",
			Description: "The middle name(s) of the User (e.g., 'Jane' given the full name 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
		{
			Name:        "honorificPrefix",
			Description: "The honorific prefix(es) of the User, or title in most Western languages (e.g., 'Ms.' given the full name 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
		{
			Name:        "honorificSuffix",
			Description: "The honorific suffix(es) of the User, or suffix in most Western languages (e.g., 'III' given the full name 'Ms. Barbara J Jensen, III').",
			Type:        "string",
			MultiValued: false,
			Required:    false,
			CaseExact:   false,
			Mutability:  "readWrite",
			Returned:    "default",
			Uniqueness:  "none",
		},
	}

	assert.Contains(t, sch.Attributes, &core.Attribute{
		Name:          "name",
		Description:   "The components of the user's real name. Providers MAY return just the full name as a single string in the formatted sub-attribute, or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both.  If both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the component attributes should be combined.",
		Type:          "complex",
		MultiValued:   false,
		Required:      false,
		Mutability:    "readWrite",
		Returned:      "default",
		Uniqueness:    "none",
		SubAttributes: sa,
	})
}

// https://tools.ietf.org/html/rfc7643#section-2.2
func assertAttributeDefaults(t *testing.T, a *core.Attribute) {
	assert.False(t, a.Required)
	assert.Empty(t, a.CanonicalValues)
	assert.False(t, a.CaseExact)
	assert.Equal(t, "readWrite", a.Mutability)
	assert.Equal(t, "default", a.Returned)
	assert.Equal(t, "none", a.Uniqueness)
	assert.Equal(t, "string", a.Type)
}

func TestNewAttribute(t *testing.T) {
	a := core.NewAttribute()
	assertAttributeDefaults(t, a)
}

func TestAttributeUnmarshalWithDefaults(t *testing.T) {
	data := []byte(`[{},{"type":"integer"}]`)
	list := make(core.Attributes, 0)

	require.NoError(t, json.Unmarshal(data, &list))
	require.Len(t, list, 2)

	first := list[0]
	require.IsType(t, (*core.Attribute)(nil), first)
	assertAttributeDefaults(t, first)

	second := list[1]
	require.IsType(t, (*core.Attribute)(nil), second)
	assert.Equal(t, "integer", second.Type)

	second.Type = first.Type
	assertAttributeDefaults(t, first)
}

func TestAttributeValidation(t *testing.T) {
	ok := core.NewAttribute()
	notOk := core.NewAttribute()

	var err error

	ok.Name = "bar"
	err = v.Validator.Var(ok, "attrname")
	require.NoError(t, err)

	notOk.Name = "0bar"
	err = v.Validator.Var(notOk, "attrname")
	require.Error(t, err)
}
