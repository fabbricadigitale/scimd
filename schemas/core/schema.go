package core

import (
	"encoding/json"
	"regexp"

	"github.com/mcuadros/go-defaults"
)

// URIExpr is the source text used to compile a regular expression matching a SCIM "schema" URI
const URIExpr = `urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:[A-Za-z0-9()+,\-.:=@;$_!*'%\/?#]+` // TODO double check, should match SCIM "schema" URI only?

// AttrNameExpr is the source text used to compile a regular expression macching a SCIM attribute name
const AttrNameExpr = `[A-Za-z][\-$_0-9A-Za-z]*`

// AttrNameRegexp is the compiled Regexp built from AttrNameExpr
var AttrNameRegexp = regexp.MustCompile("^" + AttrNameExpr + "$")

// A SchemaError is a description of a SCIM error.
type SchemaError struct {
	msg string // description of error
}

func (e *SchemaError) Error() string { return e.msg }

// Attributes is a slice of Attribute that holds definitions of attributes included within a Schema Definition.
type Attributes []*Attribute

// Attribute describes a single attribute included within a Schema Definition.
// It includes the characteristics of a SCIM Attribute as per https://tools.ietf.org/html/rfc7643#section-2.2.
type Attribute struct {
	Name            string     `json:"name,omitempty" validate:"attrname"`
	Type            string     `json:"type,omitempty" default:"string"`
	SubAttributes   Attributes `json:"subAttributes,omitempty"`
	MultiValued     bool       `json:"multiValued"`
	Description     string     `json:"description,omitempty"`
	Required        bool       `json:"required"`
	CanonicalValues []string   `json:"canonicalValues,omitempty"`
	CaseExact       bool       `json:"caseExact,omitempty"`
	Mutability      string     `json:"mutability,omitempty" default:"readWrite"`
	Returned        string     `json:"returned,omitempty" default:"default"`
	Uniqueness      string     `json:"uniqueness,omitempty" default:"none"`
	ReferenceTypes  []string   `json:"referenceTypes,omitempty"`
}

// NewAttribute returns a new Attribute filled with defaults
func NewAttribute() *Attribute {
	a := new(Attribute)
	defaults.SetDefaults(a)
	return a
}

// UnmarshalJSON unmarshals an Attribute taking into account defaults
func (a *Attribute) UnmarshalJSON(data []byte) error {
	defaults.SetDefaults(a)

	type aliasType Attribute
	alias := aliasType(*a)
	err := json.Unmarshal(data, &alias)

	*a = Attribute(alias)
	return err
}

// Schema is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:Schema"
type Schema struct {
	Common
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Attributes  Attributes `json:"attributes,omitempty"`
}

var _ Resource = (*Schema)(nil)

// GetIdentifier ...
func (s Schema) GetIdentifier() string {
	return s.ID
}
