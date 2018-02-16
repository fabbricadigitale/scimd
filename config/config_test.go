package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func filesFromDir(dir string) (res []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		path := filepath.Join(dir, f.Name())
		if filepath.Ext(path) == ".json" {
			res = append(res, path)
		}
	}

	return
}

func TestDefaults(t *testing.T) {
	fmt.Printf("=> %+v\n", Values)
}

func TestConfig(t *testing.T) {
	spc := Get()
	require.IsType(t, spc, &core.ServiceProviderConfig{})

	schemaList := filesFromDir(fmt.Sprintf("%s/schemas", defaultConfigPath))

	schemaRepo := core.GetSchemaRepository()
	require.Equal(t, len(schemaList), len(schemaRepo.List()))

	resTypeList := filesFromDir(fmt.Sprintf("%s/resources", defaultConfigPath))

	resTypeRepo := core.GetResourceTypeRepository()
	require.Equal(t, len(resTypeList), len(resTypeRepo.List()))

	// (todo)
	// phase 2 - requires parametrization (path of service provider config JSON) of config function
	// - test panics when wrong path
	// - test panics with unmarshalling errors
	// - test panics with validation errors
}
