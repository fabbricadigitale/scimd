package api

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
	validator "gopkg.in/go-playground/validator.v9"
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
	s.Sorting.SortOrder = "sam"
	s.Pagination.StartIndex = 0

	errors := validation.Validator.Struct(s)
	require.NotNil(t, errors)
	require.IsType(t, (validator.ValidationErrors)(nil), errors)
	require.Len(t, errors, 2)

	structs := []string{"Sorting", "Pagination"}
	fields := []string{"SortOrder", "StartIndex"}
	failtags := []string{"eq=ascending|eq=descending", "gt"}

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "Search." + structs[e]
		if len(fields[e]) > 0 {
			exp += "." + fields[e]
		} else {
			fields[e] = structs[e]
		}
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	s.Sorting.SortOrder = "ascending"

	errors = validation.Validator.Struct(s)
	require.Len(t, errors, 1)

	structs = structs[1:]
	fields = fields[1:]
	failtags = failtags[1:]

	for e, err := range errors.(validator.ValidationErrors) {
		exp := "Search." + structs[e]
		if len(fields[e]) > 0 {
			exp += "." + fields[e]
		} else {
			fields[e] = structs[e]
		}
		require.Equal(t, exp, err.Namespace())
		require.Equal(t, fields[e], err.Field())
		require.Equal(t, failtags[e], err.ActualTag())
	}

	s.Pagination.StartIndex = 1

	errors = validation.Validator.Struct(s)
	require.NoError(t, errors)

}
