package messages

import (
	"encoding/json"
	"net/http"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

//ErrorURI error message urn
const ErrorURI = "urn:ietf:params:scim:api:messages:2.0:Error"

//Error is a struct for wrapping scim error
type Error struct {
	Schemas  []string `json:"schemas"`
	Status   string   `json:"status,required"`
	ScimType string   `json:"scimType,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}

// NewError wraps error in a scim Error struct
func NewError(e error) Error {

	var scimError Error

	switch e.(type) {
	case *json.SyntaxError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidSyntax"
	case *datatype.InvalidDataTypeError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidValue"
	case *json.UnmarshalTypeError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidValue"
	case *api.InvalidPathError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidPath"
	case *api.InvalidFilterError:
		scimError.Status = string(http.StatusBadRequest)
		scimError.ScimType = "invalidFilter"
	case *api.ResourceNotFoundError:
		scimError.Status = "404"
	default:
		scimError.Status = string(http.StatusInternalServerError)
	}

	scimError.Schemas = append(scimError.Schemas, ErrorURI)
	scimError.Detail = e.Error()

	return scimError
}

func (e Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b)
}
