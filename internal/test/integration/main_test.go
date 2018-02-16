package integration

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/ory/dockertest"
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
	resTypeRepo.PushFromFile("../../testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.PushFromFile("../../testdata/user_schema.json")
	schemaRepo.PushFromFile("../../testdata/enterprise_user_schema.json")

	// Run our tests
	code := m.Run()

	// Kill and remove the container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
