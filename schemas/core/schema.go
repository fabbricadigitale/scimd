package core

import (
	"encoding/json"
	"strings"

	"github.com/mcuadros/go-defaults"
)

// A SchemaError is a description of a SCIM error.
type SchemaError struct {
	msg string // description of error
}

func (e *SchemaError) Error() string { return e.msg }

// Attributes is a slice of Attribute that holds definitions of attributes included within a Schema Definition.
type Attributes []*Attribute

// ByName returns the *Attribute with the given name, performing a insensitive match. It returns nil if no attribute was found.
func (attributes Attributes) ByName(name string) *Attribute {
	name = strings.ToLower(name)
	for _, a := range attributes {
		if name == strings.ToLower(a.Name) {
			return a
		}
	}
	return nil
}

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

// SchemaURI is the Schema Definitions Schema used by Schema
const SchemaURI = "urn:ietf:params:scim:schemas:core:2.0:Schema"

// Schema is a structured resource for "urn:ietf:params:scim:schemas:core:2.0:Schema"
type Schema struct {
	CommonAttributes
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Attributes  Attributes `json:"attributes,omitempty"`
}

// NewSchema returns a new Schema filled with defaults
func NewSchema() *Schema {
	return &Schema{
		CommonAttributes: *NewCommon(SchemaURI, "Schema", SchemaURI),
		Name:             "Schema",
	}
}

var _ ResourceTyper = (*Schema)(nil)

// GetIdentifier ...
func (s Schema) GetIdentifier() string {
	return s.ID
}
