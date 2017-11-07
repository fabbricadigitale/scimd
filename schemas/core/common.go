package core

import (
	"time"
)

// Complex ...
type Complex map[string]interface{}

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
