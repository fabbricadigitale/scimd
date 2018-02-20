package defaults

import (
	"github.com/fabbricadigitale/scimd/validation"
	"fmt"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

// UserResourceType is the default resource type for users
var UserResourceType core.ResourceType

func init() {
	schema := "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
	commonResType := "ResourceType"
	id := "User"

	commons := core.NewCommon(schema, commonResType, id)

	UserResourceType.CommonAttributes = *commons
	UserResourceType.Name = "User"
	UserResourceType.Endpoint = "/User"
	UserResourceType.Description = "User Account "
	UserResourceType.Schema = "urn:ietf:params:scim:schemas:core:2.0:User"

	UserResourceType.Meta.Location = fmt.Sprintf("/v2/ResourceTypes/%s", id)

	// No user extension by default, for now
	// resType.SchemaExtensions = []core.SchemaExtension{
	// 	core.SchemaExtension{
	// 		Schema:   "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
	// 		Required: true,
	// 	},
	// }

	if errors := validation.Validator.Struct(UserResourceType); errors != nil {
		panic("user resourcetype default configuration incorrect")
	}
	
	// (todo) > mold

}
