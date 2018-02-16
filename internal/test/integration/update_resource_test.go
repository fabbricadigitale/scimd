package integration

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/api/update"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {

	notExistingID := "abcdefgh-xxxx-yyyy-zzzz-jkilmnopqrst"

	id := "2819c223-7f76-453a-919d-ab1234567891"
	dat, err := ioutil.ReadFile("../../testdata/enterprise_user_resource_1.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	retRes, err := update.Resource(adapter, resTypeRepo.Pull("User"), notExistingID, res)
	require.Error(t, err)

	res = &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	retRes, err = update.Resource(adapter, resTypeRepo.Pull("User"), id, res)
	require.Nil(t, err)
	r := retRes.(*resource.Resource)
	require.NotNil(t, r)
	require.Equal(t, res.Meta.Version, r.Meta.Version)

}
