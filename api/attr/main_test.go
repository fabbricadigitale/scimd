package attr

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository
var userRes = resource.Resource{}

func TestMain(m *testing.M) {
	// Test setup
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.PushFromFile("../../internal/testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.PushFromFile("../../internal/testdata/user_schema.json")
	schemaRepo.PushFromFile("../../internal/testdata/enterprise_user_schema.json")

	resTypeRepo.PushFromFile("../../internal/testdata/group.json")
	schemaRepo.PushFromFile("../../internal/testdata/group_schema.json")

	userData, err := ioutil.ReadFile("../../internal/testdata/enterprise_user_resource_1.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(userData, &userRes); err != nil {
		panic(err)
	}

	// Test run
	m.Run()

	// No teardown, came back to home please
}
