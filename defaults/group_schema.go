package defaults

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/validation"
)

// GroupSchema is the default schema for the Group entity
var GroupSchema core.Schema

func init() {
	id := "urn:ietf:params:scim:schemas:core:2.0:Group"

	schema := core.NewSchema(id, "Group")
	schema.CommonAttributes.Meta.Location = fmt.Sprintf("/v2/Schemas/%s", id)
	schema.CommonAttributes.Meta.Version = v
	schema.Description = "Group"

	displayNameAttr := &core.Attribute{
		Name:        "displayName",
		Type:        "string",
		MultiValued: false,
		Description: "A human-readable name for the Group. REQUIRED.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	membersAttr := &core.Attribute{
		Name: "members",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "string",
				MultiValued: false,
				Description: "Identifier of the member of this Group.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "immutable",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "$ref",
				Type:        "reference",
				MultiValued: false,
				Description: "The URI corresponding to a SCIM resource that is a member of this Group.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "immutable",
				Returned:    "default",
				Uniqueness:  "none",
				ReferenceTypes: []string{
					"User",
					"Group",
				},
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the type of resource, e.g., 'User' or 'Group'.",
				Required:    false,
				CanonicalValues: []string{
					"User",
					"Group",
				},
				CaseExact:  false,
				Mutability: "immutable",
				Returned:   "default",
				Uniqueness: "none",
			},
		},
		MultiValued: true,
		Description: "A list of members of the Group.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	schema.Attributes = []*core.Attribute{
		displayNameAttr,
		membersAttr,
	}

	if errors := validation.Validator.Struct(schema); errors != nil {
		fmt.Println(errors)
		panic("default group schema configuration incorrect")
	}

	GroupSchema = *schema
}
