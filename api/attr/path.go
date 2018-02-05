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
	subAttr = `(?:\.(?P<SUBATTRNAME>\` + schemas.ReferenceAttrName + `|` + schemas.AttrNameExpr + `))?`
)

// A Path represents a parsed SCIM attribute path as per https://tools.ietf.org/html/rfc7644#section-3.10
type Path struct {
	URI  string
	Name string
	Sub  string
}

var (
	attrNameExp = regexp.MustCompile("(?::" + attrName + subAttr + ")$")
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

// Undefined returns true if p does not satisfy the minimal path attribute definition.
// An Path is defined when its minimal component (ie., Name) exists, that's not implay validity
// (syntactically validation is a Parse responsibility).
func (p Path) Undefined() bool {
	return len(p.Name) == 0
}

func (p Path) String() string {
	if p.Undefined() {
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

// IsSubAttribute returns true if the path refers to an attribute with a parent
//
// While p.Name refers to the name of the parent, p.Sub refers to the name of its children.
func (p Path) IsSubAttribute() bool {
	return !p.Undefined() && len(p.Sub) > 0
}

// Transform applies f(facet) to p's facets (URN, attribute, and sub-attribute name) and returns a new transformed Path
func (p Path) Transform(f func(facet string) string) *Path {
	return &Path{
		URI:  f(p.URI),
		Name: f(p.Name),
		Sub:  f(p.Sub),
	}
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

// Paths returns a slice of Path given a resource type rt.
//
// It flattens the attributes of rt's schemas returning their contextualized Path representations.
// When a fx is provided it returns only the attribute paths statisfying fx(attribute).
func Paths(rt *core.ResourceType, fx func(attribute *core.Attribute) bool) []*Path {
	ctxs := Contexts(rt, fx)
	acc := make([]*Path, len(ctxs))

	for i, ctx := range ctxs {
		acc[i] = ctx.Path()
	}

	return acc
}
