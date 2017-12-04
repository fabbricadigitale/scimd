package messages

import (
	"github.com/fabbricadigitale/scimd/api"
)

// SearchRequestURI is the URI identifies the search query request
const SearchRequestURI = "urn:ietf:params:scim:api:messages:2.0:SearchRequest"

// SearchRequest represents a resource query as per https://tools.ietf.org/html/rfc7644#section-3.4.3
type SearchRequest struct {
	Schemas []string `json:"schemas" validate:"eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:SearchRequest"`
	api.Search
}
