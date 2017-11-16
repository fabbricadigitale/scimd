package messages

const PatchOpURN = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

type PatchOp struct {
	Schemas    []string `json:"schemas"`
	Operations []struct {
		Op    string `json:"op"`
		Path  string `json:"path,omitempty"`
		Value []struct {
			Display string `json:"display,omitempty"`
			Ref     string `json:"$ref,omitempty"`
			Value   string `json:"value"`
			Type    string `json:"type,omitempty"`
		} `json:"value,omitempty"`
	} `json:"Operations"`
}
