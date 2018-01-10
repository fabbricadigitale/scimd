package main

import (
	"github.com/stretchr/testify/require"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"testing"
)

func TestConfig(t *testing.T) {

	spc := config()
	require.IsType(t, spc, &core.ServiceProviderConfig{})

	// (todo)
	// phase 1
	// - x = # of JSON files within default/schemas directory
	// - y = # of JSON files within default/resources directory
	// - call config()
	// - test it returns a ServiceProviderConfig instance
	// - test a schema repository containing x schemas has been created
	// - test a resource type repository containing y resource types has been created
	// phase 2 - requires parametrization (path of service provider config JSON) of config function
	// - test panics when wrong path
	// - test panics with unmarshalling errors
	// - test panics with validation errors
}
