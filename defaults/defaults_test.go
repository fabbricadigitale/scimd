package defaults

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func TestDefaults(t *testing.T) {
	require.IsType(t, ServiceProviderConfig, core.ServiceProviderConfig{})
	require.IsType(t, UserSchema, core.Schema{})
	require.IsType(t, GroupSchema, core.Schema{})
	require.IsType(t, UserResourceType, core.ResourceType{})
	require.IsType(t, GroupResourceType, core.ResourceType{})

	// (todo) > test the defaults
}
