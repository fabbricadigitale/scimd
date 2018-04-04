package integration

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/api/patch"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

func TestPatchAdd(t *testing.T) {
	log.Println("TestUpdate")
	setupDB()
	defer teardownDB()

	id := "2819c223-7f76-453a-919d-ab1234567891"

	res, err := patch.Resource(adapter, resTypeRepo.Pull("User"), id, "add", "emails", datatype.Complex{"value": "gigi@gmail.com", "type": "private"})
	require.NoError(t, err)
	require.NotNil(t, res)

	values := res.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	e := (*values)["emails"]
	emails := e.([]datatype.DataTyper)
	newValue := emails[2].(*datatype.Complex)

	require.Equal(t, datatype.String("gigi@gmail.com"), (*newValue)["value"])
	require.Equal(t, datatype.String("private"), (*newValue)["type"])
}

func TestPatchRemovePull(t *testing.T) {
	log.Println("TestUpdate")
	setupDB()
	defer teardownDB()

	id := "2819c223-7f76-453a-919d-ab1234567891"

	res, err := patch.Resource(adapter, resTypeRepo.Pull("User"), id, "remove", "emails", datatype.Complex{"value": "tiffy@fork.org", "type": "home"})
	require.NoError(t, err)
	require.NotNil(t, res)

	values := res.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	e := (*values)["emails"]
	emails := e.([]datatype.DataTyper)
	value := emails[0].(*datatype.Complex)

	require.Equal(t, 1, len(emails))
	require.NotEqual(t, datatype.String("tiffy@fork.org"), (*value)["value"])
	require.NotEqual(t, datatype.String("home"), (*value)["type"])

}

func TestPatchRemoveUnset(t *testing.T) {
	log.Println("TestUpdate")
	setupDB()
	defer teardownDB()

	id := "2819c223-7f76-453a-919d-ab1234567891"

	res, err := patch.Resource(adapter, resTypeRepo.Pull("User"), id, "remove", "emails", nil)
	require.NoError(t, err)
	require.NotNil(t, res)

	values := res.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	emails := (*values)["emails"]

	require.Nil(t, emails)
}

func TestPatchReplace(t *testing.T) {

	log.Println("TestUpdate")
	setupDB()
	defer teardownDB()

	id := "2819c223-7f76-453a-919d-ab1234567891"

	res, err := patch.Resource(adapter, resTypeRepo.Pull("User"), id, "replace", "locale", "it-IT")
	require.NoError(t, err)
	require.NotNil(t, res)

	values := res.(*resource.Resource).Values("urn:ietf:params:scim:schemas:core:2.0:User")
	locale := (*values)["locale"]

	require.Equal(t, datatype.String("it-IT"), locale)

}
