package core

// SchemaExtension ...
type SchemaExtension struct {
	Schema   string `json:"schema" validate:"required"`
	Required bool   `json:"required" validate:"required"`
}

// ResourceType ...
type ResourceType struct {
	Common
	Name             string            `json:"name" validate:"required"`
	Endpoint         string            `json:"endpoint" validate:"startswith=/,required"`
	Description      string            `json:"description"`
	Schema           string            `json:"schema" validate:"uri,required"`
	SchemaExtensions []SchemaExtension `json:"schemaExtensions"`
}

// GetIdentifier ...
func (rt ResourceType) GetIdentifier() string {
	return rt.Name
}
