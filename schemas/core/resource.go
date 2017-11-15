package core

type Resource interface {
	GetResourceType() *ResourceType
	GetSchema() *Schema
	GetSchemaExtensions() map[string]*Schema

	GetCommon() *Common
}
