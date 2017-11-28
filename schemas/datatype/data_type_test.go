package datatype

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

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

	c["empty-multivalue-key"] = make([]DataTyper, 2)
	assert.False(t, IsNull(c["empty-multivalue-key"]))
}

func TestIsNull(t *testing.T) {
	// nil
	assert.True(t, IsNull(nil))

	// multi-value type with length equals to 0
	assert.True(t, IsNull(make([]DataTyper, 0)))
	assert.False(t, IsNull(make([]DataTyper, 1)))

	// values of types not included in single-value nor multi-value ones
	assert.True(t, IsNull("hello world!"))
	assert.False(t, IsNull(String("hello world!")))
}

func TestIsSingleValue(t *testing.T) {
	single := String("one")
	var emptySingle String
	multi := []String{"one", "two"}
	var emptyMulti []String

	// nil
	assert.False(t, IsSingleValue(nil), "nil")

	// multi Data Type
	assert.False(t, IsSingleValue(multi))

	// empty Data Type
	assert.True(t, IsSingleValue(emptySingle), "empty")
	assert.False(t, IsSingleValue(emptyMulti))

	// Data Type String
	assert.True(t, IsSingleValue(String(single)), "valued")
	assert.False(t, IsSingleValue("test"))

	// Data type Boolean
	assert.True(t, IsSingleValue(Boolean(false)))
	assert.False(t, IsSingleValue(true))

	//Data type Decimal
	assert.True(t, IsSingleValue(Decimal(3.14)))
	assert.False(t, IsSingleValue(3.14))

	// Data Type Integer
	assert.True(t, IsSingleValue(Integer(123)))
	assert.False(t, IsSingleValue(123))

	// Data Type DateTime
	assert.True(t, IsSingleValue(DateTime(time.Now())))
	assert.False(t, IsSingleValue(time.Now()))

	// Data Type Binary
	assert.True(t, IsSingleValue(Binary([]byte{1, 2})))
	assert.False(t, IsSingleValue([]byte{1, 2}))

	// Data Type Reference
	assert.True(t, IsSingleValue(Reference("https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646")))
	assert.False(t, IsSingleValue("https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646"))

	// Data Type Complex
	var c = map[string]interface{}{}
	assert.True(t, IsSingleValue(Complex(c)))
	assert.False(t, IsSingleValue(c))
}

func TestIsMultiValue(t *testing.T) {
	single := String("one")

	var emptySingle String

	multi := []DataTyper{String("one"), String("two")}

	var emptyMulti []DataTyper

	// nil
	assert.False(t, IsMultiValue(nil), "nil")

	// Data type empty
	assert.True(t, IsMultiValue(emptyMulti))
	assert.False(t, IsMultiValue(emptySingle), "empty")

	// Data type String
	assert.True(t, IsMultiValue(multi))
	assert.False(t, IsMultiValue(single), "valued")

	// Data type Boolean
	assert.True(t, IsMultiValue([]DataTyper{Boolean(true), Boolean(false)}))
	assert.False(t, IsMultiValue(Boolean(true)))

	// Data type Decimal
	assert.True(t, IsMultiValue([]DataTyper{Decimal(1.23), Decimal(4.56)}))
	assert.False(t, IsMultiValue(Decimal(6.78)))

	// Data type Integer
	assert.True(t, IsMultiValue([]DataTyper{Integer(1), Integer(2)}))
	assert.False(t, IsMultiValue(Integer(3)))

	// Data type DateTime
	dt1 := DateTime(time.Now())
	dt2 := DateTime(time.Now())

	assert.True(t, IsMultiValue([]DataTyper{DateTime(dt1), DateTime(dt2)}))
	assert.False(t, IsMultiValue(DateTime(dt1)))

	// Data type Binary
	assert.True(t, IsMultiValue([]DataTyper{Binary([]byte{1}), Binary([]byte{2})}))
	assert.False(t, IsMultiValue(Binary([]byte{1})))

	// Data type Reference
	assert.True(t, IsMultiValue([]DataTyper{Reference("urn"), Reference("URL")}))
	assert.False(t, IsMultiValue(Reference("urn")))

	// Data type Complex
	var c1 = map[string]interface{}{}
	var c2 = map[string]interface{}{}
	assert.True(t, IsMultiValue([]DataTyper{Complex(c1), Complex(c2)}))
	assert.False(t, IsMultiValue(Complex(c1)))
}

func TestDateTimeUnmarshal(t *testing.T) {

	tt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	data := (json.RawMessage)(`"` + tt.Format(time.RFC3339Nano) + `"`)

	d := &DateTime{}
	err := d.UnmarshalJSON(data)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	assert.Equal(t, tt.Format(time.RFC3339Nano), (time.Time)(*d).Format(time.RFC3339Nano))

	byt, err := json.Marshal((time.Time)(*d))

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	assert.Equal(t, []byte(data), []byte(byt))

}

func TestDateTimeMarshal(t *testing.T) {

	tt := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	data := (json.RawMessage)(`"` + tt.Format(time.RFC3339Nano) + `"`)

	d := &DateTime{}
	(*d) = DateTime(tt)

	byt, err := json.Marshal(d)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.Equal(t, data, json.RawMessage(byt))
}

func TestBinaryUnmarshalMarshal(t *testing.T) {
	b := []byte(`"R28gR28gR29sYW5nIQ=="`)

	d := &Binary{}
	err := d.UnmarshalJSON(b)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.Equal(t, "Go Go Golang!", string(*d))

	byt, err := json.Marshal(*d)

	require.Equal(t, b, byt)
}
