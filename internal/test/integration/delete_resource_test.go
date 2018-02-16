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

	notExistingID := "fake-id-doent-exist"

	err := delete.Resource(adapter, resTypeRepo.Pull("User"), notExistingID)
	require.Error(t, err, err)

	// Create a new resource to be deleted
	res := &resource.Resource{
		CommonAttributes: core.CommonAttributes{
			Schemas:    []string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"},
			ExternalID: "5666",
			Meta: core.Meta{
				ResourceType: "User",
				Location:     "something",
			},
		},
	}

	res.SetValues("urn:ietf:params:scim:schemas:core:2.0:User", &datatype.Complex{
		"userName": datatype.String("alelb"),
	})

	res.SetValues("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User", &datatype.Complex{
		"employeeNumber": "701984",
	})

	create.Resource(adapter, resTypeRepo.Pull("User"), res)

	id := res.ID

	// delete the newly created resource
	err = delete.Resource(adapter, resTypeRepo.Pull("User"), id)
	require.Nil(t, err)

	// checks the successfully deletion
	retRes, err := adapter.Get(resTypeRepo.Pull("User"), id, "", nil)
	require.Nil(t, retRes)

}
