package core

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeNullWhenSingle(t *testing.T) {

	a := Attribute{
		Type:          "complex",
		SubAttributes: Attributes{&Attribute{Type: "string", Name: "a"}},
	}

	// Test state: value must be cleared
	data := (json.RawMessage)(`{"a": null}`)
	r, err := a.Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c := r.(Complex)
	v, ok := c[a.SubAttributes[0].Name]

	// equivalent of `ok && core.IsNull(v)`
	assert.Nil(t, v)
	assert.True(t, ok)

	// Test state: unassigned
	data = (json.RawMessage)(`{}`)
	r, err = a.Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[a.SubAttributes[0].Name]

	assert.Nil(t, v)
	assert.False(t, ok)

}

func TestDecodeNullWhenMulti(t *testing.T) {

	a := Attribute{
		Type:          "complex",
		SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "string", Name: "a"}},
	}

	// Test state: value must be cleared
	test := func(data json.RawMessage) {
		r, err := a.Unmarshal(data)
		if err != nil {
			t.Log(err)
			t.Fail()
		}
		require.IsType(t, Complex{}, r)

		c := r.(Complex)
		v, ok := c[a.SubAttributes[0].Name]
		require.IsType(t, []DataType{}, v)
		mv := v.([]DataType)

		// equivalent of `ok && core.IsNull(v)`
		assert.Len(t, mv, 0)
		assert.True(t, ok)
	}
	test((json.RawMessage)(`{"a": null}`))
	test((json.RawMessage)(`{"a": []}`))

	// Test state: unassigned
	data := (json.RawMessage)(`{}`)
	r, err := a.Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c := r.(Complex)
	v, ok := c[a.SubAttributes[0].Name]

	assert.Nil(t, v)
	assert.False(t, ok)
}
