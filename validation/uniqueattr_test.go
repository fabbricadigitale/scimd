package validation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Tests []*Test 

type Test struct {
	Name 	string 
	Surname string 
	Age 	int
}

type Testy struct {
	Tests Tests `validate:"uniqueattr=Name"`
}

type NonExistingFieldTest struct {
	Tests Tests `validate:"uniqueattr=Nonexisting"`
}

type NonStringFieldTest struct {
	Tests Tests `validate:"uniqueattr=Age"`
}

func TestUniqueAttr(t *testing.T) {

	// Valid test, fields values are unique
	x := Testy{}

	y := Tests{
		{
			Name: "Waldo",
			Surname: "Baldo",
		},
		{
			Name: "Baldo",
			Surname: "Waldo",
		},
	}

	x.Tests = y

	err := Validator.Struct(x)
	require.NoError(t, err)

	// Fail test, fields values are not unique
	yz := Tests{
		{
			Name: "Waldo",
			Surname: "Baldo",
		},
		{
			Name: "waldo",
			Surname: "Waldo",
		},
	}

	x.Tests = yz

	err = Validator.Struct(x)
	require.Error(t, err)

	// Fail test, validator works on string fields only
	xx := NonStringFieldTest{}

	yy := Tests{
		{
			Age: 27,
		},
	}

	xx.Tests = yy

	err = Validator.Struct(xx)
	require.Error(t, err)

	// Fail test, field is not existing
	z := NonExistingFieldTest{}
	err = Validator.Struct(z)
	require.Error(t, err)

	// Fail test, validator can't be used on types that are not Slices
	xy := "Waldo"
	require.PanicsWithValue(t, "Can't be used on string, only on a Slice", func() {
		Validator.Var(xy, "uniqueattr=Name")
	})
}