package schemas

import "regexp"

// URIExpr is the source text used to compile a regular expression matching a SCIM "schema" URI
const URIExpr = `[uU][rR][nN]:[A-Za-z0-9][A-Za-z0-9-]{0,31}:[A-Za-z0-9()+,\-.:=@;$_!*'%\/?#]+` // TODO double check, case NID equals to "urn" is not permitted but regex matches

// AttrNameExpr is the source text used to compile a regular expression macching a SCIM attribute name
const AttrNameExpr = `[A-Za-z]([\-_0-9A-Za-z])*` // TODO '$' RFC 7643 includes '$' but does not include '$ref' that's ambiguous, furthermore '$' cannot be used by filtering (which grammars does not include '$')

// AttrNameRegexp is the compiled Regexp built from AttrNameExpr
var AttrNameRegexp = regexp.MustCompile("^" + AttrNameExpr + "$")

// ReferenceAttrName is the name of attribute with type "reference"
const ReferenceAttrName = "$ref"
