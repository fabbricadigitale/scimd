# About

This document is a draft that describes how `scimd` has been designed, 
it should include primary design goals, internal specifications, and implementation details. 
It may not always be completely up to date.

## Design

*TBD*

## Attributes

Within `scimd` a `core.Attribute` is the definition of a SCIM attribute. It's can be included within a `core.Schema`, has a name and some characteristics.
Also, it defines the **Data Type** of the value that can be bound to the attribute.

### Case sensitivity

All runtime operations match attribute names in a *case-sensitive* way (exact match), expect the following cases:

* Unmarshalling matches JSON keys to the names defined by the schema, preferring an exact match but also accepting a *case-insensitive* match. As result, decoded names are normalized according to its schema.
*  **(TODO)** When an attribute name is included in a request body, it should be normalized as unmarshalling does.
*  **(TODO)** Filtering should accept *case-insensitive* matches for attribute paths

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
* **SCIM Data Type** `scimd`'s `type`s *(with the same name)* are also defined within the [`datatype`](../schemas/datatype/) package
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
package datatype

// ...

type DataTyper interface {
    // ...
}
```

### Singular and Multi-Valued

By convention, the `type` of a **single-value** (for *singular attribute*) is any `type` that implements `DataTyper`, `type` thus must be one of:
```go
String, Boolean, Decimal, Integer, DateTime, Binary, Reference, Complex
```

Instead, the `type` of a **multi-value** (for *multi-valued attribute*) must be `[]DataTyper`.

To check if a value is **single** or **multi** 

> A value that's not a `DataTyper` nor `[]DataTyper` is considered a **Null** value.

### Unassigned and Null Values

The internal convention for [Section 2.5 of [RFC7643]](https://tools.ietf.org/html/rfc7643#section-2.5) implements the following rules:

**Null** values are:
* `nil` *(equivalent to null, ie. JSON `null`)*
* **multi-value** `type` with length equals to 0 *(equivalent to empty array, ie. JSON `[]`)*
* values of `type`s not included in **single-value** nor **multi-value** ones

**Unassigned** values are:
* keys of a `map` not set
* keys of a `map` holding a **Null** value

Setting **Null** expresses the willing to make the attribute "unassigned" (ie. clear the attribute's value), eg:

Overriding default values when [Creating Resources](https://tools.ietf.org/html/rfc7644#section-3.3)
> Clients that intend to override existing or server-defaulted values for attributes MAY specify "null" for a single-valued attribute or an empty array "[]" for a multi-valued attribute to clear all values.

[Replacing with PUT](https://tools.ietf.org/html/rfc7644#section-3.5.1)
> Clients that want to override a server's defaults MAY specify "null" for a single-valued attribute, or an empty array "[]" for a multi-valued attribute, to clear all values.

**Unassigned** and **Null** values have particular meaning when using `map` (`datatype.Complex` is a `map` too). Thus, assuming:
```go
// m is a map[string]interface{}
v, ok := m[key]
```

You should use:

* `datatype.IsNull(v)` to check if a value is **Null**
* `!datatype.IsNull(v)` to check if an attribute's value is NOT **Unassigned** (ie. is assigned)
* `ok && datatype.IsNull(v)` to check if an attribute's **value must be cleared**


## Resources

A resource is an artifact managed by `scimd` that's described by a Resource Type and can holds values. 

All resources:
* are `struct`
* embed `core.Common` for common attribures (ie. `Schemas`, `ID`, `Meta`, etc)
> Members of `core.Common` use Go primitive `type`s, not *Data Types*

`scimd` implements two different kinds of resources.

### Structured Resources

All members of *strucutred resource* are Go primitive `type`s, not *Data Types*.

For performance and simplicity, the resources defined by the followings schemas URI:

* `urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig`
* `urn:ietf:params:scim:schemas:core:2.0:ResourceType`
* `urn:ietf:params:scim:schemas:core:2.0:Schema`

For instance, `core.ServiceProviderConfig` (that's a *Structured Resource*) can be used to load a service provider config from a JSON file.

#### Observations

* There is not advantage in storing thos resources into a database
* Those resources can not be filtered by attribute name since they are not implemented with `map`s
* The corresponding Endpoint of each of those resources ignores query [parameters](https://tools.ietf.org/html/rfc7644#section-3.4.2) (eg., filtering, sorting) as per [RFC](https://tools.ietf.org/html/rfc7644#section-4)

### Mapped Resources

To handle any other type of resource (including *User Resource*, *Group Resource* and new ones may be defined in future), 
`scimd` uses `resource.Resource` that implements a flexible data rapresentation using Go `map` internally.

A *mapped resource*:

* still embed `core.Common` for common attribures (ie. `Schemas`, `ID`, `Meta`, etc)
* has a `map` of `datatype.Complex` indexed by schema URI
* each `datatype.Complex` (that's another `map`) can hold values needed by the bound schema

Notes:
> Members of `core.Common` use Go primitive `type`s, instead `map`s hold *Data Types*

> `resource.Resource` is **NOT** responsible to enforce attributes structure and characteristics defined by bound schemas *(other APIs are needed to accomplish that)*.

In this way, it can represent data for any schemas (or composition of them in case of extensions).

Finally, you need to know that:
* it can hold data that may not be consistent with the schemas definition
* JSON marshalling/unmarshalling will ignore extraneous attributes and will enforce Data Types according to schemas definition
* to determinate the "state" of an attribute within `map`s use [Unassigned and Null](#unassigned-and-null-values) rules



## Error handling

In [RFC 7644](https://tools.ietf.org/html/rfc7644#section-3.12) are defined the following detail error types.

| scimType | Description | Applicability | Status Code |
|----------|-------------|---------------|-------------|
| **invalidFilter** | The specified filter syntax was invalid or the specified attribute and filter comparison combination is not supported | GET, POST (Search), PATCH | 400 |
| **invalidPath** | The "path" attribute was invalid or malformed  | PATCH | 400 |
| **invalidSyntax** | The request body message is invalid or not conform to the request schema | POST (Search, Create and Bulk), PUT | 400 |
| **invalidValue** | A required value was missing, or the value specified was not compatible with the operation or attribute type | GET, POST (Create, Query), PUT, PATCH | 400 |
| **invalidVers** | The specified SCIM protocol version is not supported | ALL | 400 |
| **mutability** | The attempted modification is not compatible with the target attribute's mutability or current state (modification of an "immutable" attribute with an existing value) | PUT, PATCH | 400 |
| **noTarget** | The specified "path" did not yield an attribute or a valid attribute value. This occurs when the specified "path" value contains a filter that yields no match | PATCH | 400 |
| **sensitive** | The request cannot be completed if it contains sensitive information | GET | 403 |
| **tooMany** | The specified filter yields many more results than the server is willing to calculate or process | GET, POST (Search) | 400 |
| **uniqueness** | One or more of the attribute values are already in use or are reserved | POST (Create), PUT, PATCH | 409 |

Other errors with the related status code

| Status Code | Applicability | Description |
|-------------|---------------|-------------|
| 404 (Bad Request) | ALL | The request is unparsable, syntatically incorrect or violates schema |
| 401 (Unauthorized) | ALL | Authorization failed. Authorization header is invalid or missing |
| 403 (Forbidden) | ALL | Operation is not permitted based on the supplied authorization |
| 404 (Not Found)| ALL | Specified Resource or Endpoint doesn't exist |
| 409 (Conflict) | POST, PUT, PATCH, DELETE | The specified version number does not match the resource's latest version number, or service provide refuse to create a new duplicate resource |
| 412 (Preconditon Failed) | PUT, PATCH, DELETE | Resource has changed on the server |
| 413 (Payload to large) | POST | {"maxOperations": 1000,"maxPayloadSize": 1048576} |
| 500 (Internal error) | ALL | An Internal error |
| 501 (Not Implemented) | ALL | Service provider doesn't support the request operation |