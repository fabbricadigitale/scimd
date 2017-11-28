package core

import (
	"encoding/json"

	defaults "github.com/mcuadros/go-defaults"
)

type patch struct {
	Supported bool `json:"supported" validate:"required"`
}

type bulk struct {
	Supported      bool `json:"supported" validate:"required"`
	MaxOperations  int  `json:"maxOperations" validate:"required"`
	MaxPayloadSize int  `json:"maxPayloadSize" validate:"required"`
}

type filter struct {
	Supported  bool `json:"supported" validate:"required"`
	MaxResults int  `json:"maxResults" validate:"required"`
}
type changePassword struct {
	Supported bool `json:"supported" validate:"required"`
}

type sort struct {
	Supported bool `json:"supported" validate:"required"`
}

type etag struct {
	Supported bool `json:"supported" validate:"required"`
}

type authenticationScheme struct {
	Type             string `json:"type" validate:"required,eq=oauth|eq=oauth2|eq=oauthbearertoken|eq=httpbasic|eq=httpdigest"`
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description" validate:"required"`
	SpecURI          string `json:"specUri,omitempty" validate:"omitempty,uri"`
	DocumentationURI string `json:"documentationUri,omitempty" validate:"omitempty,uri"`
	Primary          bool   `json:"primary,omitempty" default:"false"`
}

// ServiceProviderConfig is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"
type ServiceProviderConfig struct {
	Common
	DocumentationURI      string                 `json:"documentationUri,omitempty" validate:"omitempty,uri"`
	Patch                 patch                  `json:"patch" validate:"required"`
	Bulk                  bulk                   `json:"bulk" validate:"required"`
	Filter                filter                 `json:"filter" validate:"required"`
	ChangePassword        changePassword         `json:"changePassword" validate:"required"`
	Sort                  sort                   `json:"sort" validate:"required"`
	Etag                  etag                   `json:"etag" validate:"required"`
	AuthenticationSchemes []authenticationScheme `json:"authenticationSchemes" validate:"required"`
}

// ServiceProviderConfigURI is the Service Provider Configuration schema used by ServiceProviderConfig
const ServiceProviderConfigURI = "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"

// NewServiceProviderConfig returns a new ServiceProviderConfig filled with defaults
func NewServiceProviderConfig() *ServiceProviderConfig {
	spc := &ServiceProviderConfig{
		Common: *NewCommon(ServiceProviderConfigURI, "ServiceProviderConfig", ""),
	}
	defaults.SetDefaults(spc)
	return spc
}

var _ ResourceTyper = (*ServiceProviderConfig)(nil)

// UnmarshalJSON unmarshals an Attribute taking into account defaults
func (spc *ServiceProviderConfig) UnmarshalJSON(data []byte) error {
	defaults.SetDefaults(spc)

	type aliasType ServiceProviderConfig
	alias := aliasType(*spc)
	err := json.Unmarshal(data, &alias)

	*spc = ServiceProviderConfig(alias)
	return err
}
