package filter

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
)

type parserErrorListener struct {
	antlr.DefaultErrorListener
	err error
}

func (l *parserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.err = &api.InvalidFilterError{
		Detail: "syntax error at character " + strconv.Itoa(column) + ", " + msg,
	}
}

// CompileString parses a SCIM filter string and returns, if successful, a Filter object.
func CompileString(s string) (Filter, error) {

	errListener := new(parserErrorListener)
	stream := antlr.NewInputStream(s)
	lexer := NewFilterLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewFilterParser(tokens)
	parser.AddErrorListener(errListener)

	ctx := parser.Root()

	if errListener.err != nil {
		return nil, errListener.err
	}

	return compileFilter(ctx.GetChild(0).(IFilterContext)), nil
}

func compileFilter(ctx IFilterContext) Filter {
	var f Filter

	switch c := ctx.(type) {
	case *AttributeExprFilterContext:
		f = compileAttributeExpression(c.GetAttributeExpr())

	case *AndFilterContext:
		f = And{
			compileFilter(c.GetLeft()),
			compileFilter(c.GetRight()),
		}

	case *OrFilterContext:
		f = Or{
			compileFilter(c.GetLeft()),
			compileFilter(c.GetRight()),
		}

	case *ValueExprFilterContext:
		f = compileValueExpression(c.GetValueExpr())

	case *NotFilterContext:
		f = Not{
			compileFilter(c.GetInnerFilter()),
		}

	case *GroupFilterContext:
		f = Group{
			compileFilter(c.GetInnerFilter()),
		}
	default:
		panic("filter: unexpected context")
	}

	return f
}

func compileAttributePath(c IAttributePathContext) *attr.Path {
	path := attr.Path{}
	if uri := c.GetURI(); uri != nil {
		path.URI = uri.GetText()
	}
	if name := c.GetName(); name != nil {
		path.Name = name.GetText()
	}
	if sub := c.GetSub(); sub != nil {
		path.Sub = sub.GetText()
	}
	return &path
}

func compileAttributeExpression(c IAttributeExpressionContext) *AttrExpr {

	// Path
	pctx := c.GetPath()
	path := compileAttributePath(pctx)

	// Operator
	octx := c.GetOp()
	var op string
	if octx != nil {
		op = strings.TrimSpace(octx.GetText())
	}

	// Value
	vctx := c.GetValue()
	var value interface{}
	if vctx != nil {
		if err := json.Unmarshal([]byte(vctx.GetText()), &value); err != nil {
			panic(err)
		}
	}

	return &AttrExpr{
		Path:  *path,
		Op:    op,
		Value: value,
	}
}

func compileValueExpression(c IValueExpressionContext) *ValuePath {
	return &ValuePath{
		Path:        *compileAttributePath(c.GetPath()),
		ValueFilter: compileValueFilter(c.GetInnerFilter()),
	}
}

func compileValueFilter(ctx IValueFilterContext) ValueFilter {
	var f ValueFilter

	switch c := ctx.(type) {
	case *AttributeExprValueFilterContext:
		f = compileAttributeExpression(c.GetAttributeExpr())

	case *AndValueFilterContext:
		f = ValueAnd{
			*compileAttributeExpression(c.GetLeft()),
			*compileAttributeExpression(c.GetRight()),
		}

	case *OrValueFilterContext:
		f = ValueOr{
			*compileAttributeExpression(c.GetLeft()),
			*compileAttributeExpression(c.GetRight()),
		}

	case *NotValueFilterContext:
		f = ValueNot{
			compileValueFilter(c.GetInnerFilter()),
		}

	case *GroupValueFilterContext:
		f = ValueGroup{
			compileValueFilter(c.GetInnerFilter()),
		}
	default:
		panic("filter: unexpected context")
	}

	return f
}
