package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fabbricadigitale/scimd/schemas/core"
	v "github.com/fabbricadigitale/scimd/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

const defaultConfigPath = "default"
const dbURL = "0.0.0.0:32770"
const dbName = "scimd"
const dbCollection = "resources"

var (
	defaultResourcesPath string
	defaultSchemasPath   string
)

func init() {
	defaultResourcesPath = filepath.Join(defaultConfigPath, "resources")
	defaultSchemasPath = filepath.Join(defaultConfigPath, "schemas")
}

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

func config() (*core.ServiceProviderConfig, *core.ResourceTypeRepository, *core.SchemaRepository) {
	// Schemas
	schemas := core.GetSchemaRepository()
	for _, p := range filesFromDir(defaultSchemasPath) {
		schemas.Add(p)
	}

	// Resource types
	rtypes := core.GetResourceTypeRepository()
	for _, p := range filesFromDir(defaultResourcesPath) {
		rtypes.Add(p)
	}

	// Service provider configuration
	dat, err := ioutil.ReadFile(filepath.Join(defaultConfigPath, "service_provider_config.json"))
	if err != nil {
		panic("error reading service provider config file")
	}
	spc := core.NewServiceProviderConfig()
	err = json.Unmarshal(dat, &spc)
	if err != nil {
		panic("unmarshalling errors")
	}

	if err := v.Validator.Struct(spc); err != nil {
		errs := ""
		for _, e := range err.(validator.ValidationErrors) {
			errs += fmt.Sprintf("%s\n", e)
		}
		panic(fmt.Sprintf("validation errors\n%s", errs))
	}

	return spc, rtypes, schemas
}
