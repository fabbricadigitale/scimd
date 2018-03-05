package integration

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/api/create"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

func TestCreate(t *testing.T) {
	setupDB()
	setup()
	defer teardownDB()

	retRes, err := create.Resource(adapter, resTypeRepo.Pull("User"), &res)
	require.Nil(t, err)
	require.NotNil(t, retRes)

	// check urn:ietf:params:scim:schemas:core:2.0:User.userName
	r := retRes.(*resource.Resource)
	values := r.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	userName := (*values)["userName"]

	expValues := res.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	expUserName := (*expValues)["userName"]

	require.Equal(t, expUserName, userName)

	// check urn:ietf:params:scim:schemas:extension:enterprise:2.0:User.employeeNumber
	extensionValues := r.Values("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
	employeeNumber := (*extensionValues)["employeeNumber"]
	require.Equal(t, datatype.String("701984"), employeeNumber)

	require.Equal(t, res.ID, r.ID)
	require.Equal(t, res.Meta.Version, r.Meta.Version)
}

func TestUniqueness(t *testing.T) {
	setupDB()
	setup()
	defer teardownDB()

	retRes, err := create.Resource(adapter, resTypeRepo.Pull("User"), &res)
	require.Nil(t, err)
	require.NotNil(t, retRes)

	retRes, err = create.Resource(adapter, resTypeRepo.Pull("User"), &res)
	require.Error(t, err)
	require.Nil(t, retRes)
}
