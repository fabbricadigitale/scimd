package messages

import (
	"github.com/fabbricadigitale/scimd/api"
)

const SearchRequestURI = "urn:ietf:params:scim:api:messages:2.0:SearchRequest"

type SearchRequest struct {
	Schemas []string `json:"schemas"`
	api.Search
}
