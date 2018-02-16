package messages

import (
	"encoding/json"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"
)

func TestNewError(t *testing.T) {

	// Unexpected Data Type
	var s = "randomvalue"
	_, err := datatype.New(s)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURI),
		Status:   400, /* (fixme) why http.StatusBadRequest gives weird chars */
		ScimType: "invalidValue",
		Detail:   err.Error(),
	}, NewError(err))

	// Unmarshal type error
	type WrongValueAttr struct {
		Num string `json:"num"`
	}
	p := WrongValueAttr{}
	byt := `{"num": 13}`

	err = json.Unmarshal([]byte(byt), &p)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURI),
		Status:   400, /* http.StatusBadRequest */
		ScimType: "invalidValue",
		Detail:   err.Error(),
	}, NewError(err))

	// Error in json syntax
	byt = `{num: 17}`

	err = json.Unmarshal([]byte(byt), &p)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURI),
		Status:   400, /* http.StatusBadRequest */
		ScimType: "invalidSyntax",
		Detail:   err.Error(),
	}, NewError(err))

	// NewError returns a correct json format
	e := NewError(err)
	byt2, _ := json.Marshal(e)

	require.JSONEq(t, `{
		"schemas":["urn:ietf:params:scim:api:messages:2.0:Error"],
		"status":400,
		"scimType":"invalidSyntax",
		"detail":"invalid character 'n' looking for beginning of object key string"
		}`, string(byt2))

	// todo InvalidPathError

	// todo InvalidFilterError

}
