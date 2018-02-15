package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/fabbricadigitale/scimd/storage"
	"github.com/fabbricadigitale/scimd/storage/mongo"
	"github.com/ory/dockertest"
)

var pul *dockertest.Pool
var res *dockertest.Resource
var err error
var ada storage.Storer

func setup() {
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
	}
	res, err = pul.RunWithOptions(&opt)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// The application in the container might not be ready to accept connections yet
	if err = pul.Retry(func() error {
		endpoint := fmt.Sprintf("localhost:%s", res.GetPort("27017/tcp"))
		ada, err = mongo.New(endpoint, "scimd", "resources")
		if err != nil {
			return err
		}

		return ada.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func teardown() {
	// Kill and remove the container
	if err := pul.Purge(res); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestSimpleGet(t *testing.T) {
	setup()
	defer teardown()

	eng := GetEngine()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v2/Users", nil)
	eng.ServeHTTP(rec, req)

	fmt.Println(rec)
}
