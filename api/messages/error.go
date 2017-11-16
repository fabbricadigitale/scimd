package messages

const ErrorURN = "urn:ietf:params:scim:api:messages:2.0:Error"

type Error struct {
	Schemas  []string `json:"schemas"`
	Status   string   `json:"status"`
	ScimType string   `json:"scimType,omitempty"`
	Detail   string   `json:"detail,omitempty"`
}
