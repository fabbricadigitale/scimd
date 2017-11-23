package filter

import (
	"encoding/json"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/fabbricadigitale/scimd/api/attr"
)

func CompileString(s string) (Filter, error) {

	stream := antlr.NewInputStream(s)
	lexer := NewFilterLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewFilterParser(tokens)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	ctx := parser.Root()
	/*
		symbols := lexer.GetSymbolicNames()

		for _, tkn := range tokens.GetAllTokens() {
			sym := "//"
			if t := tkn.GetTokenType(); t >= 0 {
				sym = symbols[t]
			}
			fmt.Printf("%+v \t\t => %s\n", tkn, sym)
		}
	*/
	f := compileFilter(ctx.GetChild(0).(IFilterContext))

	return f, nil
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
		json.Unmarshal([]byte(vctx.GetText()), &value)
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
		f = Or{
			*compileAttributeExpression(c.GetLeft()),
			*compileAttributeExpression(c.GetRight()),
		}

	case *NotValueFilterContext:
		f = Not{
			compileValueFilter(c.GetInnerFilter()),
		}

	case *GroupValueFilterContext:
		f = Group{
			compileValueFilter(c.GetInnerFilter()),
		}
	}

	return f
}
