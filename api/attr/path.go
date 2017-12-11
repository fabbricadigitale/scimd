package attr

import (
	"regexp"
	"strings"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	urn "github.com/leodido/go-urn"
)

const (
	// nameChar  = "-" / "_" / DIGIT / ALPHA
	// ATTRNAME  = ALPHA *(nameChar)
	attrName = `(?P<ATTRNAME>` + schemas.AttrNameExpr + `)`

	// subAttr   = "." ATTRNAME ; a sub-attribute of a complex attribute
	subAttr = `(?:\.(?P<SUBATTRNAME>` + schemas.AttrNameExpr + `))?`
)

// A Path represents a parsed SCIM attribute path as per https://tools.ietf.org/html/rfc7644#section-3.10
type Path struct {
	URI  string
	Name string
	Sub  string
	ctx  map[string]*Context
}

var (
	attrNameExp = regexp.MustCompile("(?::" + attrName + subAttr + ")$")
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

	var in string
	// Parse input string as an URN
	u, ok := urn.Parse(s)
	// Switch input
	if ok {
		// Grab specific string
		in = u.SS
	} else {
		// Assume input was <attrname>.<subattrname>
		in = ":" + s
	}

	// Parse current input as attribute name
	matches := attrNameExp.FindStringSubmatch(in)

	// Any valid attribute name give us always 3 matches:
	// * the full match
	// * the attribute name
	// * the subattribute name (empty string when not present)
	//
	// When we found an attribute name expression ...
	if len(matches) == 3 {
		// And primary input was an URN
		if ok {
			// Remove attribute name from the URN's specific string
			u.SS = strings.TrimSuffix(u.SS, matches[0])
			// Normalize current URN and store it
			p.URI = u.Normalize().String()
			// Store attribute name and subattribute name
			p.Name = matches[1]
			p.Sub = matches[2]
		} else if in == matches[0] {
			// Store attribute name and subattribute name also when:
			// original input was not an URN but full match equals our internal input,
			// which means it was a syntactically valid attribute name expression.
			p.Name = matches[1]
			p.Sub = matches[2]
		}
	}

	return p
}

// Valid returns true if p is valid attribute path
// (todo) > rename in defined or undefined (negating condition)?
// (note) > an attr exist iff its minimal component (ie., name) exists and it is syntactically valid (parse responsibility)
func (p Path) Valid() bool {
	return len(p.Name) > 0
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

	s := rt.GetSchema()
	// Simple equivalence assuming schema ID is a valid already normalized URN
	if p.URI == s.ID {
		return s
	}

	for _, s := range rt.GetSchemaExtensions() {
		// Same assumption as above
		if p.URI == s.ID {
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
