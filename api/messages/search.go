package messages

const SEARCH_REQUEST_URN = "urn:ietf:params:scim:api:messages:2.0:SearchRequest"

type SearchRequest struct {
	Schemas            []string `json:"schemas"`
	Attributes         []string `json:"attributes,omitempty"`
	ExcludedAttributes []string `json:"excludedAttributes,omitempty"`
	Filter             string   `json:"filter,omitempty"`
	SortBy             string   `json:"sortBy,omitempty"`
	SortOrder          string   `json:"sortOrder,omitempty"`
	StartIndex         int      `json:"startIndex,omitempty"`
	Count              int      `json:"count,omitempty"`
}
