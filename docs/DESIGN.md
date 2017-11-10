# About

This document is a draft that describes how `scimd` has been designed, 
it should include primary design goals, internal specifications, and implementation details. 
It may not always be completely up to date.

## Design

*TBD*

### Attributes

#### Data Types

Mapping from [SCIM Data Types](https://tools.ietf.org/html/rfc7643#section-2.3) to `go` type are defined as following:

Golang `type` 	| SCIM Data Type 	|  SCIM Schema "type" 	| JSON Type 	|
|----------------	|---------------------	|-------------	|-------------	|
`string`   | String    | "string"    | String per Section 7 of [RFC7159]       |
`bool`   | Boolean   | "boolean"   | Value per Section 3 of [RFC7159]        |
`float64`   | Decimal   | "decimal"   | Number per Section 6 of [RFC7159]       |
`int64`   | Integer   | "integer"   | Number per Section 6 of [RFC7159]       |
`time.Time`   | DateTime  | "dateTime"  | String per Section 7 of [RFC7159]       |
`[]byte`   | Binary    | "binary"    | Binary value base64 encoded per Section 4 of [RFC4648], or with URL and filename safe alphabet URL per Section 5 of [RFC4648] that is passed as a JSON string per Section 7 of [RFC7159]       |
`string`   | Reference | "reference" | String per Section 7 of [RFC7159]       |
`map[string]interface{}`   | Complex   | "complex"   | Object per Section 4 of [RFC7159]       |

The [`core`](../schemas/core/data_type.go) package defines the above `type`s with the same name defined by SCIM Data Type, so:

```go
package core

// ...

type String string
type Boolean bool
type Decimal float64
type Integer int64
type DateTime time.Time
type Binary []byte
type Reference string
type Complex map[string]interface{}
```

Furthermore, all `type`s implement:
```go 
type DataType interface {
    // ...
}
```

#### Singular and Multi-Valued

By convention, the `type` of a **single-value** (for *singular attribute*) must be one of:
```go
String, Boolean, Decimal, Integer, DateTime, Binary, Reference, Complex
```
Instead, the `type` of a **multi-value** (for *multi-valued attribute*) must be on of:
```go
[]String, []Boolean, []Decimal, []Integer, []DateTime, []Binary, []Reference, []Complex
```

> A value of another `type` is considered a **Null** value.

#### Unassigned and Null Values

The internal convention for [Section 2.5 of [RFC7643]](https://tools.ietf.org/html/rfc7643#section-2.5) implements the following rules:

**Null** values are:
* `nil` *(equivalent to null, ie. JSON `null`)*
* **multi-value** `type` with length equals to 0 *(equivalent to empty array, ie. JSON `[]`)*
* values of `type`s not included in **single-value** nor **multi-value** ones

**Unassigned** values are:
* keys of a `map` not set
* keys of a `map` with **Null** value

Setting **Null** expresses the willing to make the attribute "unassigned" (ie. clear the attribute's value), eg:

Overriding default values when [Creating Resources](https://tools.ietf.org/html/rfc7644#section-3.3)
> Clients that intend to override existing or server-defaulted values for attributes MAY specify "null" for a single-valued attribute or an empty array "[]" for a multi-valued attribute to clear all values.

[Replacing with PUT](https://tools.ietf.org/html/rfc7644#section-3.5.1)
> Clients that want to override a server's defaults MAY specify "null" for a single-valued attribute, or an empty array "[]" for a multi-valued attribute, to clear all values.

Assuming:
```go
// m is a map[string]interface{}
v, ok := m[key]
```

Use:

* `core.IsNull(v)` to check if a value is **Null**
* `!core.IsNull(v)` to check if an attribute's value is **Unassigned**
* `ok && core.IsNull(v)` to check if an attribute's **value must be cleared**


