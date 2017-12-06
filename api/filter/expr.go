package filter

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// Filter is implemented by any value that has a String method,
// which returns a SCIM filtering expression as per https://tools.ietf.org/html/rfc7644#section-3.4.2.2,
// and Normalize method, which differentiates from other interfaces.
type Filter interface {
	String() string
	// Normalize returns the contestualized form of Filter for the given rt
	Normalize(rt *core.ResourceType) Filter
}

// ValueFilter is implemented by any value that has a String method,
// which returns a SCIM expression for filtering Complex attributes (eg. emails[type eq "work"]) as per https://tools.ietf.org/html/rfc7644#section-3.4.2.2,
// and ToFilter method, which differentiates from other interfaces.
type ValueFilter interface {
	String() string
	// ToFilter returns an equivalent and normalized Filter assiming ValueFilter within the given ctx
	ToFilter(ctx *attr.Context) Filter
}

// Attribute Operators
const (
	OpEqual              = "eq"
	OpNotEqual           = "ne"
	OpContains           = "co"
	OpStartsWith         = "sw"
	OpEndsWith           = "ew"
	OpGreaterThan        = "gt"
	OpLessThan           = "lt"
	OpGreaterOrEqualThan = "ge"
	OpLessOrEqualThan    = "le"
	OpPresent            = "pr"
)

var opMap = map[string]bool{
	OpEqual:              true,
	OpNotEqual:           true,
	OpContains:           true,
	OpStartsWith:         true,
	OpEndsWith:           true,
	OpGreaterThan:        true,
	OpLessThan:           true,
	OpGreaterOrEqualThan: true,
	OpLessOrEqualThan:    true,
	OpPresent:            true,
}

var _ Filter = (*AttrExpr)(nil)
var _ ValueFilter = (*AttrExpr)(nil)

// AttrExpr is an attribute expression, implements both Filter and ValueFilter.
// When OpPresent is used, Value should be nil
type AttrExpr struct {
	Path  attr.Path
	Op    string      // Should be an Attribute Operator
	Value interface{} // Type allowed: nil, bool, string, int64 or float64
}

func (e AttrExpr) String() string {
	if e.Op == OpPresent {
		return e.Path.String() + " " + e.Op
	}
	compValue, _ := json.Marshal(e.Value)
	return e.Path.String() + " " + e.Op + " " + string(compValue)
}

func (e AttrExpr) Normalize(rt *core.ResourceType) Filter {
	ctx := e.Path.Context(rt)

	var p attr.Path
	o := e.Op
	v := e.Value

	if ctx != nil && ctx.Schema != nil && ctx.Attribute != nil {

		p = attr.Path{
			URI:  ctx.Schema.ID,
			Name: ctx.Attribute.Name,
		}

		if ctx.SubAttribute == nil {
			if ctx.Attribute.Type == datatype.ComplexType {
				if a := ctx.Attribute.SubAttributes.ByName(schemas.ComplexValueAttrName); a != nil {
					p.Sub = a.Name
				}
			}
		} else {
			p.Sub = ctx.SubAttribute.Name
		}

	} else {
		// For filtered attributes that are not part of a particular resource
		// type, the service provider SHALL treat the attribute as if there is
		// no attribute value, as per https://tools.ietf.org/html/rfc7644#section-3.4.2.1
		p = attr.Path{} // using zero value to indicate an undefined and not valid attribute path
	}

	if !opMap[o] {
		panic("Invalid attribute expression: operator '" + o + "' is not supported")
	}

	if o == OpPresent {
		v = nil
	}

	exp := &AttrExpr{
		Path:  p,
		Op:    o,
		Value: v,
	}

	return exp
}

func (e AttrExpr) ToFilter(ctx *attr.Context) Filter {

	p := attr.Path{}

	// For filtered attributes that are not part of a particular resource
	// type, the service provider SHALL treat the attribute as if there is
	// no attribute value, as per https://tools.ietf.org/html/rfc7644#section-3.4.2.1
	// Futhermore, complex's sub-attribute cannot have sub-attribute, so ignore them
	if ctx != nil && ctx.Schema != nil && ctx.Attribute != nil && ctx.SubAttribute == nil && e.Path.Valid() {
		if e.Path.URI != "" || e.Path.Sub != "" {
			panic(&api.InvalidFilterError{
				Filter: e.String(),
				Detail: "paths with URI or sub-attribute are not supported within a complex attribute filter grouping",
			})
		}

		if ctx.Attribute.Type != datatype.ComplexType {
			panic(&api.InvalidFilterError{
				Filter: e.String(),
				Detail: "complex attribute filter grouping not allowed for non-complex attributes",
			})
		}

		leaf := ctx.Attribute.SubAttributes.ByName(e.Path.Name)
		if leaf != nil {
			p.URI = ctx.Schema.ID
			p.Name = ctx.Attribute.Name
			p.Sub = leaf.Name
		}
		// else p will left to its zero value to indicate an undefined and not valid attribute path
	}

	return &AttrExpr{
		Path:  p,
		Op:    e.Op,
		Value: e.Value,
	}

}

