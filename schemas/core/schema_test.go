package core

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/stretchr/testify/assert"
)

func TestAttributesWithName(t *testing.T) {

	a := &Attribute{
		Name: "Test",
	}

	list := Attributes{
		a,
	}

	assert.Equal(t, a, list.WithName("Test"))
	assert.Equal(t, a, list.WithName("tEst"))
	assert.Nil(t, list.WithName("nope"))
}

func TestAttributesSome(t *testing.T) {

	r := &Attribute{
		Name:     "R",
		Returned: schemas.ReturnedAlways,
	}

	n := &Attribute{
		Name:     "N",
		Returned: schemas.ReturnedNever,
	}

	list := Attributes{
		r,
		n,
	}

	assert.Equal(t, Attributes{r}, list.Some(func(attribute *Attribute) bool {
		return attribute.Returned == schemas.ReturnedAlways
	}))

	assert.Equal(t, list, list.Some(func(attribute *Attribute) bool {
		return true
	}))

	assert.Empty(t, list.Some(func(attribute *Attribute) bool {
		return false
	}))

}
