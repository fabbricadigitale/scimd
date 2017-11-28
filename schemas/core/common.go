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

// Common ...
type Common struct {
	Schemas []string `json:"schemas" validate:"gt=0,dive,urn,required"`

	// Common attributes
	ID         string `json:"id" validate:"excludes=bulkId,required"`
	ExternalID string `json:"externaId,omitempty"`
	Meta       Meta   `json:"meta" validate:"required"`
}

// NewCommon returns a Common filled with schema, resourceType, and ID
func NewCommon(schema, resourceType, ID string) *Common {
	return &Common{
		Schemas: []string{schema},
		ID:      ID,
		Meta:    Meta{ResourceType: resourceType},
	}
}

func (c *Common) GetCommon() *Common {
	return c
}

func (c *Common) ResourceType() *ResourceType {
	return GetResourceTypeRepository().Get(c.Meta.ResourceType)
}
