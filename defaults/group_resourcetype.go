package defaults

import (
	"github.com/fabbricadigitale/scimd/validation"
	"fmt"

	"github.com/fabbricadigitale/scimd/schemas/core"
)

// GroupResourceType is the default resource type for groups
var GroupResourceType core.ResourceType

func init() {
	schema := "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
	commonResType := "ResourceType"
	id := "Group"

	commons := core.NewCommon(schema, commonResType, id)

	GroupResourceType.CommonAttributes = *commons
	GroupResourceType.Name = "Group"
	GroupResourceType.Endpoint = "/Group"
	GroupResourceType.Description = "Group"
	GroupResourceType.Schema = "urn:ietf:params:scim:schemas:core:2.0:Group"

	GroupResourceType.Meta.Location = fmt.Sprintf("/v2/ResourceTypes/%s", id)

	if errors := validation.Validator.Struct(GroupResourceType); errors != nil {
		panic("group resourcetype default configuration incorrect")
	}

	// (todo) > mold

}
