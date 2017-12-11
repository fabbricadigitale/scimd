package api

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
)

func TestSearchResource(t *testing.T) {
	dat, err := ioutil.ReadFile("testdata/search.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	s := Search{}
	json.Unmarshal(dat, &s)

	var filter = new(Filter)
	*filter = "userName eq john"
	var index uint
	index = 1

	equalities := []struct {
		value interface{}
		field interface{}
	}{
		{[]string{"userName", "email"}, s.Attributes.Attributes},
		{[]string{"age", "phoneNumber"}, s.ExcludedAttributes},
		{*filter, s.Filter},
		{"userName", s.SortBy},
		{"ascending", s.SortOrder},
		{index, s.StartIndex},
		{0, s.Count},
	}

	for _, row := range equalities {
		require.Equal(t, row.value, row.field)
	}

}

func TestSearchValidation(t *testing.T) {
	testWrong := Search{}
	testWrong.Attributes.Attributes = []string{"urn:ietf:params:scim:schemas:core:2.0"}
	testWrong.Attributes.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0"}
	testWrong.Pagination.StartIndex = 0
	testWrong.Sorting.SortOrder = "sam"

	testRight := Search{}
	testRight.Attributes.Attributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:userName"}
	testRight.Attributes.ExcludedAttributes = []string{"urn:ietf:params:scim:schemas:core:2.0:User:userName"}
	testRight.Pagination.StartIndex = 1
	testRight.Sorting.SortOrder = "ascending"

	var err error

	// Wrong Search struct tags
	// Attributes attrpath
	err = validation.Validator.Var(testWrong.Attributes, "attrpath")
	require.Error(t, err)

	// ExcludedAttributes attrpath
	err = validation.Validator.Var(testWrong.Attributes, "attrpath")
	require.Error(t, err)

	// Pagination StartIndex
	err = validation.Validator.Var(testWrong.Pagination, "gt")
	require.Error(t, err)

	// Sorting SortOrder
	err = validation.Validator.Var(testWrong.Sorting, "eq=ascending|eq=descending")
	require.Error(t, err)

	// Right Search struct tags
	// Attributes attrpath
	err = validation.Validator.Var(testRight.Attributes, "attrpath")
	require.NoError(t, err)

	// ExcludedAttributes attrpath
	err = validation.Validator.Var(testRight.Attributes, "attrpath")
	require.NoError(t, err)

	// Pagination StartIndex
	err = validation.Validator.Var(testRight.Pagination, "gt")
	require.NoError(t, err)

	// Sorting SortOrder
	err = validation.Validator.Var(testRight.Sorting, "eq=ascending|eq=descending")
	require.NoError(t, err)
}
