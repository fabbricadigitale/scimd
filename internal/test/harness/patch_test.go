package harness

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/server"
)

func TestPatch(t *testing.T) {
	setup()
	defer teardown()

	r, err := os.Open("../../testdata/patch_add_complex.json")
	require.NoError(t, err)

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/v2/Users/2819c223-7f76-453a-919d-ab1234567891", r)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	user := resource.Resource{}
	json.Unmarshal([]byte(rec.Body.String()), &user)

	values := user.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	emails := (*values)["emails"].([]datatype.DataTyper)

	found := false

	for _, email := range emails {
		if email.(datatype.Complex)["value"] == datatype.String("gigi@gmail.com") {
			found = true
			break
		}
	}

	require.True(t, found)
}
