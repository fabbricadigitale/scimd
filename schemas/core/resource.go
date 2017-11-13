package core

type Resource interface {
	GetResourceType() *ResourceType
	GetSchema() *Schema
	GetSchemaExtensions() map[string]*Schema

	GetCommon() *Common

	// TODO: how to extend this interface to handle both structured and mapped resources?
	// Data() *Complex
	// ExtensionsData(ns string) *Complex
}
