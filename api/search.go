package api

import "fmt"

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

// InvalidPathError is thrown when path attribute is invalid or malformed
type InvalidPathError struct {
	Path   string
	Detail string
}

func (e *InvalidPathError) Error() string {
	if len(e.Path) > 0 {
		return fmt.Sprintf("Path [%s] is invalid: %s", e.Path, e.Detail)
	}
	return fmt.Sprintf("Path is invalid: %s", e.Detail)
}

// InvalidFilterError is thrown when the specified filter syntax is invalid or the specified attribute and filter comparison is not supported
type InvalidFilterError struct {
	Filter string
	Detail string
}

func (e *InvalidFilterError) Error() string {
	if len(e.Filter) > 0 {
		return fmt.Sprintf("Filter [%s] is invalid: %s", e.Filter, e.Detail)
	}
	return fmt.Sprintf("Filter is invalid: %s", e.Detail)
}
