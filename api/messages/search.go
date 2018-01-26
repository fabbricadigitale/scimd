package messages

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/api"
	defaults "github.com/mcuadros/go-defaults"
)

// SearchRequestURI is the URI identifies the search query request
const SearchRequestURI = "urn:ietf:params:scim:api:messages:2.0:SearchRequest"

// SearchRequest represents a resource query as per https://tools.ietf.org/html/rfc7644#section-3.4.3
type SearchRequest struct {
	Schemas []string `json:"schemas" validate:"eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest"`
	api.Search
}

// UnmarshalJSON unmarshals an Attribute taking into account defaults
func (s *SearchRequest) UnmarshalJSON(data []byte) error {
	defaults.SetDefaults(s)

	type aliasType SearchRequest
	alias := aliasType(*s)
	// Notice that the promoted UnmarshalJSON of api.Search results in an instance (ie., `s`) missing the Schemas field
	err := json.Unmarshal(data, &alias)
	*s = SearchRequest(alias)

	// So we need to unmarshal the current data as raw
	var parts map[string]json.RawMessage
	if err := json.Unmarshal(data, &parts); err != nil {
		return err
	}

	// Grabbing the raw "schemas" part
	if part, ok := parts["schemas"]; ok {
		if err := json.Unmarshal(part, &s.Schemas); err != nil {
			return err
		}
	}

	return err
}
