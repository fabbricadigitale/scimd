package query

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository

func TestMain(m *testing.M) {
	// Test setup
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.PushFromFile("../../internal/testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.PushFromFile("../../internal/testdata/user_schema.json")
	schemaRepo.PushFromFile("../../internal/testdata/enterprise_user_schema.json")

	// Test run
	m.Run()

	// No teardown, came back to home please
}
