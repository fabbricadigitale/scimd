package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

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

	// Escape - Unescape key from json

	data, err := ioutil.ReadFile("../testdata/keyescaping_document.json")
	require.NotNil(t, data)
	require.NoError(t, err)

	r := make(document)
	err = json.Unmarshal(data, &r)
	require.NoError(t, err)

	r.escapeKeys()
	assert.Nil(t, r[invalidKey])
	assert.Equal(t, "foo", r[keyEscape(invalidKey)])

	r.unescapeKeys()

	assert.Equal(t, "foo", r[invalidKey])
}
