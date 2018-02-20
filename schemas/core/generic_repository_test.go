package core

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var m = map[string]string{
	"errortype": "missing identifier",
}

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

func TestClean(t *testing.T) {
	// Schema Repostory 
	repo := GetSchemaRepository()

	data1, _ := repo.PushFromFile("../../internal/testdata/user_schema.json")
	require.IsType(t, Schema{}, data1)
	data2, _ := repo.PushFromFile("../../internal/testdata/enterprise_user_schema.json")
	require.IsType(t, Schema{}, data2)

	list := repo.List()
	require.Len(t, list, 2)

	repo.Clean()
	list = repo.List()
	require.Len(t, list, 0)

	// ResourceType Repository

	repo2 := GetResourceTypeRepository()

	data3, _ := repo2.PushFromFile("../../internal/testdata/user_schema.json")
	require.IsType(t, ResourceType{}, data3)
	data4, _ := repo2.PushFromFile("../../internal/testdata/enterprise_user_schema.json")
	require.IsType(t, ResourceType{}, data4)

	list2 := repo2.List()
	require.Len(t, list2, 2)

	repo2.Clean()
	list2 = repo2.List()
	require.Len(t, list2, 0)
}

func TestSchemaRepository(t *testing.T) {
	schemas := GetSchemaRepository()

	// PushFromFile - Malformed JSON
	x0, err0 := schemas.PushFromFile("../../internal/testdata/malformed.json")
	require.Error(t, err0)
	require.Zero(t, x0)

	// PushFromData - Malformed JSON
	malformed := `{"malformed": "json",}`
	x2, err5 := schemas.PushFromData([]byte(malformed))
	require.Error(t, err5)
	require.Zero(t, x2)

	// Push - Malformed JSON
	x3, err10 := schemas.Push(x0)
	require.Error(t, err10)
	require.Zero(t, x3)

	//PushFromFile - Wrong path
	x1, err1 := schemas.PushFromFile("WRONG/uschema.json")
	require.Error(t, err1)
	require.Zero(t, x1)

	// PushFromFile - Wrong structure (ie., missing ID) - Returns it but do not stores it
	_, err2 := schemas.PushFromFile("../../internal/testdata/service_provider_config.json")
	require.EqualError(t, err2, "missing identifier")
	require.Equal(t, 0, len(schemas.List()))

	// PushFromData - Wrong structure (ie., missing ID) - Returns it but do not stores it
	data, _ := json.Marshal(m)
	_, err4 := schemas.PushFromData(data)
	require.EqualError(t, err4, "missing identifier")
	require.Equal(t, 0, len(schemas.List()))

	// Push - Wrong structure (ie., missing ID) - Returns it but do not stores it
	var missingIDSchema Schema
	_, err7 := schemas.Push(missingIDSchema)
	require.EqualError(t, err7, "missing identifier")
	require.Equal(t, 0, len(schemas.List()))

	// PushFromFile - Successful loading of a schema from file
	data3, err3 := schemas.PushFromFile("../../internal/testdata/user_schema.json")
	require.NoError(t, err3)
	require.Implements(t, (*Identifiable)(nil), data3)
	require.IsType(t, Schema{}, data3)

	key := "urn:ietf:params:scim:schemas:core:2.0:User"
	schema := schemas.Pull(key)

	require.Equal(t, schema.GetIdentifier(), key)
	require.Equal(t, 1, len(schemas.List()))

	// PushFromData - Successful loading of a schema from bytes
	byt, _ := ioutil.ReadFile("../../internal/testdata/enterprise_user_schema.json")
	data5, err6 := schemas.PushFromData(byt)
	require.NoError(t, err6)
	require.Implements(t, (*Identifiable)(nil), data5)
	require.IsType(t, Schema{}, data5)

	key2 := "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"
	schema = schemas.Pull(key2)

	require.Equal(t, schema.GetIdentifier(), key2)
	require.Equal(t, 2, len(schemas.List()))

	ktm := &Schema{}
	if err := json.Unmarshal(byt, ktm); err != nil {
		t.Fatalf("error unmarshalling")
	}
	ktm.ID = "urn:ietf:params:scim:schemas:core:2.0:Shop"
	data7, err8 := schemas.Push(*ktm)
	require.NoError(t, err8)
	require.Implements(t, (*Identifiable)(nil), data7)
	require.IsType(t, Schema{}, data7)

	key3 := "urn:ietf:params:scim:schemas:core:2.0:Shop"
	schema = schemas.Pull(key3)

	require.Equal(t, schema.GetIdentifier(), key3)
	require.Equal(t, 3, len(schemas.List()))

	// (todo) > test lock

	schemas.Clean() // teardown
}

