package messages

// PatchOpURI is the URI identifies the patch query requests
const PatchOpURI = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

// Operation struct
type Operation struct {
	Op    string      `json:"op" binding:"required" validate:"required,eq=add|eq=remove|eq=replace"`
	Path  string      `json:"path,omitempty" validate:"isfield=Op:add|isfield=Op:replace|required,omitempty,attrpath"`
	Value interface{} `json:"value,omitempty"`
}

// PatchOp represents a patch query as per https://tools.ietf.org/html/rfc7644#section-3.5.2
type PatchOp struct {
	Schemas    []string     `json:"schemas" validate:"eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:PatchOp"`
	Operations []*Operation `json:"Operations" binding:"required" validate:"required"`
}
