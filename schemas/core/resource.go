package core

// ResourceTyper is the interface implemented by types that can hold Common resource attribute and are can return it ResourceType
type ResourceTyper interface {
	ResourceType() *ResourceType
	GetCommon() *Common
}
