package messages

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func TestErrorWrapper(t *testing.T) {

	// Unexpected Data Type
	var s = "randomvalue"
	_, err := core.NewDataType(s)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURN),
		Status:   string(http.StatusBadRequest),
		ScimType: "invalidValue",
		Detail:   err.Error(),
	}, ErrorWrapper(err))

	// Unmarshal type error
	type Prova struct {
		Num string `json:"num"`
	}
	p := Prova{}
	byt := `{"num": 13}`

	err = json.Unmarshal([]byte(byt), &p)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURN),
		Status:   string(http.StatusBadRequest),
		ScimType: "invalidValue",
		Detail:   err.Error(),
	}, ErrorWrapper(err))

	// Error in json syntax
	byt = `{num: 17}`

	err = json.Unmarshal([]byte(byt), &p)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURN),
		Status:   string(http.StatusBadRequest),
		ScimType: "invalidSyntax",
		Detail:   err.Error(),
	}, ErrorWrapper(err))

	// ErrorWrapper returns a correct json format
	e := ErrorWrapper(err)
	byt2, _ := json.Marshal(e)

	require.JSONEq(t, `{
		"schemas":["urn:ietf:params:scim:api:messages:2.0:Error"],
		"status":"Ɛ",
		"scimType":"invalidSyntax",
		"detail":"invalid character 'n' looking for beginning of object key string"
		}`, string(byt2))

	// todo InvalidPathError

	// todo InvalidFilterError

}
