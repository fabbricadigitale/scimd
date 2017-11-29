package core

import (
	"time"
)

// A ScimError is a description of a SCIM error.
type ScimError struct {
	Msg string // description of error
}

func (e *ScimError) Error() string { return e.Msg }

// Meta ...
type Meta struct {
	Location     string     `json:"location" validate:"uri,required"`
	ResourceType string     `json:"resourceType" validate:"required"`
	Created      *time.Time `json:"created" validate:"required"`
	LastModified *time.Time `json:"lastModified" validate:"required"`
	Version      string     `json:"version,omitempty"`
}

// CommonAttributes represents SCIM Common Attributes as per https://tools.ietf.org/html/rfc7643#section-3.1
type CommonAttributes struct {
	Schemas []string `json:"schemas" validate:"gt=0,dive,urn,required"`

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
