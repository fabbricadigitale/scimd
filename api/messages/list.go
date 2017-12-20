package messages

import (
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// ListResponseURI is the URI identifies the list of responses
const ListResponseURI = "urn:ietf:params:scim:api:messages:2.0:ListResponse"

// ListResponse represents a list of responses as per http://myfabbrica.fabbricadigitale.it
type ListResponse struct {
	Schemas      []string             `json:"schemas" validate:"eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:ListResponse"`
	TotalResults int                  `json:"totalResults" validate:"required"`
	ItemsPerPage int                  `json:"itemsPerPage" validate:"required"`    // we set this always required even if the RFC 7644 specifies that it's required due to pagination
	StartIndex   int                  `json:"startIndex" validate:"gt=0,required"` // we set this always required even if the RFC 7644 specifies that it's required due to pagination
	Resources    []*resource.Resource `json:"Resources" validate:"required"`
}
