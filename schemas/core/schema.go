package core

// Attribute ...
type Attribute struct {
	Name            string       `json:"name,omitempty"`
	Type            string       `json:"type,omitempty"`
	SubAttributes   []*Attribute `json:"subAttributes,omitempty"`
	MultiValued     bool         `json:"multiValued"`
	Description     string       `json:"description,omitempty"`
	Required        bool         `json:"required"`
	CanonicalValues []string     `json:"canonicalValues,omitempty"`
	CaseExact       bool         `json:"caseExact,omitempty"`
	Mutability      string       `json:"mutability,omitempty"`
	Returned        string       `json:"returned,omitempty"`
	Uniqueness      string       `json:"uniqueness,omitempty"`
	ReferenceTypes  []string     `json:"referenceTypes,omitempty"`
}

// Schema ...
type Schema struct {
	Resource
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Attributes  []*Attribute `json:"attributes,omitempty"`
}
