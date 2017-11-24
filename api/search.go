package api

// TODO: validation

type Attributes struct {
	Attributes         []string `json:"attributes,omitempty"`
	ExcludedAttributes []string `json:"excludedAttributes,omitempty"`
}

type Filter string

const (
	AscendingOrder  = "ascending"
	DescendingOrder = "descending"
)

type Sorting struct {
	SortBy    string `json:"sortBy,omitempty"`
	SortOrder string `json:"sortOrder,omitempty"`
}

type Pagination struct {
	StartIndex uint `json:"startIndex,omitempty"`
	Count      uint `json:"count,omitempty"`
}

type Search struct {
	Attributes
	Filter `json:"filter,omitempty"`
	Sorting
	Pagination
}
