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
	var index, count uint
	index = 1
	count = 0

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
		{count, s.Count},
	}

	for _, row := range equalities {
		require.Equal(t, row.value, row.field)
	}

}

func TestSearchValidation(t *testing.T) {
	s := Search{}

	var err error

	defer func() {
		r := recover()
		require.NotNil(t, r)
		require.Equal(t, "Field Type not found in the Struct", r)
	}()

	// Wrong Search struct tags
	// Attributes attrname
	s.Attributes.Attributes = []string{"1userName"}
	err = validation.Validator.Var(s, "attrname")
	require.Error(t, err)

	// ExcludedAttributes attrname
	s.Attributes.ExcludedAttributes = []string{"2age"}
	err = validation.Validator.Var(s, "attrname")
	require.Error(t, err)

	// Pagination StartIndex
	s.Pagination.StartIndex = 0
	err = validation.Validator.Var(s, "gt")
	require.Error(t, err)

	// Sorting SortOrder
	s.Sorting.SortOrder = "sam"
	err = validation.Validator.Var(s, "eq=ascending|eq=descending")
	require.Error(t, err)

	// Right Search struct tags
	// Attributes attrname
	s.Attributes.Attributes = []string{"userName"}
	err = validation.Validator.Var(s, "attrname")
	require.NoError(t, err)

	// ExcludedAttributes attrname
	s.Attributes.ExcludedAttributes = []string{"age"}
	err = validation.Validator.Var(s, "attrname")
	require.NoError(t, err)

	// Pagination StartIndex
	s.Pagination.StartIndex = 1
	err = validation.Validator.Var(s, "gt")
	require.NoError(t, err)

	// Sorting SortOrder
	s.Sorting.SortOrder = "ascending"
	err = validation.Validator.Var(s, "eq=ascending|eq=descending")
	require.NoError(t, err)
}
