package main

import (
	"github.com/stretchr/testify/require"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"testing"
)
// Note: testing phase 1 has static expected values.
const (
	//Number of JSON files within default/schemas directory
	x = 3
	//Number of JSON files within default/resources directory
	y = 2
)

func TestConfig(t *testing.T) {

	spc := config()
	require.IsType(t, spc, &core.ServiceProviderConfig{})

	schemaRepo := core.GetSchemaRepository()
	require.Equal(t, x, len(schemaRepo.List()))

	resTypeRepo := core.GetResourceTypeRepository()
	require.Equal(t, y, len(resTypeRepo.List()))

	// (todo)
	// phase 2 - requires parametrization (path of service provider config JSON) of config function
	// - test panics when wrong path
	// - test panics with unmarshalling errors
	// - test panics with validation errors
}

