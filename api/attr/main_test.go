package attr

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

var resTypeRepo core.ResourceTypeRepository
var schemaRepo core.SchemaRepository

func TestMain(m *testing.M) {
	// Test setup
	resTypeRepo = core.GetResourceTypeRepository()
	resTypeRepo.Add("../../internal/testdata/user.json")

	schemaRepo = core.GetSchemaRepository()
	schemaRepo.Add("../../internal/testdata/user_schema.json")

	// Test run
	m.Run()

	// No teardown, came back to home please
}
