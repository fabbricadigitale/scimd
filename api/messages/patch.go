package messages

// PatchOpURI is the URI identifies the patch query requests
const PatchOpURI = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

// Value struct
type Value struct {
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Value   string `json:"value" validate:"required"`
	Type    string `json:"type,omitempty"`
}

// Operation struct
type Operation struct {
	Op    string   `json:"op" validate:"required,eq=add|eq=remove|eq=replace"`
	Path  string   `json:"path,omitempty" validate:"isfield=Op:add|isfield=Op:replace|required,omitempty,attrpath"`
	Value []*Value `json:"value,omitempty" validate:"gt=0|isfield=Op:remove"`
}

// PatchOp represents a patch query as per https://tools.ietf.org/html/rfc7644#section-3.5.2
type PatchOp struct {
	Schemas    []string     `json:"schemas" validate:"eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:PatchOp"`
	Operations []*Operation `json:"Operations" validate:"required"`
}
