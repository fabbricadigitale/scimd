package core

// SchemaExtension ...
type SchemaExtension struct {
	Schema   string `json:"schema"`
	Required bool   `json:"required"`
}

// ResourceType ...
type ResourceType struct {
	Common
	Name             string            `json:"name"`
	Endpoint         string            `json:"endpoint"`
	Description      string            `json:"description"`
	Schema           string            `json:"schema"`
	SchemaExtensions []SchemaExtension `json:"schemaExtensions"`
}

// GetIdentifier ...
func (rt ResourceType) GetIdentifier() string {
	return rt.ID
}
