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
}



// (todo) > Test Get adapter method

// (todo) > Test Delete adapter method

// (todo) > Test Update adapter method

// (todo) > Test Find adapter method
