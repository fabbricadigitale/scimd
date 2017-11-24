package filter

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/api/attr"
)

// Filter is implemented by any value that has a String method,
// which returns a SCIM filtering expression as per https://tools.ietf.org/html/rfc7644#section-3.4.2.2,
// and IsFilter method, which differentiates from other interfaces.
type Filter interface {
	String() string
	IsFilter()
}

// ValueFilter is implemented by any value that has a String method,
// which returns a SCIM expression for filtering Complex attributes (eg. emails[type eq "work"]) as per https://tools.ietf.org/html/rfc7644#section-3.4.2.2,
// and IsValueFilter method, which differentiates from other interfaces.
type ValueFilter interface {
	String() string
	IsValueFilter()
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

func (e AttrExpr) IsFilter()      {}
func (e AttrExpr) IsValueFilter() {}

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

func (op And) IsFilter() {}

// Or implements Filter, is used to represent the logical "or"
type Or struct {
	Left  Filter
	Right Filter
}

func (op Or) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

func (op Or) IsFilter() {}

// Not implements Filter, is used to represent the logical "not"
type Not struct {
	Filter
}

func (op Not) String() string {
	return "not (" + op.Filter.String() + ")"
}

// Not implements Filter, is used to represent the precedence grouping "( )"
type Group struct {
	Filter
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

func (vp ValuePath) IsFilter() {}

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

func (op ValueAnd) IsValueFilter() {}

// ValueOr implements ValueFilter, is used to represent the logical "or" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueOr struct {
	Left  AttrExpr
	Right AttrExpr
}

func (op ValueOr) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

func (op ValueOr) IsValueFilter() {}

// ValueNot implements ValueFilter, is used to represent the logical "not" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueNot struct {
	ValueFilter
}

func (op ValueNot) String() string {
	return "not (" + op.ValueFilter.String() + ")"
}

// ValueNot implements ValueFilter, is used to represent the precedence grouping "( )" within a Complex attribute filter grouping (ie. a ValuePath).
type ValueGroup struct {
	ValueFilter
}

func (g ValueGroup) String() string {
	return "(" + g.ValueFilter.String() + ")"
}
