package api

import "strings"

// Attributes represents ...
type Attributes struct {
	Attributes         []string `form:"attributes" json:"attributes,omitempty" validate:"dive,attrpath"`
	ExcludedAttributes []string `form:"excludedAttributes" json:"excludedAttributes,omitempty" validate:"dive,attrpath"`
}

// Explode splits the attributes content by comma.
//
// Notice that this assumes that URNs containing comma/s CANNOT be used.
func (a *Attributes) Explode() {
	attributesAcc := []string{}
	for _, x := range a.Attributes {
		attributesAcc = append(attributesAcc, strings.Split(x, ",")...)

	}

	excludedAttributesAcc := []string{}
	for _, y := range a.ExcludedAttributes {
		excludedAttributesAcc = append(excludedAttributesAcc, strings.Split(y, ",")...)
	}

	a.Attributes = attributesAcc
	a.ExcludedAttributes = excludedAttributesAcc
}

type Filter string

const (
	AscendingOrder  = "ascending"
	DescendingOrder = "descending"
)

type Sorting struct {
	SortBy    string `form:"sortBy" json:"sortBy,omitempty"`
	SortOrder string `form:"sortOrder" json:"sortOrder,omitempty" default:"ascending" validate:"eq=ascending|eq=descending"`
}

type Pagination struct {
	StartIndex int `form:"startIndex" json:"startIndex,omitempty" default:"1" validate:"gt=0"`
	Count      int `form:"count" json:"count,omitempty" mold:"min=0"`
}

type Search struct {
	Attributes
	Filter `form:"filter" json:"filter,omitempty"`
	Sorting
	Pagination
}
