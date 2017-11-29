package api

type Attributes struct {
	Attributes         []string `json:"attributes,omitempty" validate:"dive,attrpath"`
	ExcludedAttributes []string `json:"excludedAttributes,omitempty" validate:"dive,attrpath"`
}

type Filter string

const (
	AscendingOrder  = "ascending"
	DescendingOrder = "descending"
)

type Sorting struct {
	SortBy    string `json:"sortBy,omitempty"`
	SortOrder string `json:"sortOrder,omitempty" default:"ascending" validate:"eq=ascending|eq=descending"`
}

type Pagination struct {
	StartIndex uint `json:"startIndex,omitempty" default:"1" validate:"gt=0"`
	Count      uint `json:"count,omitempty" validate:"gte=0"` // TODO: a negative value shall be interpreted as 0
}

type Search struct {
	Attributes
	Filter `json:"filter,omitempty"`
	Sorting
	Pagination
}
