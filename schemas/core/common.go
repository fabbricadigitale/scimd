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
	Location     string     `json:"location"`
	ResourceType string     `json:"resourceType"`
	Created      *time.Time `json:"created"`
	LastModified *time.Time `json:"lastModified"`
	Version      string     `json:"version"`
}

// Common ...
type Common struct {
	Schemas []string `json:"schemas"`

	// Common attributes
	ID         string `json:"id"`
	ExternalID string `json:"externaId,omitempty"`
	Meta       Meta   `json:"meta"`
}
