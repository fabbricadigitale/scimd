package messages

import (
	"encoding/json"
	"net/http"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

const ErrorURN = "urn:ietf:params:scim:api:messages:2.0:Error"

type Error struct {
	Schemas  []string `json:"schemas"`
	Status   string   `json:"status"`
	ScimType string   `json:"scimType,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}

// ErrorWrapper wraps error in a scim Error struct
func ErrorWrapper(e error) Error {

	var scimError Error

	switch e.(type) {
	case *core.DataTypeError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidType"
	case *json.UnmarshalTypeError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidType"
	case *api.InvalidPathError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidPath"
	case *api.InvalidFilterError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidFilter"
	default:
		scimError.Status = string(http.StatusInternalServerError)
	}

	scimError.Schemas = append(scimError.Schemas, ErrorURN)
	scimError.Detail = e.Error()

	return scimError
}
