package core

type Patch struct {
	Supported bool `json:"supported"`
}

type Bulk struct {
	Supported      bool `json:"supported"`
	MaxOperations  int  `json:"maxOperations"`
	MaxPayloadSize int  `json:"maxPayloadSize"`
}

type Filter struct {
	Supported  bool `json:"supported"`
	MaxResults int  `json:"maxResults"`
}

type ChangePassword struct {
	Supported bool `json:"supported"`
}

type Sort struct {
	Supported bool `json:"supported"`
}

type Etag struct {
	Supported bool `json:"supported"`
}

type AuthenticationSchema struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	SpecURI          string `json:"specUri"`
	DocumentationURI string `json:"documentationUri"`
	Type             string `json:"type"`
	Primary          bool   `json:"primary,omitempty"`
}

// ServiceProviderConfig is a structured resource "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"
type ServiceProviderConfig struct {
	ID string `json:"id,omitempty"`
	Common
	DocumentationURI      string                 `json:"documentationUri"`
	Patch                 Patch                  `json:"patch"`
	Bulk                  Bulk                   `json:"bulk"`
	Filter                Filter                 `json:"filter"`
	ChangePassword        ChangePassword         `json:"changePassword"`
	Sort                  Sort                   `json:"sort"`
	Etag                  Etag                   `json:"etag"`
	AuthenticationSchemas []AuthenticationSchema `json:"authenticationSchemas"`
}

var _ Resource = (*ServiceProviderConfig)(nil)
