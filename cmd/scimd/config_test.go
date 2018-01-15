package main

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {

	spc := config()
	require.IsType(t, spc, &core.ServiceProviderConfig{})

	schemaList := filesFromDir("./default/schemas")

	schemaRepo := core.GetSchemaRepository()
	require.Equal(t, len(schemaList), len(schemaRepo.List()))

	resTypeList := filesFromDir("./default/resources")

	resTypeRepo := core.GetResourceTypeRepository()
	require.Equal(t, len(resTypeList), len(resTypeRepo.List()))

	// (todo)
	// phase 2 - requires parametrization (path of service provider config JSON) of config function
	// - test panics when wrong path
	// - test panics with unmarshalling errors
	// - test panics with validation errors
}
