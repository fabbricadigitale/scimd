package schemas

import "regexp"

// URIExpr is the source text used to compile a regular expression matching a SCIM "schema" URI
const URIExpr = `[uU][rR][nN]:[A-Za-z0-9][A-Za-z0-9-]{0,31}:[A-Za-z0-9()+,\-.:=@;$_!*'%\/?#]+` // TODO double check, case NID equals to "urn" is not permitted but regex matches

// InvalidURNPrefix is the text used to validate the Path, NID cannot be equal to "urn"
const InvalidURNPrefix = "urn:urn:"

// URIRegexp is the compiled Regexp built from URIExpr
var URIRegexp = regexp.MustCompile("^" + URIExpr + "$")

// AttrNameExpr is the source text used to compile a regular expression macching a SCIM attribute name
const AttrNameExpr = `[A-Za-z]([\-_0-9A-Za-z])*` // TODO '$' RFC 7643 includes '$' but does not include '$ref' that's ambiguous, furthermore '$' cannot be used by filtering (which grammars does not include '$')

// AttrNameRegexp is the compiled Regexp built from AttrNameExpr
var AttrNameRegexp = regexp.MustCompile("^" + AttrNameExpr + "$")

// ReferenceAttrName is the name of attribute with type "reference"
const ReferenceAttrName = "$ref"

// ComplexValueAttrName is the name of complex's sub-attribute that represents the attribute's significant value.
const ComplexValueAttrName = "value"

// Keyword indicating the circumstances under which the value of the attribute can be (re)defined
const (
	MutabilityReadOnly  = "readOnly"
	MutabilityReadWrite = "readWrite" // the default
	MutabilityImmutable = "immutable"
	MutabilityWriteOnly = "writeOnly"
)

// Keyword that indicates when an attribute and associated values are returned in response
const (
	ReturnedAlways  = "always"
	ReturnedNever   = "never"
	ReturnedDefault = "default" // the default
	ReturnedRequest = "request"
)

// Keyword value that specifies how the service provider enforces uniqueness of attribute values
const (
	UniquenessNone   = "none" // the default
	UniquenessServer = "server"
	UniquenessGlobal = "global"
)
