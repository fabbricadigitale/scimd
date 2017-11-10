package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnassigned(t *testing.T) {
	c := make(Complex)

	// Map does not own the key
	assert.True(t, IsNull(c["unassigned-key"]))

	// Map owns the key but it contains nil
	var val *string
	c["pointing-key"] = val
	assert.True(t, IsNull(c["pointing-key"]))

	// Map owns the key but it does not contain a slice of a SCIM Data Type
	vals := make([]string, 2)
	c["multivalue-key"] = vals
	assert.True(t, IsNull(c["multivalue-key"]))

	// Map owns the key that contain an empty slice of a SCIM Data Type
	c["empty-multivalue-key"] = make([]String, 0)
	assert.True(t, IsNull(c["empty-multivalue-key"]))

	// Map own the key that does not contain a SCIM Data Type
	c["existent-key"] = "val"
	assert.True(t, IsNull(c["existent-key"]))

	c["existent-key"] = String("val")
	assert.False(t, IsNull(c["existent-key"]))

	c["empty-multivalue-key"] = make([]String, 2)
	assert.False(t, IsNull(c["empty-multivalue-key"]))
}

func TestIsNull(t *testing.T) {
	// nil
	assert.True(t, IsNull(nil))

	// multi-value type with length equals to 0
	assert.True(t, IsNull(make([]Integer, 0)))
	assert.False(t, IsNull(make([]Integer, 1)))

	// values of types not included in single-value nor multi-value ones
	assert.True(t, IsNull("hello world!"))
	assert.False(t, IsNull(String("hello world!")))
}

func TestSingleValueCheck(t *testing.T) {
}

func TestMultiValueCheck(t *testing.T) {
}
