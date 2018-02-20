package defaults

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/validation"
)

// UserSchema is the default schema for the User entity
var UserSchema core.Schema

func init() {
	id := "urn:ietf:params:scim:schemas:core:2.0:User"

	schema := core.NewSchema(id, "User")
	schema.CommonAttributes.Meta.Location = fmt.Sprintf("/v2/Schemas/%s", id)
	schema.CommonAttributes.Meta.Version = v
	schema.Description = "User Account"

	userNameAttr := &core.Attribute{
		Name:        "userName",
		Type:        "string",
		MultiValued: false,
		Description: "Unique identifier for the User, typically used by the user to directly authenticate to the service provider. Each User MUST include a non-empty userName value.  This identifier MUST be unique across the service provider's entire set of Users. REQUIRED.",
		Required:    true,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "server",
	}

	nameAttr := &core.Attribute{
		Name: "name",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "formatted",
				Type:        "string",
				MultiValued: false,
				Description: "The full name, including all middle names, titles, and suffixes as appropriate, formatted for display (e.g., 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "familyName",
				Type:        "string",
				MultiValued: false,
				Description: "The family name of the User, or last name in most Western languages (e.g., 'Jensen' given the full name 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "givenName",
				Type:        "string",
				MultiValued: false,
				Description: "The given name of the User, or first name in most Western languages (e.g., 'Barbara' given the full name 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "middleName",
				Type:        "string",
				MultiValued: false,
				Description: "The middle nameof the User (e.g., 'Jane' given the full name 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "honorificPrefix",
				Type:        "string",
				MultiValued: false,
				Description: "The honorific prefix(es) of the User, or title in most Western languages (e.g., 'Ms.' given the full name 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "honorificSuffix",
				Type:        "string",
				MultiValued: false,
				Description: "The honorific suffix(es) of the User, or suffix in most Western languages (e.g., 'III' given the full name 'Ms. Barbara J Jensen, III').",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: false,
		Description: "The components of the user's real name. Providers MAY return just the full name as a single string in the formatted sub-attribute, or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both.  If both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the component attributes should be combined.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	displayNameAttr := &core.Attribute{
		Name:        "displayName",
		Type:        "string",
		MultiValued: false,
		Description: "The name of the User, suitable for display to end-users.  The name SHOULD be the full name of the User being described, if known.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	nickNameAttr := &core.Attribute{
		Name:        "nickName",
		Type:        "string",
		MultiValued: false,
		Description: "The casual way to address the user in real life, e.g., 'Bob' or 'Bobby' instead of 'Robert'.  This attribute SHOULD NOT be used to represent a User's username (e.g., 'bjensen' or 'mpepperidge').",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	profileURLAttr := &core.Attribute{
		Name:        "profileUrl",
		Type:        "reference",
		MultiValued: false,
		Description: "A fully qualified URL pointing to a page representing the User's online profile.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
		ReferenceTypes: []string{
			"external",
		},
	}

	titleAttr := &core.Attribute{
		Name:        "title",
		Type:        "string",
		MultiValued: false,
		Description: "The user's title, such as \"Vice President.\"",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	userTypeAttr := &core.Attribute{
		Name:        "userType",
		Type:        "string",
		MultiValued: false,
		Description: "Used to identify the relationship between the organization and the user.  Typical values used might be 'Contractor', 'Employee', 'Intern', 'Temp', 'External', and 'Unknown', but any value may be used.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	preferredLanguageAttr := &core.Attribute{
		Name:        "preferredLanguage",
		Type:        "string",
		MultiValued: false,
		Description: "Indicates the User's preferred written or spoken language.  Generally used for selecting a localized user interface; e.g., 'en_US' specifies the language English and country US.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	localeAttr := &core.Attribute{
		Name:        "locale",
		Type:        "string",
		MultiValued: false,
		Description: "Used to indicate the User's default location for purposes of localizing items such as currency, date time format, or numerical representations.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	timezoneAttr := &core.Attribute{
		Name:        "timezone",
		Type:        "string",
		MultiValued: false,
		Description: "The User's time zone in the 'Olson' time zone database format, e.g., 'America/Los_Angeles'.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	activeAttr := &core.Attribute{
		Name:        "active",
		Type:        "boolean",
		MultiValued: false,
		Description: "A Boolean value indicating the User's administrative status.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	passwordAttr := &core.Attribute{
		Name:        "password",
		Type:        "string",
		MultiValued: false,
		Description: "The User's cleartext password.  This attribute is intended to be used as a means to specify an initial password when creating a new User or to reset an existing User's password.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "writeOnly",
		Returned:    "never",
		Uniqueness:  "none",
	}

	emailsAttr := &core.Attribute{
		Name: "emails",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "string",
				MultiValued: false,
				Description: "Email addresses for the user.  The value SHOULD be canonicalized by the service provider, e.g., 'bjensen@example.com' instead of 'bjensen@EXAMPLE.COM'. Canonical type values of 'work', 'home', and 'other'.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "display",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, e.g., 'work' or 'home'.",
				Required:    false,
				CanonicalValues: []string{
					"work",
					"home",
					"other",
				},
				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name:        "primary",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred mailing address or primary email address.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: true,
		Description: "Email addresses for the user.  The value SHOULD be canonicalized by the service provider, e.g., 'bjensen@example.com' instead of 'bjensen@EXAMPLE.COM'. Canonical type values of 'work', 'home', and 'other'.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	phoneNumbersAttr := &core.Attribute{
		Name: "phoneNumbers",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "string",
				MultiValued: false,
				Description: "Phone number of the User.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "display",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, e.g., 'work', 'home', 'mobile'.",
				Required:    false,
				CanonicalValues: []string{
					"work",
					"home",
					"mobile",
					"fax",
					"pager",
					"other",
				},
				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name:        "primary",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred phone number or primary phone number.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: true,
		Description: "Phone numbers for the User.  The value SHOULD be canonicalized by the service provider according to the format specified in RFC 3966, e.g., 'tel:+1-201-555-0123'. Canonical type values of 'work', 'home', 'mobile', 'fax', 'pager', and 'other'.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	imsAttr := &core.Attribute{
		Name: "ims",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "string",
				MultiValued: false,
				Description: "Instant messaging address for the User.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "display",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, e.g., 'aim', 'gtalk', 'xmpp'.",
				Required:    false,
				CanonicalValues: []string{
					"aim",
					"gtalk",
					"icq",
					"xmpp",
					"msn",
					"skype",
					"qq",
					"wechat",
					"yahoo",
				},
				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name:        "primary",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred messenger or primary messenger.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: true,
		Description: "Instant messaging addresses for the User.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	photosAttr := &core.Attribute{
		Name: "photos",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "reference",
				MultiValued: false,
				Description: "URL of a photo of the User.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
				ReferenceTypes: []string{
					"external",
				},
			},
			&core.Attribute{
				Name:        "display",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, i.e., 'photo' or 'thumbnail'.",
				Required:    false,
				CanonicalValues: []string{
					"photo",
					"thumbnail",
				},
				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name:        "primary",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred photo or thumbnail.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: true,
		Description: "URLs of photos of the User.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	addressesAttr := &core.Attribute{
		Name: "addresses",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "formatted",
				Type:        "string",
				MultiValued: false,
				Description: "The full mailing address, formatted for display or use with a mailing label.  This attribute MAY contain newlines.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "streetAddress",
				Type:        "string",
				MultiValued: false,
				Description: "The full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information.  This attribute MAY contain newlines.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "locality",
				Type:        "string",
				MultiValued: false,
				Description: "The city or locality component.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "region",
				Type:        "string",
				MultiValued: false,
				Description: "The state or region component.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "postalCode",
				Type:        "string",
				MultiValued: false,
				Description: "The zip code or postal code component.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "country",
				Type:        "string",
				MultiValued: false,
				Description: "The country name component.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, e.g., 'work' or 'home'.",
				Required:    false,
				CanonicalValues: []string{
					"work",
					"home",
					"other",
				},
				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name:        "primary",
				Type:        "boolean",
				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g., the preferred messenger or primary messenger.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
		},
		MultiValued: true,
		Description: "A physical mailing address for this User. Canonical type values of 'work', 'home', and 'other'.  This attribute is a complex type with the following sub-attributes.",
		Required:    false,
		CaseExact:   false,
		Mutability:  "readWrite",
		Returned:    "default",
		Uniqueness:  "none",
	}

	groupsAttr := &core.Attribute{
		Name: "groups",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name:        "value",
				Type:        "string",
				MultiValued: false,
				Description: "The identifier of the User's group.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readOnly",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "$ref",
				Type:        "reference",
				MultiValued: false,
				Description: "The URI of the corresponding 'Group' resource to which the user belongs.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readOnly",
				Returned:    "default",
				Uniqueness:  "none",
				ReferenceTypes: []string{
					"User",
					"Group",
				},
			},
			&core.Attribute{
				Name:        "display",
				Type:        "string",
				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readOnly",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name:        "type",
				Type:        "string",
				MultiValued: false,
				Description: "A label indicating the attribute's function, e.g., 'direct' or 'indirect'.",
				Required:    false,
				CanonicalValues: []string{
					"direct",
					"indirect",
				},
				CaseExact:  false,
				Mutability: "readOnly",
				Returned:   "default",
				Uniqueness: "none",
			},
		},
		MultiValued: true,
		Description: "A list of groups to which the user belongs, either through direct membership, through nested groups, or dynamically calculated.",
		Required:    false,

		CaseExact:  false,
		Mutability: "readOnly",
		Returned:   "default",
		Uniqueness: "none",
	}

	entitlementesAttr := &core.Attribute{
		Name: "entitlements",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name: "value",
				Type: "string",

				MultiValued: false,
				Description: "The value of an entitlement.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "display",
				Type: "string",

				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "type",
				Type: "string",

				MultiValued: false,
				Description: "A label indicating the attribute's function.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "primary",
				Type: "boolean",

				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
		},
		MultiValued: true,
		Description: "A list of entitlements for the User that represent a thing the User has.",
		Required:    false,

		CaseExact:  false,
		Mutability: "readWrite",
		Returned:   "default",
		Uniqueness: "none",
	}

	rolesAttr := &core.Attribute{
		Name: "roles",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name: "value",
				Type: "string",

				MultiValued: false,
				Description: "The value of a role.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "display",
				Type: "string",

				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "type",
				Type: "string",

				MultiValued: false,
				Description: "A label indicating the attribute's function.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "primary",
				Type: "boolean",

				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
		},
		MultiValued: true,
		Description: "A list of roles for the User that collectively represent who the User is, e.g., 'Student', 'Faculty'.",
		Required:    false,

		CaseExact:  false,
		Mutability: "readWrite",
		Returned:   "default",
		Uniqueness: "none",
	}

	x509certificatesAttr := &core.Attribute{
		Name: "x509Cerficates",
		Type: "complex",
		SubAttributes: []*core.Attribute{
			&core.Attribute{
				Name: "value",
				Type: "binary",

				MultiValued: false,
				Description: "The value of an X.509 certificate.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "display",
				Type: "string",

				MultiValued: false,
				Description: "A human-readable name, primarily used for display purposes.  READ-ONLY.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
			&core.Attribute{
				Name: "type",
				Type: "string",

				MultiValued: false,
				Description: "A label indicating the attribute's function.",
				Required:    false,
				CaseExact:   false,
				Mutability:  "readWrite",
				Returned:    "default",
				Uniqueness:  "none",
			},
			&core.Attribute{
				Name: "primary",
				Type: "boolean",

				MultiValued: false,
				Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute.  The primary attribute value 'true' MUST appear no more than once.",
				Required:    false,

				CaseExact:  false,
				Mutability: "readWrite",
				Returned:   "default",
				Uniqueness: "none",
			},
		},
		MultiValued: true,
		Description: "A list of certificates issued to the User.",
		Required:    false,

		CaseExact:  false,
		Mutability: "readWrite",
		Returned:   "default",
		Uniqueness: "none",
	}

	schema.Attributes = []*core.Attribute{
		userNameAttr,
		nameAttr,
		displayNameAttr,
		nickNameAttr,
		profileURLAttr,
		titleAttr,
		userTypeAttr,
		preferredLanguageAttr,
		localeAttr,
		timezoneAttr,
		activeAttr,
		passwordAttr,
		emailsAttr,
		phoneNumbersAttr,
		imsAttr,
		photosAttr,
		addressesAttr,
		groupsAttr,
		entitlementesAttr,
		rolesAttr,
		x509certificatesAttr,
	}

	if errors := validation.Validator.Struct(schema); errors != nil {
		fmt.Println(errors)
		panic("default user schema configuration incorrect")
	}
	// (todo) > mold

	UserSchema = *schema
}
