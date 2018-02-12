package core

import (
	"time"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// A ScimError is a description of a SCIM error.
type ScimError struct {
	Msg string // description of error
}

func (e *ScimError) Error() string { return e.Msg }

// Meta ...
type Meta struct {
	Location     string     `json:"location,omitempty" validate:"uri"`
	ResourceType string     `json:"resourceType" validate:"required"`
	Created      *time.Time `json:"created,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	Version      string     `json:"version,omitempty"`
}

// CommonAttributes represents SCIM Common Attributes as per https://tools.ietf.org/html/rfc7643#section-3.1
type CommonAttributes struct {
	Schemas []string `json:"schemas" validate:"gt=0,dive,urn,required" mold:"dive,normurn"`

	// Common attributes
	ID         string `json:"id" validate:"excludes=bulkId,required"`
	ExternalID string `json:"externaId,omitempty"`
	Meta       Meta   `json:"meta" validate:"required"`
}

// NewCommon returns a Common filled with schema, resourceType, and ID
func NewCommon(schema, resourceType, ID string) *CommonAttributes {
	return &CommonAttributes{
		Schemas: []string{schema},
		ID:      ID,
		Meta:    Meta{ResourceType: resourceType},
	}
}

// Common returns CommonAttributes of a SCIM resource
func (c *CommonAttributes) Common() *CommonAttributes {
	return c
}

// ResourceType returns the ResourceType of a SCIM resource
func (c *CommonAttributes) ResourceType() *ResourceType {
	return GetResourceTypeRepository().Get(c.Meta.ResourceType)
}

var comAttrs Attributes

// Commons returns Attributes are considered to be part of every base resource schema and do not use their own "schemas" URI.
func Commons() Attributes {
	return comAttrs
}

func init() {
	comAttrs = Attributes{
		&Attribute{
			Name: "schemas",
			Type: datatype.StringType,
			// SubAttributes
			MultiValued: true,
			Description: "An array of Strings containing URIs that are used to indicate the namespaces of the SCIM schemas",
			Required:    true,
			// CanonicalValues
			CaseExact:  true,
			Mutability: schemas.MutabilityReadWrite,
			Returned:   schemas.ReturnedAlways,
			Uniqueness: schemas.UniquenessNone,
			// ReferenceTypes
		},
		&Attribute{
			Name: "id",
			Type: datatype.StringType,
			// SubAttributes
			MultiValued: false,
			Description: "A unique identifier for a SCIM resource as defined by the service provider",
			Required:    true,
			// CanonicalValues
			CaseExact:  true,
			Mutability: schemas.MutabilityReadOnly,
			Returned:   schemas.ReturnedAlways,
			Uniqueness: schemas.UniquenessServer,
			// ReferenceTypes
		},
		&Attribute{
			Name: "externalId",
			Type: datatype.StringType,
			// SubAttributes
			MultiValued: false,
			Description: "A String that is an identifier for the resource as defined by the provisioning client",
			Required:    false,
			// CanonicalValues
			CaseExact:  true,
			Mutability: schemas.MutabilityReadOnly,
			Returned:   schemas.ReturnedDefault,
			Uniqueness: schemas.UniquenessNone,
			// ReferenceTypes
		},
		&Attribute{
			Name: "meta",
			Type: datatype.ComplexType,
			SubAttributes: Attributes{
				&Attribute{
					Name: "resourceType",
					Type: datatype.StringType,
					// SubAttributes
					MultiValued: false,
					Description: "The name of the resource type of the resource",
					Required:    false,
					// CanonicalValues
					CaseExact:  true,
					Mutability: schemas.MutabilityReadOnly,
					// NOTE > We need to enforce at least this (sub)attribute to be present in projections - See #55
					Returned:   schemas.ReturnedAlways,
					Uniqueness: schemas.UniquenessNone,
					// ReferenceTypes
				},
				&Attribute{
					Name: "created",
					Type: datatype.DateTimeType,
					// SubAttributes
					MultiValued: false,
					Description: "The DateTime that the resource was added to the service provider",
					Required:    false,
					// CanonicalValues
					CaseExact:  false,
					Mutability: schemas.MutabilityReadOnly,
					Returned:   schemas.ReturnedDefault,
					Uniqueness: schemas.UniquenessNone,
					// ReferenceTypes
				},
				&Attribute{
					Name: "lastModified",
					Type: datatype.DateTimeType,
					// SubAttributes
					MultiValued: false,
					Description: "The most recent DateTime that the details of this resource were updated at the service provider",
					Required:    false,
					// CanonicalValues
					CaseExact:  false,
					Mutability: schemas.MutabilityReadOnly,
					Returned:   schemas.ReturnedDefault,
					Uniqueness: schemas.UniquenessNone,
					// ReferenceTypes
				},
				&Attribute{
					Name: "location",
					Type: datatype.StringType,
					// SubAttributes
					MultiValued: false,
					Description: "The URI of the resource being returned",
					Required:    false,
					// CanonicalValues
					CaseExact:  true,
					Mutability: schemas.MutabilityReadOnly,
					Returned:   schemas.ReturnedDefault,
					Uniqueness: schemas.UniquenessNone,
					// ReferenceTypes
				},
				&Attribute{
					Name: "version",
					Type: datatype.StringType,
					// SubAttributes
					MultiValued: false,
					Description: "The version of the resource being returned",
					Required:    false,
					// CanonicalValues
					CaseExact:  true,
					Mutability: schemas.MutabilityReadOnly,
					Returned:   schemas.ReturnedDefault,
					Uniqueness: schemas.UniquenessNone,
					// ReferenceTypes
				},
			},
			MultiValued: false,
			Description: "A complex attribute containing resource metadata",
			Required:    false,
			// CanonicalValues
			CaseExact:  false,
			Mutability: schemas.MutabilityReadOnly,
			Returned:   schemas.ReturnedDefault,
			Uniqueness: schemas.UniquenessNone,
			// ReferenceTypes
		},
	}
}
