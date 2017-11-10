# About

This document is a draft that describes how `scimd` has been designed, 
it should include primary design goals, internal specifications, and implementation details. 
It may not always be completely up to date.

## Design

*TBD*

## Attributes

Within `scimd` a `core.Attribute` is the definition of a SCIM attribute. It's can be included within a `core.Schema`, has a name and some characteristics.
Also, it defines the **Data Type** of the value that can be bound to the attribute.

### Data Types

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

Rules: 
* **JSON Type** must be used when encoding/decoding
* **SCIM Schema "type"** is the set of values that `core.Attribute.Type` can assume
* **SCIM Data Type** `scimd`'s `type`s *(with the same name)* are also defined within the [`core`](../schemas/core/data_type.go) package
* **Golang `type`** are just the underlying Go `type` (ie. not use them directly)

Indeed:
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

### Singular and Multi-Valued

By convention, the `type` of a **single-value** (for *singular attribute*) must be one of:
```go
String, Boolean, Decimal, Integer, DateTime, Binary, Reference, Complex
```
Instead, the `type` of a **multi-value** (for *multi-valued attribute*) must be on of:
```go
[]String, []Boolean, []Decimal, []Integer, []DateTime, []Binary, []Reference, []Complex
```

> A value of another `type` is considered a **Null** value.

### Unassigned and Null Values

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

**Unassigned** and **Null** values have particular meaning when using `map` (`core.Complex` is a `map` too). Thus, assuming:
```go
// m is a map[string]interface{}
v, ok := m[key]
```

You should use:

* `core.IsNull(v)` to check if a value is **Null**
* `!core.IsNull(v)` to check if an attribute's value is **Unassigned**
* `ok && core.IsNull(v)` to check if an attribute's **value must be cleared**


## Resources

A resource is an artifact managed by `scimd` that's described by a Resource Type and can holds values. 
`scimd` implements two different kinds of resources.

### Structured Resources

For performance and simplicity, the resources defined by the followings schemas URI:

* `urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig`
* `urn:ietf:params:scim:schemas:core:2.0:ResourceType`
* `urn:ietf:params:scim:schemas:core:2.0:Schema`

are implemented by Go `struct`, which are very simply to use.

For instance, `core.ServiceProviderConfig` (that's a *Structured Resource*) can be used to load a service provider config from a JSON file.

### Mapped Resources

To handle any other type of resource (including *User Resource*, *Group Resource* and new ones may be defined in future), 
`scimd` uses `core.resource.Resource` that implements a flexible data rapresentation using Go `map` internally.

`core.resource.Resource` implements the following features:

* it's a Go `struct` 
* common attribures are embedded within the `struct` directly (ie. `Schemas`, `ID`, `Meta`, etc)
* has a `map` of `core.Complex` indexed by schema URI
* each `core.Complex` (that's another `map`) can hold values needed by the bound schema

In this way, it can represent data for any schemas (or composition of them in case of extensions).

`core.resource.Resource` is **NOT** responsible to enforce attributes structure and characteristics defined by bound schemas *(other APIs are needed to accomplish that)*.

Finally, you need to know that:
* it can hold data that may not be consistent with the schemas definition
* JSON marshalling/unmarshalling will ignore extraneous attributes and will enforce Data Types according to schemas definition
* to determinate the "state" of an attribute within `map`s use [Unassigned and Null](#unassigned-and-null-values) rules
