package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2/bson"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository
var adapter storage.Storer
var a Adapter

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
			path.Clean(fmt.Sprintf("%s/../../internal/testdata/initdb.d:/docker-entrypoint-initdb.d", cwd)),
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
		adapter, err = New(endpoint, "scimd", "resources")
		if err != nil {
			return err
		}

		return adapter.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Repositories are prerequisites
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.Add("../../internal/testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.Add("../../internal/testdata/user_schema.json")
	schemaRepo.Add("../../internal/testdata/enterprise_user_schema.json")

	// Run our tests
	code := m.Run()

	// Kill and remove the container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMongotoDoc(t *testing.T) {
	require.NotNil(t, resTypeRepo)
	require.NotNil(t, schemaRepo)
	require.NotNil(t, adapter)

	dat, err := ioutil.ReadFile("../../internal/testdata/enterprise_user_resource_1.json")
	require.NoError(t, err)
	require.NotNil(t, dat)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)
	require.NoError(t, err)

	doc := a.toDoc(res)
	require.NotNil(t, doc)
}

func TestMongotoResource(t *testing.T) {
	doc := &document{
		"id":         "2819c223-7f76-453a-919d-ab1234567891",
		"schemas":    []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		"externalId": "1234567890",
		"meta": bson.M{
			"created":      time.Date(2018, time.January, 12, 12, 12, 12, 12, time.Local),
			"lastModified": time.Now(),
			"version":      "W/\"a330bc54f0671c9\"",
			"location":     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
			"resourceType": "User",
		},
	}

	res := toResource(doc)
	require.NotNil(t, res)
}
