package core

type Resource interface {
	GetResourceType() *ResourceType
	GetSchemas() *[]string
	GetCommon() *Common
	//
	Get(path string) DataType
	GetExtended(ns string, path string) DataType
	// oppure
	Data() *Complex
	ExtData(ns string) *Complex
}
