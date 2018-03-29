package harness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/server"
)

func TestUpdateRelationshipAdded(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	id := "e9e30dba-f08f-4109-8486-d5c6a331660a"

	req, _ := http.NewRequest("GET", fmt.Sprintf("/v2/Groups/%s", id), nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	group := resource.Resource{}
	json.Unmarshal([]byte(rec.Body.String()), &group)

	values := group.Values("urn:ietf:params:scim:schemas:core:2.0:Group")
	members := (*values)["members"].([]datatype.DataTyper)
	members = append(members, datatype.Complex{
		"value": "2819c223-7f76-453a-919d-ab1234567891",
		"$ref":  "/v2/Users/2819c223-7f76-453a-919d-ab1234567891",
	})
	(*values)["members"] = members
	group.SetValues("urn:ietf:params:scim:schemas:core:2.0:Group", values)

	body, _ := json.Marshal(&group)

	rec2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("PUT", fmt.Sprintf("/v2/Groups/%s", id), bytes.NewBuffer(body))
	req2.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec2, req2)

	rec3 := httptest.NewRecorder()
	idUser := "2819c223-7f76-453a-919d-ab1234567891"

	req3, _ := http.NewRequest("GET", fmt.Sprintf("/v2/Users/%s", idUser), nil)
	req3.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec3, req3)

	resp := rec3.Body.String()

	usr, _ := ioutil.ReadFile("../../testdata/user_after_group_assignment.json")

	require.JSONEq(t, string(usr), resp)
}

func TestUpdateRelationshipRemoved(t *testing.T) {
	setup()
	defer teardown()

	spc := config.ServiceProviderConfig()
	srv := server.Get(&spc)
	rec := httptest.NewRecorder()

	id := "e9e30dba-f08f-4109-8486-d5c6a331660a"

	req, _ := http.NewRequest("GET", fmt.Sprintf("/v2/Groups/%s", id), nil)
	req.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec, req)

	group := resource.Resource{}
	json.Unmarshal([]byte(rec.Body.String()), &group)

	values := group.Values("urn:ietf:params:scim:schemas:core:2.0:Group")
	members := (*values)["members"].([]datatype.DataTyper)
	members = append(members[:0], members[1:]...)
	(*values)["members"] = members
	group.SetValues("urn:ietf:params:scim:schemas:core:2.0:Group", values)

	body, _ := json.Marshal(&group)

	rec2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("PUT", fmt.Sprintf("/v2/Groups/%s", id), bytes.NewBuffer(body))
	req2.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec2, req2)

	rec3 := httptest.NewRecorder()
	idUser := "2819c223-7f76-453a-919d-ab1234567891"

	req3, _ := http.NewRequest("GET", fmt.Sprintf("/v2/Users/%s", idUser), nil)
	req3.Header.Add("Authorization", aaa)
	srv.ServeHTTP(rec3, req3)

	resp := rec3.Body.String()

	usr, _ := ioutil.ReadFile("../../testdata/resp_user_full_attributes.json")

	require.JSONEq(t, string(usr), resp)
}
