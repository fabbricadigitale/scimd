package mongo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	invalidKey = "dot.dollar$"
)

func TestDocumentKeyEscaping(t *testing.T) {
	d := make(document)
	d[invalidKey] = "foo"

	d.escapeKeys()

	assert.Nil(t, d[invalidKey])
	assert.Equal(t, "foo", d[keyEscape(invalidKey)])

	d.unescapeKeys()

	fmt.Printf("%+v\n\n", d)

	assert.Equal(t, "foo", d[invalidKey])

}
