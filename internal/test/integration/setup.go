package integration

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	dc "github.com/fsouza/go-dockerclient"
	"github.com/ory/dockertest"
)

var pul *dockertest.Pool
var resDocker *dockertest.Resource
var err error
var adapter storage.Storer
var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository
var res resource.Resource

func setupDB() {
	// It uses sensible defaults for windows (tcp/http) and linux/osx (socket)
	// Regarding darwin setting DOCKER_HOST environment variable is probably required
	pul, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	cwd, _ := os.Getwd()
	opt := dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "3.4",
		Env: []string{
			"MONGO_INITDB_DATABASE=scimd",
		},
		Mounts: []string{
			path.Clean(fmt.Sprintf("%s/../../testdata/initdb.d:/docker-entrypoint-initdb.d", cwd)),
		},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"27017/tcp": {{HostIP: "", HostPort: strconv.Itoa(config.Values.Storage.Port)}},
		},
	}
	resDocker, err = pul.RunWithOptions(&opt)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// The application in the container might not be ready to accept connections yet
	if err = pul.Retry(func() error {
		endpoint := fmt.Sprintf("localhost:%s", resDocker.GetPort("27017/tcp"))
		adapter, err = mongo.New(endpoint, "scimd", "resources")
		if err != nil {
			return err
		}

		// Just for testing uniqueness
		keys, err := attr.GetUniqueAttributes()
		if err != nil {
			return err
		}
		adapter.SetIndexes(keys)

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
}

func teardownDB() {
	// Kill and remove the container
	if err := pul.Purge(resDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func setup() {
	res = resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
			ExternalID: "5666",
			Meta: core.Meta{
				ResourceType: "User",
			},
		},
	}
	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"userName": datatype.String(String(8)),
	})
	res.SetValues("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", &datatype.Complex{
		"employeeNumber": "701984",
	})
}

func teardown() {
	adapter.Delete(res.ResourceType(), res.ID, res.Meta.Version)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset is ...
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String is ...
func String(length int) string {
	return StringWithCharset(length, charset)
}
