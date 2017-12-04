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
	p := &Path{}
	matches := attrNameExp.FindStringSubmatch(s)

	// to be valid must match ATTRNAME at least
	l := len(matches)
	if l > nameIdx {
		p.URI = matches[uriIdx]
		p.Name = matches[nameIdx]
		if l > subIdx {
			p.Sub = matches[subIdx]
		}
	}
	return p
}

// Valid returns true if p is valid attribute path
func (p Path) Valid() bool {
	return len(p.Name) > 0 && !strings.HasPrefix(strings.ToLower(p.URI), schemas.InvalidURNPrefix)
}

func (p Path) String() string {
	if !p.Valid() {
		return ""
	}
	s := p.URI
	if len(s) > 0 {
		s += ":"
	}
	s += p.Name
	if len(p.Sub) > 0 {
		s += "." + p.Sub
	}
	return s
}

// FindAttribute returns the core.Attribute matched by p within the given core.ResourceType, if any
// (todo) Deprecate this in favour of Match()
func (p Path) FindAttribute(rt *core.ResourceType) *core.Attribute {
	_, att, subAtt := p.Match(rt)

	if subAtt != nil {
		return subAtt
	}

	return att
}

// MatchSchema returns schema matched by p, if any
func (p Path) MatchSchema(rt *core.ResourceType) *core.Schema {

	if !p.Valid() || rt == nil {
		return nil
	}

	if p.URI == "" {
		return rt.GetSchema()
	}

	// (fixme) ToLower() is not enough to ensure URN-equivalence as per https://tools.ietf.org/html/rfc8141#section-3
	lcURI := strings.ToLower(p.URI)

	s := rt.GetSchema()
	if lcURI == strings.ToLower(s.ID) {
		return s
	}

	for _, s := range rt.GetSchemaExtensions() {
		if lcURI == strings.ToLower(s.ID) {
			return s
		}
	}

	return nil
}

// Match returns schema, attribute, and subAttribute matched by p
func (p Path) Match(rt *core.ResourceType) (schema *core.Schema, attribute *core.Attribute, subAttribute *core.Attribute) {

	schema = p.MatchSchema(rt)
	if schema == nil {
		return
	}

	attribute = schema.Attributes.ByName(p.Name)
	if attribute == nil {
		return
	}

	if p.Sub != "" {
		subAttribute = attribute.SubAttributes.ByName(p.Sub)
	}

	return
}
