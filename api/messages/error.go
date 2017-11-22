package messages

import (
	"net/http"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

const ErrorURN = "urn:ietf:params:scim:api:messages:2.0:Error"

type Error struct {
	Schemas  []string `json:"schemas"`
	Status   string   `json:"status"`
	ScimType string   `json:"scimType,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}

func ErrorWrapper(e error) Error {

	var scimError Error

	switch e.(type) {
	case *core.DataTypeError:

		scimError.Schemas = append(scimError.Schemas, ErrorURN)
		scimError.Status = string(http.StatusBadRequest)
		scimError.Detail = e.Error()
		scimError.ScimType = "invalidType"
	}

	return scimError
}
