package integration

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/api/create"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

func TestCreate(t *testing.T) {

	res := &resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
			ExternalID: "5666",
			Meta: core.Meta{
				ResourceType: "User",
			},
		},
	}
	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"userName": datatype.String("alelb"),
	})
	res.SetValues("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", &datatype.Complex{
		"employeeNumber": "701984",
	})

	retRes, err := create.Resource(adapter, resTypeRepo.Pull("User"), res)
	require.Nil(t, err)
	require.NotNil(t, retRes)

	// check urn:ietf:params:scim:schemas:core:2.0:User.userName
	r := retRes.(*resource.Resource)
	values := r.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	userName := (*values)["userName"]
	require.Equal(t, datatype.String("alelb"), userName)

	// check urn:ietf:params:scim:schemas:extension:enterprise:2.0:User.employeeNumber
	extensionValues := r.Values("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
	employeeNumber := (*extensionValues)["employeeNumber"]
	require.Equal(t, datatype.String("701984"), employeeNumber)

	require.Equal(t, res.ID, r.ID)
	require.Equal(t, res.Meta.Version, r.Meta.Version)
}
