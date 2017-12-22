package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"gopkg.in/ory-am/dockertest.v3"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/stretchr/testify/require"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository
var adapter storage.Storer

func TestMain(m *testing.M) {
	// It uses sensible defaults for windows (tcp/http) and linux/osx (socket)
	// Regarding darwin setting DOCKER_HOST environment variable is probably required
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	cwd, _ := os.Getwd()

	opts := dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "3.4",
		Env: []string{
			"MONGO_INITDB_DATABASE=scimd",
		},
		Mounts: []string{
			path.Clean(fmt.Sprintf("%s/../../testdata/initdb.d:/docker-entrypoint-initdb.d", cwd)),
		},
	}
	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// The application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("27017/tcp"))
		var err error
		adapter, err = mongo.New(endpoint, "scimd", "resources")
		if err != nil {
			return err
		}

		return adapter.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Repositories are prerequisites
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.Add("../../testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.Add("../../testdata/user_schema.json")
	schemaRepo.Add("../../testdata/enterprise_user_schema.json")

	// Run our tests
	code := m.Run()

	// Kill and remove the container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMongoCreate(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("../../testdata/enterprise_user_resource_1.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	err = adapter.Create(res)
	require.NoError(t, err)
}

func TestMongoGet(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	id := "2819c223-7f76-453a-919d-ab1234567891"
	resource, err := adapter.Get(resTypeRepo.Get("User"), id, "", nil)
	require.NoError(t, err)

	require.NotNil(t, resource)
	require.Equal(t, id, resource.ID)

	// Excluding externalId
	id = "2819c223-7f76-453a-919d-ab1234567891"
	m := make(map[attr.Path]bool)
	m[attr.Path{
		Name: "userName",
	}] = true
	m[attr.Path{
		Name: "schemas",
	}] = true
	m[attr.Path{
		Name: "id",
	}] = true
	m[attr.Path{
		Name: "meta",
	}] = true
	resource, err = adapter.Get(resTypeRepo.Get("User"), id, "", m)
	require.NoError(t, err)

	require.NotNil(t, resource)
	require.Equal(t, "", resource.ExternalID)

	// Non-existing ID
	id = "2819c223-7f76-453a-919d-ab1234567898"
	resource, err = adapter.Get(resTypeRepo.Get("User"), id, "", nil)
	require.Nil(t, resource)
	require.EqualError(t, err, "not found")

	// Empty ID
	resource, err = adapter.Get(resTypeRepo.Get("User"), "", "", nil)
	require.Nil(t, resource)
	require.EqualError(t, err, "not found")

}

func TestMongoDelete(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	// Delete object with specified id
	id := "2819c223-7f76-453a-919d-413861904648"
	err := adapter.Delete(resTypeRepo.Get("User"), id, "")
	require.NoError(t, err)

	id = "2819c223-7f76-453a-919d-111111111111"
	err = adapter.Delete(resTypeRepo.Get("User"), id, "")
	require.Error(t, err)
}

// (todo) > Test Update adapter method

func TestMongoFind(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	res := []*core.ResourceType{
		&core.ResourceType{
			CommonAttributes: core.CommonAttributes{
				Schemas: []string{"urn:ietf:params:scim:schemas:core:2.0:ResourceType"},
				ID:      "User",
				Meta: core.Meta{
					Location:     "https://example.com/v2/ResourceTypes/User",
					ResourceType: "ResourceType",
				},
			},
			Name:        "User",
			Endpoint:    "/User",
			Description: "User Account",
			Schema:      "urn:ietf:params:scim:schemas:core:2.0:User",
			SchemaExtensions: []core.SchemaExtension{
				core.SchemaExtension{
					Schema:   "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
					Required: true,
				},
			},
		},
	}

	var f filter.Filter = filter.AttrExpr{
		Path: attr.Path{
			URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
			Name: "userName",
			Sub:  "",
		},
		Op:    filter.OpEqual,
		Value: "bjensen@example.com",
	}

	q, err := adapter.Find(res, f)
	require.NotNil(t, q)
	require.NoError(t, err)

	// Invalid schema urn

	res[0].Schema = "invalid-urn"
	q, err = adapter.Find(res, f)
	require.Nil(t, q)
	require.Error(t, err)
}
