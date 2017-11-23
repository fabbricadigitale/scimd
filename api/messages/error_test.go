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
	byt := []byte(`{"num": 13}`)

	err = json.Unmarshal(byt, &p)

	require.Equal(t, Error{
		Schemas:  append([]string{}, ErrorURN),
		Status:   string(http.StatusBadRequest),
		ScimType: "invalidValue",
		Detail:   err.Error(),
	}, ErrorWrapper(err))

}
