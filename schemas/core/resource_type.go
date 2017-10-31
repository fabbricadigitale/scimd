package core

type SchemaExtension struct {
	Schema   string `json:"schema"`
	Required bool   `json:"required"`
}

type ResourceType struct {
	Resource
	Name             string            `json:"name"`
	Endpoint         string            `json:"endpoint"`
	Description      string            `json:"description"`
	Schema           string            `json:"schema"`
	SchemaExtensions []SchemaExtension `json:"schemaExtensions"`
}
