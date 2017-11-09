package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComplexHas(t *testing.T) {
	c := make(Complex)

	// Map does not own the key
	require.False(t, IsAssigned(c, "unassigned-key"))

	// Map owns the key but it contains nil
	var val *string
	c["pointing-key"] = val
	assert.False(t, IsAssigned(c, "pointing-key"))

	// Map owns the key but it does not contain a slice of a SCIM Data Type
	vals := make([]string, 2)
	c["multivalue-key"] = vals
	assert.False(t, IsAssigned(c, "multivalue-key"))

	// Map owns the key that contain an empty slice of a SCIM Data Type
	c["empty-multivalue-key"] = make([]String, 0)
	assert.False(t, IsAssigned(c, "empty-multivalue-key"))

	// Map own the key that does not contain a SCIM Data Type
	c["existent-key"] = "val"
	assert.False(t, IsAssigned(c, "existent-key"))

	c["existent-key"] = String("val")
	assert.True(t, IsAssigned(c, "existent-key"))

	c["empty-multivalue-key"] = make([]String, 2)
	assert.True(t, IsAssigned(c, "empty-multivalue-key"))
}

func TestSingleValueCheck(t *testing.T) {
}

func TestMultiValueCheck(t *testing.T) {
}
