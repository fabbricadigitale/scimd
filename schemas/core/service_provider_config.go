package core

import (
	"encoding/json"
	"time"

	defaults "github.com/mcuadros/go-defaults"
)

type supported struct {
	Supported bool `json:"supported"`
}

type bulk struct {
	Supported      bool `json:"supported"`
	MaxOperations  int  `json:"maxOperations" validate:"gte=0"`
	MaxPayloadSize int  `json:"maxPayloadSize" validate:"gte=0"`
}

type filter struct {
	Supported  bool `json:"supported"`
	MaxResults int  `json:"maxResults" validate:"gte=0"`
}

// AuthenticationScheme is ...
type AuthenticationScheme struct {
	Type             string `json:"type" validate:"required,eq=oauth|eq=oauth2|eq=oauthbearertoken|eq=httpbasic|eq=httpdigest"`
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description" validate:"required"`
	SpecURI          string `json:"specUri,omitempty" validate:"omitempty,uri"`
	DocumentationURI string `json:"documentationUri,omitempty" validate:"omitempty,uri"`
	Primary          bool   `json:"primary,omitempty" default:"false"`
}

// ServiceProviderConfig is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"
type ServiceProviderConfig struct {
	Schemas               []string               `json:"schemas" validate:"gt=0,dive,urn,required" mold:"dive,normurn"`
	Meta                  Meta                   `json:"meta" validate:"required"`
	DocumentationURI      string                 `json:"documentationUri,omitempty" validate:"omitempty,uri"`
	Patch                 supported              `json:"patch" validate:"required"`
	Bulk                  bulk                   `json:"bulk" validate:"required"`
	Filter                filter                 `json:"filter" validate:"required"`
	ChangePassword        supported              `json:"changePassword" validate:"required"`
	Sort                  supported              `json:"sort" validate:"required"`
	Etag                  supported              `json:"etag" validate:"required"`
	AuthenticationSchemes []AuthenticationScheme `json:"authenticationSchemes" validate:"required"`
}

// ServiceProviderConfigURI is the Service Provider Configuration schema used by ServiceProviderConfig
const ServiceProviderConfigURI = "urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"

// NewServiceProviderConfig returns a new ServiceProviderConfig filled with defaults
func NewServiceProviderConfig() *ServiceProviderConfig {
	now := time.Now()
	spc := &ServiceProviderConfig{
		Schemas: []string{ServiceProviderConfigURI},
		Meta: Meta{
			ResourceType: "ServiceProviderConfig",
			Created:      &now,
			LastModified: &now,
		},
	}
	defaults.SetDefaults(spc)

	return spc
}

// UnmarshalJSON unmarshals an Attribute taking into account defaults
func (spc *ServiceProviderConfig) UnmarshalJSON(data []byte) error {
	defaults.SetDefaults(spc)

	type aliasType ServiceProviderConfig
	alias := aliasType(*spc)
	err := json.Unmarshal(data, &alias)

	*spc = ServiceProviderConfig(alias)
	return err
}