func TestResourceTypeRepository(t *testing.T) {
	rType := GetResourceTypeRepository()

	// PushFromFile - Malformed JSON
	x0, err0 := rType.PushFromFile("../../internal/testdata/malformed.json")
	require.Error(t, err0)
	require.Zero(t, x0)
	// require.Empty(t, data0)

	// PushFromData - Malformed JSON
	malformed := `{"malformed": "json",}`
	x2, err4 := rType.PushFromData([]byte(malformed))
	require.Error(t, err4)
	require.Zero(t, x2)

	// Push - Malformed JSON
	x3, err9 := rType.Push(x0)
	require.Error(t, err9)
	require.Zero(t, x3)

	// PushFromFile - Wrong path
	_, err1 := rType.PushFromFile("WRONG/urt.json")
	require.Error(t, err1)
	// require.Empty(t, data1)

	// PushFromFile - Wrong structure (ie., missing ID) - Returns it but do not stores it
	_, err2 := rType.PushFromFile("../../internal/testdata/service_provider_config.json")
	require.EqualError(t, err2, "missing identifier")
	require.Equal(t, 0, len(rType.List()))
	// require.Empty(t, data2)

	// PushFromData - Wrong structure (ie., missing ID) - Returns it but do not stores it
	data, _ := json.Marshal(m)
	_, err5 := rType.PushFromData(data)
	require.EqualError(t, err5, "missing identifier")
	require.Equal(t, 0, len(rType.List()))

	// Push - Wrong structure (ie., missing ID) - Returns it but do not stores it
	var missingIDResType ResourceType
	missingIDResType.Endpoint = "/User"
	missingIDResType.Description = "User Account "
	missingIDResType.Schema = "urn:ietf:params:scim:schemas:core:2.0:User"
	commons := NewCommon("urn:ietf:params:scim:schemas:core:2.0:ResourceType", "ResourceType", "")
	missingIDResType.CommonAttributes = *commons
	_, err7 := rType.Push(missingIDResType)
	require.EqualError(t, err7, "missing identifier")
	require.Equal(t, 0, len(rType.List()))

	// PushFromFile - Successful loading of a schema from file
	data3, err3 := rType.PushFromFile("../../internal/testdata/user.json")
	require.NoError(t, err3)
	require.Implements(t, (*Identifiable)(nil), data3)
	require.IsType(t, ResourceType{}, data3)

	key := "User"
	rT := rType.Pull(key)

	require.Equal(t, rT.GetIdentifier(), key)

	// PushFromData - Successful loading of a schema from bytes
	byt, _ := ioutil.ReadFile("../../internal/testdata/user_resource.json")
	data4, err6 := rType.PushFromData(byt)
	require.NoError(t, err6)
	require.Implements(t, (*Identifiable)(nil), data4)
	require.IsType(t, ResourceType{}, data4)

	key2 := "User 2"
	rT = rType.Pull(key2)

	require.Equal(t, rT.GetIdentifier(), key2)

	// Push - Successful loading of a schema
	var resType ResourceType
	resType.Name = "User 3"
	resType.Endpoint = "/User"
	resType.Description = "User Account "
	resType.Schema = "urn:ietf:params:scim:schemas:core:2.0:User"
	resType.CommonAttributes = *commons
	data5, err8 := rType.Push(resType)

	require.NoError(t, err8)
	require.Implements(t, (*Identifiable)(nil), data5)
	require.IsType(t, ResourceType{}, data5)

	key3 := "User 3"
	rT = rType.Pull(key3)

	require.Equal(t, rT.GetIdentifier(), key3)

	// (todo) > test lock

	rType.Clean() // teardown
}

func TestResourceTypeRepositoryList(t *testing.T) {
	repos := GetResourceTypeRepository()

	data, err := repos.PushFromFile("../../internal/testdata/user.json")
	require.NoError(t, err)
	require.Implements(t, (*Identifiable)(nil), data)
	require.IsType(t, ResourceType{}, data)

	data2, err2 := repos.PushFromFile("../../internal/testdata/user_resource.json")
	require.NoError(t, err2)
	require.Implements(t, (*Identifiable)(nil), data2)
	require.IsType(t, ResourceType{}, data2)

	list := repos.List()
	require.NotNil(t, list)
	require.Len(t, list, 2)
	require.IsType(t, []ResourceType{}, list)

	repos.Clean() // teardown
}

func TestSchemaRepositoryList(t *testing.T) {
	repos := GetSchemaRepository()

	data, err := repos.PushFromFile("../../internal/testdata/user_schema.json")
	require.NoError(t, err)
	require.Implements(t, (*Identifiable)(nil), data)
	require.IsType(t, Schema{}, data)

	data2, err2 := repos.PushFromFile("../../internal/testdata/enterprise_user_schema.json")
	require.NoError(t, err2)
	require.Implements(t, (*Identifiable)(nil), data2)
	require.IsType(t, Schema{}, data2)

	list := repos.List()
	require.NotNil(t, list)
	require.Len(t, list, 2)
	require.IsType(t, []Schema{}, list)

	repos.Clean() // teardown
}
