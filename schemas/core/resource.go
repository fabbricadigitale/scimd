package core

// ResourceTyper is the interface implemented by types representing a SCIM resource which embed CommonAttributes and can return the related ResourceType
type ResourceTyper interface {
	ResourceType() *ResourceType
	Common() *CommonAttributes
}
