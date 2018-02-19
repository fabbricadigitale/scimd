package defaults

import (
	"fmt"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
)

func TestDefaults(t *testing.T) {
	fmt.Printf("SVCPROV: %#v\n", ServiceProviderConfig)
	fmt.Printf("USCHEMA: %#v\n", UserSchema)
	fmt.Printf("GSCHEMA: %#v\n", GroupSchema)
	fmt.Printf("URESTYP: %#v\n", UserResourceType)
	fmt.Printf("GRESTYP: %#v\n", GroupResourceType)

	require.IsType(t, ServiceProviderConfig, &core.ServiceProviderConfig{})

	// (todo) > test the defaults
}
