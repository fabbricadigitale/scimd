package core

// A SchemaError is a description of a SCIM error.
type SchemaError struct {
	msg string // description of error
}

func (e *SchemaError) Error() string { return e.msg }

// Attributes is
type Attributes []*Attribute

// Attribute ...
type Attribute struct {
	Name            string     `json:"name,omitempty"`
	Type            string     `json:"type,omitempty"`
	SubAttributes   Attributes `json:"subAttributes,omitempty"`
	MultiValued     bool       `json:"multiValued"`
	Description     string     `json:"description,omitempty"`
	Required        bool       `json:"required"`
	CanonicalValues []string   `json:"canonicalValues,omitempty"`
	CaseExact       bool       `json:"caseExact,omitempty"`
	Mutability      string     `json:"mutability,omitempty"`
	Returned        string     `json:"returned,omitempty"`
	Uniqueness      string     `json:"uniqueness,omitempty"`
	ReferenceTypes  []string   `json:"referenceTypes,omitempty"`
}

// Schema ...
type Schema struct {
	Common
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Attributes  Attributes `json:"attributes,omitempty"`
}

// GetIdentifier ...
func (s Schema) GetIdentifier() string {
	return s.Name
}
