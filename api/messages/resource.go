package messages

import (
	"github.com/fabbricadigitale/scimd/schemas/core"
)

// Complex ...
type Complex map[string]interface{}

// Meta ...
type Meta struct {
	ResourceType string
	Created      string
	LastModified string
	Location     string
	Version      string
}

// Common ...
type Common struct {
	ID         string
	ExternalID string
	Meta
}

// Resource ...
type Resource struct {
	Schemas map[string]*core.Schema
	Common
	Attributes map[string]Complex
}

// func (this *Schema) Validate(r *Resource)
