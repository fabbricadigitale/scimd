package config

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	schemaRepo := core.GetSchemaRepository()
	require.Equal(t, 2, len(schemaRepo.List()))

	resTypeRepo := core.GetResourceTypeRepository()
	require.Equal(t, 2, len(resTypeRepo.List()))

	// (todo)
	// phase 2
	// - test schemas can be added
	// - test resource types can be added
	// - test default schemas can be overridden
	// - test default resource types can be overridden
}
