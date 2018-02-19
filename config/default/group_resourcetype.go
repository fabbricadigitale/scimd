package config

import (
	"fmt"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

func init() {
	resType := core.ResourceType{}
	
	schema := "urn:ietf:params:scim:schemas:core:2.0:ResourceType"
	commonResType := "ResourceType"
	id := "Group" 
	
	commons := core.NewCommon(schema, commonResType, id)

	resType.CommonAttributes = *commons
	resType.Name = "Group"
	resType.Endpoint = "/Group"
	resType.Description = "Group"
	resType.Schema = "urn:ietf:params:scim:schemas:core:2.0:Group"

	resType.Meta.Location = fmt.Sprintf("/v2/ResourceTypes/%s", id)
}