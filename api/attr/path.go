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
	ctx  map[string]Context
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
	c := p.Context(rt)

	if c.SubAttribute != nil {
		return c.SubAttribute
	}

	return c.Attribute
}

func (p Path) matchSchema(rt *core.ResourceType) *core.Schema {

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

func (p Path) getCtxCache(key string) *Context {
	if p.ctx != nil {
		if c, ok := p.ctx[key]; ok {
			return &c
		}
	}
	return nil
}

func (p Path) setCtxCache(key string, ctx *Context) {
	if p.ctx == nil {
		p.ctx = map[string]Context{key: *ctx}
	} else {
		p.ctx[key] = *ctx
	}
}

// Context fetches from rt a suitable Context for p, if any.
func (p Path) Context(rt *core.ResourceType) *Context {

	key := p.String()

	// Lookup cache
	if c := p.getCtxCache(key); c != nil {
		return c
	}

	schema := p.matchSchema(rt)
	if schema == nil {
		return nil
	}

	attribute := schema.Attributes.ByName(p.Name)
	if attribute == nil {
		return nil
	}

	c := &Context{
		Schema:    schema,
		Attribute: attribute,
	}

	if p.Sub != "" {
		c.SubAttribute = attribute.SubAttributes.ByName(p.Sub)
	}

	p.setCtxCache(key, c)
	return c
}

// A Context represents a set of definitions related to a Path
type Context struct {
	Schema       *core.Schema
	Attribute    *core.Attribute
	SubAttribute *core.Attribute
}
