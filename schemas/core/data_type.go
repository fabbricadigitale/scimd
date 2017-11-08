package core

import (
	"encoding/json"
	"time"
)

// SCIM Schema "type" as per https://tools.ietf.org/html/rfc7643#section-2.3
const (
	StringType    = "string"
	BooleanType   = "boolean"
	DecimalType   = "decimal"
	IntegerType   = "integer"
	DateTimeType  = "dateTime"
	BinaryType    = "binary"
	ReferenceType = "reference"
	ComplexType   = "complex"
)

// NewDataType function allocates for SCIM Data Types.
// The first argument is string containing the SCIM Schema "type" as per https://tools.ietf.org/html/rfc7643#section-2.3,
// and the value returnerd is a DataType interface that holds a pointer to a newly allocated zero value of that "type".
func NewDataType(t string) (DataType, error) {
	switch t {
	case StringType:
		return new(String), nil
	case BooleanType:
		return new(Boolean), nil
	case DecimalType:
		return new(Decimal), nil
	case IntegerType:
		return new(Integer), nil
	case DateTimeType:
		return new(DateTime), nil
	case BinaryType:
		return new(Binary), nil
	case ReferenceType:
		return new(Reference), nil
	case ComplexType:
		return &Complex{}, nil
	}
	return nil, &dataTypeError{"Invalid type"}
}

type dataTypeError struct {
	msg string
}

func (e *dataTypeError) Error() string {
	return e.msg
}

// String defines the equivalent SCIM Data Type and attaches the methods of DataType interface to string
type String string

// Boolean defines the equivalent SCIM Data Type and attaches the methods of DataType interface to bool
type Boolean bool

// Decimal defines the equivalent SCIM Data Type and attaches the methods of DataType interface to float64
type Decimal float64

// Integer defines the equivalent SCIM Data Type and attaches the methods of DataType interface to int64
type Integer int64

// DateTime defines the equivalent SCIM Data Type and attaches the methods of DataType interface to time.Time
type DateTime time.Time

// Binary defines the equivalent SCIM Data Type and attaches the methods of DataType interface to []byte
type Binary []byte

// Reference defines the equivalent SCIM Data Type and attaches the methods of DataType interface to []byte
type Reference string

// Complex defines the equivalent SCIM Data Type and attaches the methods of DataType interface to map[string]interface{}
type Complex map[string]interface{}

// DataType is the interface implemented by SCIM Data Types.
// DataTypes implement Indirect() returns the value that the current DataType points to.
type DataType interface {
	Indirect() interface{}
}

// Indirect returns the String value that DataType points to.
func (p *String) Indirect() interface{} {
	return *p
}

// Indirect returns the Boolean value that DataType points to.
func (p *Boolean) Indirect() interface{} {
	return *p
}

// Indirect returns the Decimal value that DataType points to.
func (p *Decimal) Indirect() interface{} {
	return *p
}

// Indirect returns the Integer value that DataType points to.
func (p *Integer) Indirect() interface{} {
	return *p
}

// Indirect returns the DateTime value that DataType points to.
func (p *DateTime) Indirect() interface{} {
	return *p
}

// Indirect returns the Binary value that DataType points to.
func (p *Binary) Indirect() interface{} {
	return *p
}

// Indirect returns the Reference value that DataType points to.
func (p *Reference) Indirect() interface{} {
	return *p
}

// Indirect returns the Complex value that DataType points to.
func (p *Complex) Indirect() interface{} {
	return *p
}

// UnmarshalJSON implements custom logic for Binary
func (p *Binary) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	*p = Binary(s)
	return err
}
