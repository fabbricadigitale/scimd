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

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/ory/dockertest"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/storage"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/require"
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
	resTypeRepo.PushFromFile("../../internal/testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.PushFromFile("../../internal/testdata/user_schema.json")
	schemaRepo.PushFromFile("../../internal/testdata/enterprise_user_schema.json")

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

	var expectedNameGivenName datatype.String
	expectedCommons := core.CommonAttributes{
		Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		ID:         "2819c223-7f76-453a-919d-ab1234567891",
		ExternalID: "1234567890",
		Meta: core.Meta{
			Version:      "W/\"a330bc54f0671c9\"",
			Location:     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
			ResourceType: "User",
		},
	}

	expectedEmailsVal := []datatype.String{"tfork@example.com", "tiffy@fork.org"}
	expectedEmailsTyp := []datatype.String{"work", "home"}
	expectedNameGivenName = "Tiffany"

	// Document with a complete set of attributes
	doc := &document{
		"id":         "2819c223-7f76-453a-919d-ab1234567891",
		"schemas":    []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		"externalId": "1234567890",
		"meta": bson.M{
			"version":      "W/\"a330bc54f0671c9\"",
			"location":     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
			"resourceType": "User",
		},
		"urn:ietf:params:scim:schemas:core:2.0:User": bson.M{
			"userName": "tfork@example.com",
			"name": bson.M{
				"givenName":       "Tiffany",
				"middleName":      "Geraldine",
				"honorificPrefix": "Ms.",
				"honorificSuffix": "II",
				"formatted":       "Ms. Tiffany G Fork, II",
				"familyName":      "Fork",
			},
			"emails": []bson.M{
				{
					"value": "tfork@example.com",
					"type":  "work",
				},
				{
					"value": "tiffy@fork.org",
					"type":  "home",
				},
			},
		},
	}

	res := toResource(doc)
	require.NotNil(t, res)
	require.Equal(t, expectedCommons, res.CommonAttributes)

	values := res.Values("urn:ietf:params:scim:schemas:core:2.0:User")

	emails := (*values)["emails"]
	for i, val := range emails.([]datatype.DataTyper) {
		e := val.(*datatype.Complex)
		require.Equal(t, expectedEmailsVal[i], (*e)["value"])
	}
	for i, typ := range emails.([]datatype.DataTyper) {
		e := typ.(*datatype.Complex)
		require.Equal(t, expectedEmailsTyp[i], (*e)["type"])
	}

	name := (*values)["name"].(*datatype.Complex)
	require.Equal(t, expectedNameGivenName, (*name)["givenName"])

	// ID, Schemas and Meta.ResourceType
	// essential attributes that are REQUIRED for a document to be valid for the toResource() method

	expectedCommons2 := core.CommonAttributes{
		Schemas: []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		ID:      "2819c223-7f76-453a-919d-ab1234567891",
		Meta: core.Meta{
			ResourceType: "User",
		},
	}

	doc = &document{
		"id":      "2819c223-7f76-453a-919d-ab1234567891",
		"schemas": []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		"meta": bson.M{
			"resourceType": "User",
		},
	}

	res = toResource(doc)
	require.NotNil(t, res)
	require.Equal(t, expectedCommons2, res.CommonAttributes)

	require.Equal(t, "", res.ExternalID)

	// This set of panics tests will follow our assumption made on toResource()
	// that Schemas, ID and Meta.ResourceType will always be present
	//
	// EMPTY DOCUMENT
	doc = &document{}
	require.Panics(t, func() {
		toResource(doc)
	})

	// ONLY RESOURCE TYPE
	doc = &document{
		"meta": bson.M{
			"resourceType": "User",
		},
	}
	require.Panics(t, func() {
		toResource(doc)
	})

	// NO ID
	doc = &document{
		"schemas": []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		"meta": bson.M{
			"created":      time.Date(2018, time.January, 12, 12, 12, 12, 12, time.Local),
			"lastModified": time.Now(),
			"version":      "W/\"a330bc54f0671c9\"",
			"location":     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
			"resourceType": "User",
		},
	}
	require.Panics(t, func() {
		toResource(doc)
	})

	// NO SCHEMAS
	doc = &document{
		"id": "2819c223-7f76-453a-919d-ab1234567891",
		"meta": bson.M{
			"created":      time.Date(2018, time.January, 12, 12, 12, 12, 12, time.Local),
			"lastModified": time.Now(),
			"version":      "W/\"a330bc54f0671c9\"",
			"location":     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
			"resourceType": "User",
		},
	}
	require.Panics(t, func() {
		toResource(doc)
	})

	// NO RESOURCE TYPE
	doc = &document{
		"id":      "2819c223-7f76-453a-919d-ab1234567891",
		"schemas": []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
		"meta": bson.M{
			"created":      time.Date(2018, time.January, 12, 12, 12, 12, 12, time.Local),
			"lastModified": time.Now(),
			"version":      "W/\"a330bc54f0671c9\"",
			"location":     "https://example.com/v2/Users/2819c223-7f76-453a-919d-ab1234567890",
		},
	}
	require.Panics(t, func() {
		toResource(doc)
	})

	// NO META
	doc = &document{
		"id":      "2819c223-7f76-453a-919d-ab1234567891",
		"schemas": []interface{}{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
	}
	require.Panics(t, func() {
		toResource(doc)
	})
}