// Logical Expression
var _ Filter = (*And)(nil)
var _ Filter = (*Or)(nil)
var _ Filter = (*Not)(nil)
var _ Filter = (*Group)(nil)

// And implements Filter, is used to represent the logical "and"
type And struct {
	Left  Filter
	Right Filter
}

func (op And) String() string {
	return op.Left.String() + " and " + op.Right.String()
}

func (op And) Normalize(rt *core.ResourceType) Filter {
	return &And{
		Left:  op.Left.Normalize(rt),
		Right: op.Right.Normalize(rt),
	}
}

// Or implements Filter, is used to represent the logical "or"
type Or struct {
	Left  Filter
	Right Filter
}

func (op Or) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

func (op Or) Normalize(rt *core.ResourceType) Filter {
	return &Or{
		Left:  op.Left.Normalize(rt),
		Right: op.Right.Normalize(rt),
	}
}

// Not implements Filter, is used to represent the logical "not"
type Not struct {
	Filter
}

func (op Not) Normalize(rt *core.ResourceType) Filter {
	return &Not{
		op.Filter.Normalize(rt),
	}
}

func (op Not) String() string {
	return "not (" + op.Filter.String() + ")"
}

// Not implements Filter, is used to represent the precedence grouping "( )"
type Group struct {
	Filter
}

func (op Group) Normalize(rt *core.ResourceType) Filter {
	return &Group{
		op.Filter.Normalize(rt),
	}
}

func (g Group) String() string {
	return "(" + g.Filter.String() + ")"
}

// Value Path
var _ Filter = (*ValuePath)(nil)

// ValuePath implements Filter, is used to represent a value path expression (eg. emails[type eq "work" and value co "@example.com"]).
// The filter for Complex attribute filter grouping MUST implement the ValueFilter interface.
type ValuePath struct {
	Path attr.Path
	ValueFilter
}

func (vp ValuePath) String() string {
	return vp.Path.String() + "[" + vp.ValueFilter.String() + "]"
}

func (vp ValuePath) Normalize(rt *core.ResourceType) Filter {
	ctx := vp.Path.Context(rt)
	return &Group{
		Filter: vp.ValueFilter.ToFilter(ctx),
	}
}

// Logical operators for value filtering
var _ ValueFilter = (*ValueAnd)(nil)
var _ ValueFilter = (*ValueOr)(nil)
var _ ValueFilter = (*ValueNot)(nil)
var _ ValueFilter = (*ValueGroup)(nil)

// ValueAnd implements ValueFilter, is used to represent the logical "and" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueAnd struct {
	Left  AttrExpr
	Right AttrExpr
}

func (op ValueAnd) String() string {
	return op.Left.String() + " and " + op.Right.String()
}

func (op ValueAnd) ToFilter(ctx *attr.Context) Filter {
	return &And{
		Left:  op.Left.ToFilter(ctx),
		Right: op.Right.ToFilter(ctx),
	}
}

// ValueOr implements ValueFilter, is used to represent the logical "or" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueOr struct {
	Left  AttrExpr
	Right AttrExpr
}

func (op ValueOr) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

func (op ValueOr) ToFilter(ctx *attr.Context) Filter {
	return &Or{
		Left:  op.Left.ToFilter(ctx),
		Right: op.Right.ToFilter(ctx),
	}
}

// ValueNot implements ValueFilter, is used to represent the logical "not" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueNot struct {
	ValueFilter
}

func (op ValueNot) ToFilter(ctx *attr.Context) Filter {
	return &Not{
		op.ValueFilter.ToFilter(ctx),
	}
}

func (op ValueNot) String() string {
	return "not (" + op.ValueFilter.String() + ")"
}

// ValueGroup implements ValueFilter, is used to represent the precedence grouping "( )" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueGroup struct {
	ValueFilter
}

func (op ValueGroup) ToFilter(ctx *attr.Context) Filter {
	return &Group{
		op.ValueFilter.ToFilter(ctx),
	}
}

func (g ValueGroup) String() string {
	return "(" + g.ValueFilter.String() + ")"
}
