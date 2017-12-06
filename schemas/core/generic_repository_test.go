package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenericRepository(t *testing.T) {
	repo1 := GetGenericRepository()
	repo2 := GetGenericRepository()

	// Is it a singleton?
	require.IsType(t, (*repositoryGeneric)(nil), repo1)
	require.IsType(t, (*repositoryGeneric)(nil), repo2)
	require.Implements(t, (*GenericRepository)(nil), repo1)
	require.Implements(t, (*GenericRepository)(nil), repo2)
	require.Exactly(t, repo1, repo2)
}

func TestSchemaRepository(t *testing.T) {
	schemas := GetSchemaRepository()

	// Malformed JSON
	_, err0 := schemas.Add("../../internal/testdata/malformed.json")
	require.Error(t, err0)
	// require.Empty(t, data0)

	// Wrong path
	_, err1 := schemas.Add("WRONG/uschema.json")
	require.Error(t, err1)
	// require.Empty(t, data1)

	// Wrong structure
	_, err2 := schemas.Add("../../internal/testdata/service_provider_config.json")
	require.EqualError(t, err2, "missing identifier")
	// require.Empty(t, data2)

	data3, err3 := schemas.Add("../../internal/testdata/user_schema.json")
	require.NoError(t, err3)
	require.Implements(t, (*Identifiable)(nil), data3)
	require.IsType(t, Schema{}, data3)

	key := "urn:ietf:params:scim:schemas:core:2.0:User"
	schema := schemas.Get(key)

	require.Equal(t, schema.GetIdentifier(), key)

	// (todo): test lock
}

func TestResourceTypeRepository(t *testing.T) {
	// (todo)

	rType := GetResourceTypeRepository()

	// Malformed JSON
	_, err0 := rType.Add("../../internal/testdata/malformed.json")
	require.Error(t, err0)
	// require.Empty(t, data0)

	// Wrong path
	_, err1 := rType.Add("WRONG/urt.json")
	require.Error(t, err1)
	// require.Empty(t, data1)

	// Wrong structure
	_, err2 := rType.Add("../../internal/testdata/service_provider_config.json")
	require.EqualError(t, err2, "missing identifier")
	// require.Empty(t, data2)

	data3, err3 := rType.Add("../../internal/testdata/user.json")
	require.NoError(t, err3)
	require.Implements(t, (*Identifiable)(nil), data3)
	require.IsType(t, ResourceType{}, data3)

	key := "User"
	rT := rType.Get(key)

	require.Equal(t, rT.GetIdentifier(), key)

	// (todo): test lock
}
