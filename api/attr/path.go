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
	ctx  map[string]*Context
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

func (p Path) matchSchema(rt *core.ResourceType) *core.Schema {

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

func (p Path) getCtxCache(key string) (*Context, bool) {
	if p.ctx != nil {
		if c, ok := p.ctx[key]; ok {
			return c, ok
		}
	}
	return nil, false
}

func (p Path) setCtxCache(key string, ctx *Context) {
	if p.ctx == nil {
		p.ctx = map[string]*Context{key: ctx}
	} else {
		p.ctx[key] = ctx
	}
}

// Context fetches from rt a suitable Context for p, if any.
func (p Path) Context(rt *core.ResourceType) (ctx *Context) {

	key := p.String()

	defer func() {
		p.setCtxCache(key, ctx)
	}()

	// Lookup cache
	if c, ok := p.getCtxCache(key); ok {
		ctx = c
		return
	}

	if rt == nil || !p.Valid() {
		return
	}

	ctx = &Context{}

	// Try common attributes
	if p.URI == "" {
		ctx.Attribute = core.Commons().ByName(p.Name)
		if ctx.Attribute != nil {
			if p.Sub != "" {
				ctx.SubAttribute = ctx.Attribute.SubAttributes.ByName(p.Sub)
				if ctx.SubAttribute == nil {
					// Unmached path
					return nil
				}
			}
			return
		}
	}

	// Try schema attributes
	ctx.Schema = p.matchSchema(rt)
	if ctx.Schema == nil {
		// Unmached path
		return nil
	}

	ctx.Attribute = ctx.Schema.Attributes.ByName(p.Name)
	if ctx.Attribute != nil {
		if p.Sub != "" {
			ctx.SubAttribute = ctx.Attribute.SubAttributes.ByName(p.Sub)
			if ctx.SubAttribute == nil {
				// Unmached path
				return nil
			}
		}
	}

	return
}

// A Context represents a set of definitions related to a Path
type Context struct {
	Schema       *core.Schema
	Attribute    *core.Attribute
	SubAttribute *core.Attribute
}
