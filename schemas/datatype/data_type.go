package datatype

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// DataTyper is the interface implemented by SCIM Data Types.
// DataTypers implement Value() that returns the value that DataType holds,
// Type() that returns the corrisponding SCIM Schema "type"
type DataTyper interface {
	Value() DataTyper
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

// String defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to string
type String string

// Value returns the DataTyper's value
func (p String) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p String) Type() string { return StringType }

// Boolean defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to bool
type Boolean bool

// Value returns the DataTyper's value
func (p Boolean) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Boolean) Type() string { return BooleanType }

// Decimal defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to float64
type Decimal float64

// Value returns the DataTyper's value
func (p Decimal) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Decimal) Type() string { return DecimalType }

// Integer defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to int64
type Integer int64

// Value returns the DataTyper's value
func (p Integer) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Integer) Type() string { return IntegerType }

// DateTime defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to time.Time
type DateTime time.Time

// Value returns the DataTyper's value
func (p DateTime) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p DateTime) Type() string { return DateTimeType }

// UnmarshalJSON implements custom logic for DateTime
func (p *DateTime) UnmarshalJSON(b []byte) error {
	return (*time.Time)(p).UnmarshalJSON(b)
}

// MarshalJSON implements custom logic for DateTime
func (p *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal((time.Time)(*p))
}

// Binary defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to []byte
type Binary []byte

// Value returns the DataTyper's value
func (p Binary) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Binary) Type() string { return BinaryType }

// UnmarshalJSON implements custom logic for Binary
func (p *Binary) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, (*[]byte)(p))
}

// MarshalJSON implements custom logic for Binary
func (p *Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal((*[]byte)(p))
}

// Reference defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to []byte
type Reference string

// Value returns the DataTyper's value
func (p Reference) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Reference) Type() string { return ReferenceType }

// Complex defines the equivalent SCIM Data Type and attaches the methods of DataTyper interface to map[string]interface{}
type Complex map[string]interface{}

// Value returns the DataTyper's value
func (p Complex) Value() DataTyper { return p }

// Type returns DataTyper's "type"
func (p Complex) Type() string { return ComplexType }

// New function allocates for SCIM Data Types.
// The first argument is string containing the SCIM Schema "type" as per https://tools.ietf.org/html/rfc7643#section-2.3,
// and the value returnerd is a DataTyper interface that holds a pointer to a newly allocated zero value of that "type".
func New(t string) (DataTyper, error) {
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
	return nil, &InvalidDataTypeError{t}
}

var rTypes = map[string]reflect.Type{
	StringType:    reflect.TypeOf(*new(String)),
	BooleanType:   reflect.TypeOf(*new(Boolean)),
	DecimalType:   reflect.TypeOf(*new(Decimal)),
	IntegerType:   reflect.TypeOf(*new(Integer)),
	DateTimeType:  reflect.TypeOf(*new(DateTime)),
	BinaryType:    reflect.TypeOf(*new(Binary)),
	ReferenceType: reflect.TypeOf(*new(Reference)),
	ComplexType:   reflect.TypeOf(Complex{}),
}

// Cast function returns a SCIM Data Type of the given type t that holds the v's value.
func Cast(v interface{}, t string) (DataTyper, error) {

	if rt := rTypes[t]; rt != nil {
		rv := reflect.ValueOf(v)
		rv = reflect.Indirect(rv)
		if rv.Type().ConvertibleTo(rt) {
			return rv.Convert(rt).Interface().(DataTyper), nil
		}
		// Those values are not within Data Types must be considered to be Null
		return nil, nil
	}

	return nil, &InvalidDataTypeError{t}
}

// InvalidDataTypeError is a generic invalid type error
type InvalidDataTypeError struct {
	t string
}

func (e *InvalidDataTypeError) Error() string {
	return fmt.Sprintf("Invalid type %s", e.t)
}

// IsSingleValue hecks if v holds a Data Type value
func IsSingleValue(v interface{}) bool {
	_, ok := v.(DataTyper)
	return ok
}

// IsMultiValue checks if v holds a slices of Data Type values
func IsMultiValue(v interface{}) bool {
	_, ok := v.([]DataTyper)
	return ok
}

// If v is a multi-value return its length, otherwise zero
func multiValueLen(v interface{}) int {
	if dt, ok := v.([]DataTyper); ok {
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
