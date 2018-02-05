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

	id := "2819c223-7f76-453a-919d-ab1234567891"
	dat, err := ioutil.ReadFile("../../testdata/enterprise_user_resource_1.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	err = update.Resource(adapter, id, res)

	retRes, err := adapter.Get(resTypeRepo.Get("User"), res.ID, res.Meta.Version, nil)
	require.Nil(t, err)
	require.NotNil(t, retRes)
	require.Equal(t, res.Meta.Version, retRes.Meta.Version)

}
