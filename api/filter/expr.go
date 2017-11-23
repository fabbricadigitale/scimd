package filter

import (
	"encoding/json"

	"github.com/fabbricadigitale/scimd/api/attr"
)

type Filter interface {
	String() string
}

type ValueFilter interface {
	String() string
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
)

var _ Filter = (*AttrExpr)(nil)
var _ ValueFilter = (*AttrExpr)(nil)

// AttrExpr is an attribute expression
type AttrExpr struct {
	Path  attr.Path
	Op    string
	Value interface{}
}

func (e AttrExpr) String() string {
	compValue, _ := json.Marshal(e.Value)
	return e.Path.String() + " " + e.Op + " " + string(compValue)
}

// Logical Expression
var _ Filter = (*And)(nil)
var _ Filter = (*Or)(nil)
var _ Filter = (*Not)(nil)
var _ Filter = (*Group)(nil)

type And struct {
	Left  Filter
	Right Filter
}

func (op And) String() string {
	return op.Left.String() + " and " + op.Right.String()
}

type Or struct {
	Left  Filter
	Right Filter
}

func (op Or) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

type Not struct {
	Filter
}

func (op Not) String() string {
	return "not (" + op.Filter.String() + ")"
}

type Group struct {
	Filter
}

func (g Group) String() string {
	return "(" + g.Filter.String() + ")"
}

// Value Path
var _ Filter = (*ValuePath)(nil)

type ValuePath struct {
	Path attr.Path
	ValueFilter
}

func (vp ValuePath) String() string {
	return vp.Path.String() + "[" + vp.ValueFilter.String() + "]"
}

// Logical operators for value filtering
var _ ValueFilter = (*ValueAnd)(nil)
var _ ValueFilter = (*ValueOr)(nil)
var _ ValueFilter = (*ValueNot)(nil)
var _ ValueFilter = (*ValueGroup)(nil)

type ValueAnd struct {
	Left  AttrExpr
	Right AttrExpr
}

func (op ValueAnd) String() string {
	return op.Left.String() + " and " + op.Right.String()
}

type ValueOr struct {
	Left  AttrExpr
	Right AttrExpr
}

func (op ValueOr) String() string {
	return op.Left.String() + " or " + op.Right.String()
}

type ValueNot struct {
	ValueFilter
}

func (op ValueNot) String() string {
	return "not (" + op.ValueFilter.String() + ")"
}

type ValueGroup struct {
	ValueFilter
}

func (g ValueGroup) String() string {
	return "(" + g.ValueFilter.String() + ")"
}
