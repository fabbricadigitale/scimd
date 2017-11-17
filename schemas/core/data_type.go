package core

import (
	"encoding/json"
	"time"
)

// DataType is the interface implemented by SCIM Data Types.
// DataTypes implement Value() that returns the value that DataType holds,
// Type() that returns the corrisponding SCIM Schema "type"
type DataType interface {
	Value() DataType
	Type() string
}

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

// String defines the equivalent SCIM Data Type and attaches the methods of DataType interface to string
type String string

// Value returns the DataType's value
func (p String) Value() DataType { return p }

// Type returns DataType's "type"
func (p String) Type() string { return StringType }

// Boolean defines the equivalent SCIM Data Type and attaches the methods of DataType interface to bool
type Boolean bool

// Value returns the DataType's value
func (p Boolean) Value() DataType { return p }

// Type returns DataType's "type"
func (p Boolean) Type() string { return BooleanType }

// Decimal defines the equivalent SCIM Data Type and attaches the methods of DataType interface to float64
type Decimal float64

// Value returns the DataType's value
func (p Decimal) Value() DataType { return p }

// Type returns DataType's "type"
func (p Decimal) Type() string { return DecimalType }

// Integer defines the equivalent SCIM Data Type and attaches the methods of DataType interface to int64
type Integer int64

// Value returns the DataType's value
func (p Integer) Value() DataType { return p }

// Type returns DataType's "type"
func (p Integer) Type() string { return IntegerType }

// DateTime defines the equivalent SCIM Data Type and attaches the methods of DataType interface to time.Time
type DateTime time.Time

// Value returns the DataType's value
func (p DateTime) Value() DataType { return p }

// Type returns DataType's "type"
func (p DateTime) Type() string { return DateTimeType }

// UnmarshalJSON implements custom logic for DateTime
func (p *DateTime) UnmarshalJSON(b []byte) error {
	return (*time.Time)(p).UnmarshalJSON(b)
}

// Binary defines the equivalent SCIM Data Type and attaches the methods of DataType interface to []byte
type Binary []byte

// Value returns the DataType's value
func (p Binary) Value() DataType { return p }

// Type returns DataType's "type"
func (p Binary) Type() string { return BinaryType }

// UnmarshalJSON implements custom logic for Binary
func (p *Binary) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, (*[]byte)(p))
}

// MarshalJSON implements custom logic for Binary
func (p *Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal((*[]byte)(p))
}

// Reference defines the equivalent SCIM Data Type and attaches the methods of DataType interface to []byte
type Reference string

// Value returns the DataType's value
func (p Reference) Value() DataType { return p }

// Type returns DataType's "type"
func (p Reference) Type() string { return ReferenceType }

// Complex defines the equivalent SCIM Data Type and attaches the methods of DataType interface to map[string]interface{}
type Complex map[string]interface{}

// Value returns the DataType's value
func (p Complex) Value() DataType { return p }

// Type returns DataType's "type"
func (p Complex) Type() string { return ComplexType }

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

// IsSingleValue hecks if v holds a Data Type value
func IsSingleValue(v interface{}) bool {
	_, ok := v.(DataType)
	return ok
}

// IsMultiValue checks if v holds a slices of Data Type values
func IsMultiValue(v interface{}) bool {
	_, ok := v.([]DataType)
	return ok
}

// If v is a multi-value return its length, otherwise zero
func multiValueLen(v interface{}) int {
	if dt, ok := v.([]DataType); ok {
		return len(dt)
	}
	return 0
}

// IsNull checks if v holds a Null value
// Convention for Null values is defined as following:
//  - nil is the "null" value
//  - zero-length multi-value is the empty array
//  - values that are not IsSingleValue nor IsMultiValue (ie. those values are not Data Types)
//
// As per https://tools.ietf.org/html/rfc7643#section-2.5
func IsNull(v interface{}) bool {
	return v == nil || (multiValueLen(v) == 0 && !IsSingleValue(v))
}
