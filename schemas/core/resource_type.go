package core

// SchemaExtension ...
type SchemaExtension struct {
	Schema   string `json:"schema" validate:"required"`
	Required bool   `json:"required" validate:"required"`
}

// ResourceType is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
type ResourceType struct {
	Common
	Name             string            `json:"name" validate:"required"`
	Endpoint         string            `json:"endpoint" validate:"startswith=/,required"`
	Description      string            `json:"description"`
	Schema           string            `json:"schema" validate:"uri,required"`
	SchemaExtensions []SchemaExtension `json:"schemaExtensions"`
}

var _ Resource = (*ResourceType)(nil)

// GetIdentifier ...
func (rt ResourceType) GetIdentifier() string {
	return rt.Name
}
