package core

import (
	"encoding/json"
	"testing"
	"time"

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

func TestDecodeValuedWhenSingle(t *testing.T) {

	attributes := []Attribute{
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "string", Name: "a"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "boolean", Name: "b"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "decimal", Name: "c"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "integer", Name: "d"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "dateTime", Name: "e"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "binary", Name: "f"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{Type: "reference", Name: "g"}},
		},
	}

	// String
	data := (json.RawMessage)(`{"a": "some value"}`)
	r, err := attributes[0].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c := r.(Complex)
	v, ok := c[attributes[0].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, String("some value"), v)

	// Boolean
	data = (json.RawMessage)(`{"b": true}`)
	r, err = attributes[1].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[1].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, Boolean(true), v)

	//Decimal
	data = (json.RawMessage)(`{"c": 3.14}`)
	r, err = attributes[2].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[2].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, Decimal(3.14), v)

	//Decimal
	data = (json.RawMessage)(`{"d": 123}`)
	r, err = attributes[3].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[3].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, Integer(123), v)

	//DateTime
	t2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	data = (json.RawMessage)(`{"e": "` + t2.Format(time.RFC3339Nano) + `"}`)

	r, err = attributes[4].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[4].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))
	assert.Equal(t, DateTime(t2), v)

	//Binary
	b := []byte{'g'}
	data = (json.RawMessage)(`{"f": "` + string(b) + `"}`)
	r, err = attributes[5].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[5].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, Binary(b), v)

	//Reference
	data = (json.RawMessage)(`{"g": "https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646"}`)
	r, err = attributes[6].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[6].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsSingleValue(v))

	assert.Equal(t, Reference("https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646"), v)
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

func TestDecodeValuedWhenMulti(t *testing.T) {

	attributes := []Attribute{
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "string", Name: "a"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "boolean", Name: "b"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "decimal", Name: "c"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "integer", Name: "d"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "dateTime", Name: "e"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "binary", Name: "f"}},
		},
		Attribute{
			Type:          "complex",
			SubAttributes: Attributes{&Attribute{MultiValued: true, Type: "reference", Name: "g"}},
		},
	}

	// String
	data := (json.RawMessage)(`{"a": ["value a", "value b"]}`)
	r, err := attributes[0].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c := r.(Complex)
	v, ok := c[attributes[0].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, String("value a"))
	assert.Contains(t, v, String("value b"))

	// Boolean
	data = (json.RawMessage)(`{"b": [true, false]}`)
	r, err = attributes[1].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[1].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, Boolean(true))
	assert.Contains(t, v, Boolean(false))

	// Decimal
	data = (json.RawMessage)(`{"c": [3.14, 2.718]}`)
	r, err = attributes[2].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[2].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, Decimal(3.14))
	assert.Contains(t, v, Decimal(2.718))

	// Integer
	data = (json.RawMessage)(`{"d": [123, 456]}`)
	r, err = attributes[3].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[3].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, Integer(123))
	assert.Contains(t, v, Integer(456))

	// DateTime
	t2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t3 := time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC)
	data = (json.RawMessage)(`{"e": ["` + t2.Format(time.RFC3339Nano) + `","` + t3.Format(time.RFC3339Nano) + `"]}`)
	r, err = attributes[4].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[4].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, DateTime(t2))
	assert.Contains(t, v, DateTime(t3))

	// Binary
	b := []byte{'g'}
	b2 := []byte{'g', 'o'}
	data = (json.RawMessage)(`{"e": ["` + string(b) + `","` + string(b2) + `"]}`)
	r, err = attributes[5].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[5].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, Binary(b))
	assert.Contains(t, v, Binary(b2))

	// Reference
	data = (json.RawMessage)(`{"g": ["https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646", "https://example.com/v2/Users/2819c223-7f76-453a-919d-413861906464" ]}`)
	r, err = attributes[6].Unmarshal(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	require.IsType(t, Complex{}, r)

	c = r.(Complex)
	v, ok = c[attributes[6].SubAttributes[0].Name]

	assert.True(t, ok)
	assert.True(t, IsMultiValue(v))

	assert.Contains(t, v, Reference("https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646"))
	assert.Contains(t, v, Reference("https://example.com/v2/Users/2819c223-7f76-453a-919d-413861906464"))
}
