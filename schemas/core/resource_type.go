package core

// SchemaExtension ...
type SchemaExtension struct {
	Schema   string `json:"schema" validate:"urn,required"`
	Required bool   `json:"required" validate:"required"`
}

// ResourceType is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
type ResourceType struct {
	Common
	Name             string            `json:"name" validate:"required"`
	Endpoint         string            `json:"endpoint" validate:"startswith=/,required"`
	Description      string            `json:"description,omitempty"`
	Schema           string            `json:"schema" validate:"urn,required"`
	SchemaExtensions []SchemaExtension `json:"schemaExtensions,omitempty"`
}

// ResourceTypeURI is the Resource Type Configuration schema used by ResourceType
const ResourceTypeURI = "urn:ietf:params:scim:schemas:core:2.0:ResourceType"

// NewResourceType returns a ResourceType filled with min values set which identify a particular schema and resourceType (eg. User)
func NewResourceType(schema, resourceType string) *ResourceType {
	return &ResourceType{
		Common: *NewCommon(ResourceTypeURI, "ResourceType", resourceType),
		Schema: schema,
		Name:   resourceType,
	}
}

var _ Resource = (*ResourceType)(nil)

// GetIdentifier ...
func (rt ResourceType) GetIdentifier() string {
	return rt.Name
}
