package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/api/create"
	"github.com/fabbricadigitale/scimd/api/delete"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

func TestDeleteResource(t *testing.T) {

	// Delete a not-existing resource
	notExistingID := "fake-id-doesnt-exist"

	err := delete.Resource(adapter, resTypeRepo.Pull("User"), notExistingID)
	require.Error(t, err)

	// Create a new resource to be deleted
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

	_, err = create.Resource(adapter, resTypeRepo.Pull("User"), res)
	require.NoError(t, err)

	id := res.ID

	// delete the newly created resource
	err = delete.Resource(adapter, resTypeRepo.Pull("User"), id)
	require.Nil(t, err)

	// check the successfully deletion
	_, err = adapter.Get(resTypeRepo.Pull("User"), id, "", nil)
	require.Error(t, err)

}
