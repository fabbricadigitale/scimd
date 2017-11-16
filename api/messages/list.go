package messages

import (
	"github.com/fabbricadigitale/scimd/schemas/core/resource"
)

const ListResponseURN = "urn:ietf:params:scim:api:messages:2.0:ListResponse"

// ListResponse ...
type ListResponse struct {
	Schemas      []string             `json:"schemas"`
	TotalResults int                  `json:"totalResults"`
	ItemsPerPage int                  `json:"itemsPerPage"`
	StartIndex   int                  `json:"startIndex"`
	Resources    []*resource.Resource `json:"Resources"`
}
