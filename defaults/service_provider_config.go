package defaults

import (
	"time"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/validation"
	"github.com/fabbricadigitale/scimd/version"
)

// ServiceProviderConfig is the default service provider configuration
var ServiceProviderConfig core.ServiceProviderConfig

var v = version.GenerateVersion(true, time.Now().String())

func init() {
	ServiceProviderConfig = *core.NewServiceProviderConfig()
	ServiceProviderConfig.Meta.Version = v
	ServiceProviderConfig.Meta.Location = "/v2/ServiceProviderConfigs"
	ServiceProviderConfig.DocumentationURI = "/help/scim.html"
	ServiceProviderConfig.Patch.Supported = false
	ServiceProviderConfig.Bulk.Supported = false
	ServiceProviderConfig.Bulk.MaxOperations = 1000
	ServiceProviderConfig.Bulk.MaxPayloadSize = 1048576
	ServiceProviderConfig.Filter.Supported = true
	ServiceProviderConfig.Filter.MaxResults = 200
	ServiceProviderConfig.ChangePassword.Supported = false
	ServiceProviderConfig.Sort.Supported = true
	ServiceProviderConfig.Etag.Supported = true

	basicAuth := core.AuthenticationScheme{
		Type:             "httpbasic",
		Name:             "HTTP Basic",
		Description:      "Authentication scheme using the HTTP Basic Standard",
		Primary:          false,
		DocumentationURI: "/help/httpBasic.html",
		SpecURI:          "http://www.rfc-editor.org/info/rfc2617",
	}
	ServiceProviderConfig.AuthenticationSchemes = []core.AuthenticationScheme{
		basicAuth,
	}

	if errors := validation.Validator.Struct(ServiceProviderConfig); errors != nil {
		panic("service provider default configuration incorrect")
	}
	// (todo) > mold
}
