package defaults

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

// UserResourceType is the default resource type for users
var UserResourceType core.ResourceType

func init() {
	resType := core.ResourceType{}

	schema := "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
	commonResType := "ResourceType"
	id := "User"

	commons := core.NewCommon(schema, commonResType, id)

	resType.CommonAttributes = *commons
	resType.Name = "User"
	resType.Endpoint = "/User"
	resType.Description = "User Account "
	resType.Schema = "urn:ietf:params:scim:schemas:core:2.0:User"

	resType.Meta.Location = fmt.Sprintf("/v2/ResourceTypes/%s", id)

	// No user extension by default, for now
	// resType.SchemaExtensions = []core.SchemaExtension{
	// 	core.SchemaExtension{
	// 		Schema:   "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
	// 		Required: true,
	// 	},
	// }

	// (todo) > validation
	// (todo) > mold

	UserResourceType = resType
}
