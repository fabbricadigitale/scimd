package attr

import (
	"regexp"
	"strings"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

const (
	// attrPath  = [URI ":"] ATTRNAME *1subAttr ; SCIM attribute name ; URI is SCIM "schema" URI
	attrPath = `((?P<URI>` + schemas.URIExpr + `)\:)?` + attrName + subAttr + `?`

	// nameChar  = "-" / "_" / DIGIT / ALPHA
	// ATTRNAME  = ALPHA *(nameChar)
	attrName = `(?P<ATTRNAME>(` + schemas.AttrNameExpr + `))`

	// subAttr   = "." ATTRNAME ; a sub-attribute of a complex attribute
	subAttr = `(\.(?P<SUBATTRNAME>` + schemas.AttrNameExpr + `))`
)

// A Path represents a parsed SCIM attribute path as per https://tools.ietf.org/html/rfc7644#section-3.10
type Path struct {
	URI  string
	Name string
	Sub  string
}

var (
	attrNameExp = regexp.MustCompile("^" + attrPath + "$")
)

// TODO automatize with attrNameExp.SubexpNames()
const (
	uriIdx  = 2
	nameIdx = 3
	subIdx  = 7
)

// Parse parses a SCIM attribute notation into a Path structure.
func Parse(s string) *Path {
	a := &Path{}
	matches := attrNameExp.FindStringSubmatch(s)

	// to be valid must match ATTRNAME at least
	l := len(matches)
	if l > nameIdx {
		a.URI = matches[uriIdx]
		a.Name = matches[nameIdx]
		if l > subIdx {
			a.Sub = matches[subIdx]
		}
	}
	return a
}

// Valid returns true if a is valid attribute path
func (a Path) Valid() bool {
	return len(a.Name) > 0 && !strings.HasPrefix(strings.ToLower(a.URI), schemas.InvalidURNPrefix)
}

func (a Path) String() string {
	if !a.Valid() {
		return ""
	}
	s := a.URI
	if len(s) > 0 {
		s += ":"
	}
	s += a.Name
	if len(a.Sub) > 0 {
		s += "." + a.Sub
	}
	return s
}

func (a *Path) findAttrDef(s *core.Schema) *core.Attribute {
	def := s.Attributes.ByName(a.Name)
	if def != nil && a.Sub != "" {
		def = def.SubAttributes.ByName(a.Sub)
	}
	return def
}

// FindAttribute returns the core.Attribute matched by Path within the given core.ResourceType, if any
func (a Path) FindAttribute(rt *core.ResourceType) *core.Attribute {

	if !a.Valid() {
		return nil
	}

	s := rt.GetSchema()

	// if no URI, assume base schema
	if a.URI == "" {
		return a.findAttrDef(s)
	}

	// (fixme) ToLower() is not enough to ensure URN-equivalence as per https://tools.ietf.org/html/rfc8141#section-3
	lcURI := strings.ToLower(a.URI)

	if lcURI == strings.ToLower(s.ID) {
		return a.findAttrDef(s)
	}

	for _, s := range rt.GetSchemaExtensions() {
		if lcURI == strings.ToLower(s.ID) {
			return a.findAttrDef(s)
		}
	}

	return nil
}
