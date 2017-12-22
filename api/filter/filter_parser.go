package filter // FilterParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)



import "github.com/fabbricadigitale/scimd/api/attr"


// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 88, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3, 
	2, 5, 2, 16, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 7, 3, 39, 10, 3, 12, 3, 14, 3, 42, 11, 3, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 4, 3, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 5, 5, 54, 10, 5, 3, 5, 3, 5, 
	3, 5, 5, 5, 59, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 86, 10, 7, 3, 7, 2, 3, 4, 8, 2, 4, 
	6, 8, 10, 12, 2, 3, 3, 2, 9, 17, 2, 94, 2, 15, 3, 2, 2, 2, 4, 30, 3, 2, 
	2, 2, 6, 50, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 85, 
	3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 
	16, 17, 3, 2, 2, 2, 17, 18, 7, 2, 2, 3, 18, 3, 3, 2, 2, 2, 19, 20, 8, 3, 
	1, 2, 20, 31, 5, 6, 4, 2, 21, 31, 5, 10, 6, 2, 22, 23, 7, 20, 2, 2, 23, 
	24, 5, 4, 3, 2, 24, 25, 7, 21, 2, 2, 25, 31, 3, 2, 2, 2, 26, 27, 7, 22, 
	2, 2, 27, 28, 5, 4, 3, 2, 28, 29, 7, 21, 2, 2, 29, 31, 3, 2, 2, 2, 30, 
	19, 3, 2, 2, 2, 30, 21, 3, 2, 2, 2, 30, 22, 3, 2, 2, 2, 30, 26, 3, 2, 2, 
	2, 31, 40, 3, 2, 2, 2, 32, 33, 12, 7, 2, 2, 33, 34, 7, 18, 2, 2, 34, 39, 
	5, 4, 3, 8, 35, 36, 12, 6, 2, 2, 36, 37, 7, 19, 2, 2, 37, 39, 5, 4, 3, 
	7, 38, 32, 3, 2, 2, 2, 38, 35, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 
	3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 5, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 
	43, 44, 5, 8, 5, 2, 44, 45, 7, 8, 2, 2, 45, 51, 3, 2, 2, 2, 46, 47, 5, 
	8, 5, 2, 47, 48, 9, 2, 2, 2, 48, 49, 7, 4, 2, 2, 49, 51, 3, 2, 2, 2, 50, 
	43, 3, 2, 2, 2, 50, 46, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 54, 7, 6, 2, 
	2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 58, 
	7, 3, 2, 2, 56, 57, 7, 24, 2, 2, 57, 59, 7, 3, 2, 2, 58, 56, 3, 2, 2, 2, 
	58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 8, 5, 1, 2, 61, 62, 6, 
	5, 4, 3, 62, 9, 3, 2, 2, 2, 63, 64, 5, 8, 5, 2, 64, 65, 7, 25, 2, 2, 65, 
	66, 5, 12, 7, 2, 66, 67, 7, 23, 2, 2, 67, 11, 3, 2, 2, 2, 68, 86, 5, 6, 
	4, 2, 69, 70, 5, 6, 4, 2, 70, 71, 7, 18, 2, 2, 71, 72, 5, 6, 4, 2, 72, 
	86, 3, 2, 2, 2, 73, 74, 5, 6, 4, 2, 74, 75, 7, 19, 2, 2, 75, 76, 5, 6, 
	4, 2, 76, 86, 3, 2, 2, 2, 77, 78, 7, 20, 2, 2, 78, 79, 5, 12, 7, 2, 79, 
	80, 7, 21, 2, 2, 80, 86, 3, 2, 2, 2, 81, 82, 7, 22, 2, 2, 82, 83, 5, 12, 
	7, 2, 83, 84, 7, 21, 2, 2, 84, 86, 3, 2, 2, 2, 85, 68, 3, 2, 2, 2, 85, 
	69, 3, 2, 2, 2, 85, 73, 3, 2, 2, 2, 85, 77, 3, 2, 2, 2, 85, 81, 3, 2, 2, 
	2, 86, 13, 3, 2, 2, 2, 10, 15, 30, 38, 40, 50, 53, 58, 85,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "':'", "' pr'", "' eq'", "' ne'", "' co'", "' sw'", 
	"' ew'", "' gt'", "' ge'", "' lt'", "' le'", "' and '", "' or '", "", "')'", 
	"'('", "']'", "'.'", "'['",
}
var symbolicNames = []string{
	"", "AttributeName", "ComparisonValue", "Space", "Urn", "Colon", "PrOperator", 
	"EqOperator", "NeOperator", "CoOperator", "SwOperator", "EwOperator", "GtOperator", 
	"GeOperator", "LtOperator", "LeOperator", "AndOperator", "OrOperator", 
	"NotOperator", "RxBracket", "LxBracket", "RxSquareBracket", "Dot", "LxSquareBracket",
}

var ruleNames = []string{
	"root", "filter", "attributeExpression", "attributePath", "valueExpression", 
	"valueFilter",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FilterParser struct {
	*antlr.BaseParser
}

func NewFilterParser(input antlr.TokenStream) *FilterParser {
	this := new(FilterParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FilterParser.g4"

	return this
}

// FilterParser tokens.
const (
	FilterParserEOF = antlr.TokenEOF
	FilterParserAttributeName = 1
	FilterParserComparisonValue = 2
	FilterParserSpace = 3
	FilterParserUrn = 4
	FilterParserColon = 5
	FilterParserPrOperator = 6
	FilterParserEqOperator = 7
	FilterParserNeOperator = 8
	FilterParserCoOperator = 9
	FilterParserSwOperator = 10
	FilterParserEwOperator = 11
	FilterParserGtOperator = 12
	FilterParserGeOperator = 13
	FilterParserLtOperator = 14
	FilterParserLeOperator = 15
	FilterParserAndOperator = 16
	FilterParserOrOperator = 17
	FilterParserNotOperator = 18
	FilterParserRxBracket = 19
	FilterParserLxBracket = 20
	FilterParserRxSquareBracket = 21
	FilterParserDot = 22
	FilterParserLxSquareBracket = 23
)

// FilterParser rules.
const (
	FilterParserRULE_root = 0
	FilterParserRULE_filter = 1
	FilterParserRULE_attributeExpression = 2
	FilterParserRULE_attributePath = 3
	FilterParserRULE_valueExpression = 4
	FilterParserRULE_valueFilter = 5
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(FilterParserEOF, 0)
}

func (s *RootContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *FilterParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterParserRULE_root)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FilterParserAttributeName) | (1 << FilterParserUrn) | (1 << FilterParserNotOperator) | (1 << FilterParserLxBracket))) != 0) {
		{
			p.SetState(12)
			p.filter(0)
		}

	}
	{
		p.SetState(15)
		p.Match(FilterParserEOF)
	}



	return localctx
}


// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) CopyFrom(ctx *FilterContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type AndFilterContext struct {
	*FilterContext
	Left IFilterContext 
	Right IFilterContext 
}

func NewAndFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndFilterContext {
	var p = new(AndFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *AndFilterContext) GetLeft() IFilterContext { return s.Left }

func (s *AndFilterContext) GetRight() IFilterContext { return s.Right }


func (s *AndFilterContext) SetLeft(v IFilterContext) { s.Left = v }

func (s *AndFilterContext) SetRight(v IFilterContext) { s.Right = v }

func (s *AndFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndFilterContext) AndOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserAndOperator, 0)
}

func (s *AndFilterContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *AndFilterContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}



type ValueExprFilterContext struct {
	*FilterContext
	ValueExpr IValueExpressionContext 
}

func NewValueExprFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprFilterContext {
	var p = new(ValueExprFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *ValueExprFilterContext) GetValueExpr() IValueExpressionContext { return s.ValueExpr }


func (s *ValueExprFilterContext) SetValueExpr(v IValueExpressionContext) { s.ValueExpr = v }

func (s *ValueExprFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprFilterContext) ValueExpression() IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}



type NotFilterContext struct {
	*FilterContext
	InnerFilter IFilterContext 
}

func NewNotFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotFilterContext {
	var p = new(NotFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *NotFilterContext) GetInnerFilter() IFilterContext { return s.InnerFilter }


func (s *NotFilterContext) SetInnerFilter(v IFilterContext) { s.InnerFilter = v }

func (s *NotFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotFilterContext) NotOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserNotOperator, 0)
}

func (s *NotFilterContext) RxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserRxBracket, 0)
}

func (s *NotFilterContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}



type GroupFilterContext struct {
	*FilterContext
	InnerFilter IFilterContext 
}

func NewGroupFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupFilterContext {
	var p = new(GroupFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *GroupFilterContext) GetInnerFilter() IFilterContext { return s.InnerFilter }


func (s *GroupFilterContext) SetInnerFilter(v IFilterContext) { s.InnerFilter = v }

func (s *GroupFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupFilterContext) LxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserLxBracket, 0)
}

func (s *GroupFilterContext) RxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserRxBracket, 0)
}

func (s *GroupFilterContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}



type AttributeExprFilterContext struct {
	*FilterContext
	AttributeExpr IAttributeExpressionContext 
}

func NewAttributeExprFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeExprFilterContext {
	var p = new(AttributeExprFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *AttributeExprFilterContext) GetAttributeExpr() IAttributeExpressionContext { return s.AttributeExpr }


func (s *AttributeExprFilterContext) SetAttributeExpr(v IAttributeExpressionContext) { s.AttributeExpr = v }

func (s *AttributeExprFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeExprFilterContext) AttributeExpression() IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}



type OrFilterContext struct {
	*FilterContext
	Left IFilterContext 
	Right IFilterContext 
}

func NewOrFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrFilterContext {
	var p = new(OrFilterContext)

	p.FilterContext = NewEmptyFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FilterContext))

	return p
}


func (s *OrFilterContext) GetLeft() IFilterContext { return s.Left }

func (s *OrFilterContext) GetRight() IFilterContext { return s.Right }


func (s *OrFilterContext) SetLeft(v IFilterContext) { s.Left = v }

func (s *OrFilterContext) SetRight(v IFilterContext) { s.Right = v }

func (s *OrFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrFilterContext) OrOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserOrOperator, 0)
}

func (s *OrFilterContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *OrFilterContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}




func (p *FilterParser) Filter() (localctx IFilterContext) {
	return p.filter(0)
}

func (p *FilterParser) filter(_p int) (localctx IFilterContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFilterContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFilterContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FilterParserRULE_filter, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttributeExprFilterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(18)

			var _x = p.AttributeExpression()


			localctx.(*AttributeExprFilterContext).AttributeExpr = _x
		}


	case 2:
		localctx = NewValueExprFilterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)

			var _x = p.ValueExpression()


			localctx.(*ValueExprFilterContext).ValueExpr = _x
		}


	case 3:
		localctx = NewNotFilterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(FilterParserNotOperator)
		}
		{
			p.SetState(21)

			var _x = p.filter(0)

			localctx.(*NotFilterContext).InnerFilter = _x
		}
		{
			p.SetState(22)
			p.Match(FilterParserRxBracket)
		}


	case 4:
		localctx = NewGroupFilterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(FilterParserLxBracket)
		}
		{
			p.SetState(25)

			var _x = p.filter(0)

			localctx.(*GroupFilterContext).InnerFilter = _x
		}
		{
			p.SetState(26)
			p.Match(FilterParserRxBracket)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndFilterContext(p, NewFilterContext(p, _parentctx, _parentState))
				localctx.(*AndFilterContext).Left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, FilterParserRULE_filter)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(31)
					p.Match(FilterParserAndOperator)
				}
				{
					p.SetState(32)

					var _x = p.filter(6)

					localctx.(*AndFilterContext).Right = _x
				}


			case 2:
				localctx = NewOrFilterContext(p, NewFilterContext(p, _parentctx, _parentState))
				localctx.(*OrFilterContext).Left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, FilterParserRULE_filter)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(34)
					p.Match(FilterParserOrOperator)
				}
				{
					p.SetState(35)

					var _x = p.filter(5)

					localctx.(*OrFilterContext).Right = _x
				}

			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}



	return localctx
}


// IAttributeExpressionContext is an interface to support dynamic dispatch.
type IAttributeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the Op token.
	GetOp() antlr.Token 

	// GetValue returns the Value token.
	GetValue() antlr.Token 


	// SetOp sets the Op token.
	SetOp(antlr.Token) 

	// SetValue sets the Value token.
	SetValue(antlr.Token) 


	// GetPath returns the Path rule contexts.
	GetPath() IAttributePathContext


	// SetPath sets the Path rule contexts.
	SetPath(IAttributePathContext)


	// IsAttributeExpressionContext differentiates from other interfaces.
	IsAttributeExpressionContext()
}

type AttributeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Path IAttributePathContext 
	Op antlr.Token
	Value antlr.Token
}

func NewEmptyAttributeExpressionContext() *AttributeExpressionContext {
	var p = new(AttributeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_attributeExpression
	return p
}

func (*AttributeExpressionContext) IsAttributeExpressionContext() {}

func NewAttributeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeExpressionContext {
	var p = new(AttributeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_attributeExpression

	return p
}

func (s *AttributeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeExpressionContext) GetOp() antlr.Token { return s.Op }

func (s *AttributeExpressionContext) GetValue() antlr.Token { return s.Value }


func (s *AttributeExpressionContext) SetOp(v antlr.Token) { s.Op = v }

func (s *AttributeExpressionContext) SetValue(v antlr.Token) { s.Value = v }


func (s *AttributeExpressionContext) GetPath() IAttributePathContext { return s.Path }


func (s *AttributeExpressionContext) SetPath(v IAttributePathContext) { s.Path = v }


func (s *AttributeExpressionContext) AttributePath() IAttributePathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributePathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributePathContext)
}

func (s *AttributeExpressionContext) PrOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserPrOperator, 0)
}

func (s *AttributeExpressionContext) ComparisonValue() antlr.TerminalNode {
	return s.GetToken(FilterParserComparisonValue, 0)
}

func (s *AttributeExpressionContext) EqOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserEqOperator, 0)
}

func (s *AttributeExpressionContext) NeOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserNeOperator, 0)
}

func (s *AttributeExpressionContext) CoOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserCoOperator, 0)
}

func (s *AttributeExpressionContext) SwOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserSwOperator, 0)
}

func (s *AttributeExpressionContext) EwOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserEwOperator, 0)
}

func (s *AttributeExpressionContext) GtOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserGtOperator, 0)
}

func (s *AttributeExpressionContext) LtOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserLtOperator, 0)
}

func (s *AttributeExpressionContext) GeOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserGeOperator, 0)
}

func (s *AttributeExpressionContext) LeOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserLeOperator, 0)
}

func (s *AttributeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *FilterParser) AttributeExpression() (localctx IAttributeExpressionContext) {
	localctx = NewAttributeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterParserRULE_attributeExpression)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)

			var _x = p.AttributePath()


			localctx.(*AttributeExpressionContext).Path = _x
		}
		{
			p.SetState(42)

			var _m = p.Match(FilterParserPrOperator)

			localctx.(*AttributeExpressionContext).Op = _m
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)

			var _x = p.AttributePath()


			localctx.(*AttributeExpressionContext).Path = _x
		}
		p.SetState(45)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AttributeExpressionContext).Op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FilterParserEqOperator) | (1 << FilterParserNeOperator) | (1 << FilterParserCoOperator) | (1 << FilterParserSwOperator) | (1 << FilterParserEwOperator) | (1 << FilterParserGtOperator) | (1 << FilterParserGeOperator) | (1 << FilterParserLtOperator) | (1 << FilterParserLeOperator))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AttributeExpressionContext).Op = _ri
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(46)

			var _m = p.Match(FilterParserComparisonValue)

			localctx.(*AttributeExpressionContext).Value = _m
		}

	}


	return localctx
}


// IAttributePathContext is an interface to support dynamic dispatch.
type IAttributePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPath returns the path attribute.
	GetPath() *attr.Path


	// SetPath sets the path attribute.
	SetPath(*attr.Path)


	// IsAttributePathContext differentiates from other interfaces.
	IsAttributePathContext()
}

type AttributePathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	path *attr.Path
}

func NewEmptyAttributePathContext() *AttributePathContext {
	var p = new(AttributePathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_attributePath
	return p
}

func (*AttributePathContext) IsAttributePathContext() {}

func NewAttributePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributePathContext {
	var p = new(AttributePathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_attributePath

	return p
}

func (s *AttributePathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributePathContext) GetPath() *attr.Path { return s.path }


func (s *AttributePathContext) SetPath(v *attr.Path) { s.path = v }


func (s *AttributePathContext) AllAttributeName() []antlr.TerminalNode {
	return s.GetTokens(FilterParserAttributeName)
}

func (s *AttributePathContext) AttributeName(i int) antlr.TerminalNode {
	return s.GetToken(FilterParserAttributeName, i)
}

func (s *AttributePathContext) Urn() antlr.TerminalNode {
	return s.GetToken(FilterParserUrn, 0)
}

func (s *AttributePathContext) Dot() antlr.TerminalNode {
	return s.GetToken(FilterParserDot, 0)
}

func (s *AttributePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *FilterParser) AttributePath() (localctx IAttributePathContext) {
	localctx = NewAttributePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterParserRULE_attributePath)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == FilterParserUrn {
		{
			p.SetState(50)
			p.Match(FilterParserUrn)
		}

	}
	{
		p.SetState(53)
		p.Match(FilterParserAttributeName)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(54)
			p.Match(FilterParserDot)
		}
		{
			p.SetState(55)
			p.Match(FilterParserAttributeName)
		}


	}
	localctx.(*AttributePathContext).path = attr.Parse(localctx.(*AttributePathContext).GetText())
	p.SetState(59)

	if !(!localctx.(*AttributePathContext).path.Undefined()) {
		panic(antlr.NewFailedPredicateException(p, "!$path.Undefined()", "cannot match any valid attribute path"))
	}



	return localctx
}


// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPath returns the Path rule contexts.
	GetPath() IAttributePathContext

	// GetInnerFilter returns the InnerFilter rule contexts.
	GetInnerFilter() IValueFilterContext


	// SetPath sets the Path rule contexts.
	SetPath(IAttributePathContext)

	// SetInnerFilter sets the InnerFilter rule contexts.
	SetInnerFilter(IValueFilterContext)


	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Path IAttributePathContext 
	InnerFilter IValueFilterContext 
}

func NewEmptyValueExpressionContext() *ValueExpressionContext {
	var p = new(ValueExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_valueExpression
	return p
}

func (*ValueExpressionContext) IsValueExpressionContext() {}

func NewValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_valueExpression

	return p
}

func (s *ValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressionContext) GetPath() IAttributePathContext { return s.Path }

func (s *ValueExpressionContext) GetInnerFilter() IValueFilterContext { return s.InnerFilter }


func (s *ValueExpressionContext) SetPath(v IAttributePathContext) { s.Path = v }

func (s *ValueExpressionContext) SetInnerFilter(v IValueFilterContext) { s.InnerFilter = v }


func (s *ValueExpressionContext) LxSquareBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserLxSquareBracket, 0)
}

func (s *ValueExpressionContext) RxSquareBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserRxSquareBracket, 0)
}

func (s *ValueExpressionContext) AttributePath() IAttributePathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributePathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributePathContext)
}

func (s *ValueExpressionContext) ValueFilter() IValueFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueFilterContext)
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *FilterParser) ValueExpression() (localctx IValueExpressionContext) {
	localctx = NewValueExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterParserRULE_valueExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)

		var _x = p.AttributePath()


		localctx.(*ValueExpressionContext).Path = _x
	}
	{
		p.SetState(62)
		p.Match(FilterParserLxSquareBracket)
	}
	{
		p.SetState(63)

		var _x = p.ValueFilter()


		localctx.(*ValueExpressionContext).InnerFilter = _x
	}
	{
		p.SetState(64)
		p.Match(FilterParserRxSquareBracket)
	}



	return localctx
}


// IValueFilterContext is an interface to support dynamic dispatch.
type IValueFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueFilterContext differentiates from other interfaces.
	IsValueFilterContext()
}

type ValueFilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueFilterContext() *ValueFilterContext {
	var p = new(ValueFilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_valueFilter
	return p
}

func (*ValueFilterContext) IsValueFilterContext() {}

func NewValueFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueFilterContext {
	var p = new(ValueFilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_valueFilter

	return p
}

func (s *ValueFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueFilterContext) CopyFrom(ctx *ValueFilterContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type NotValueFilterContext struct {
	*ValueFilterContext
	InnerFilter IValueFilterContext 
}

func NewNotValueFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotValueFilterContext {
	var p = new(NotValueFilterContext)

	p.ValueFilterContext = NewEmptyValueFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueFilterContext))

	return p
}


func (s *NotValueFilterContext) GetInnerFilter() IValueFilterContext { return s.InnerFilter }


func (s *NotValueFilterContext) SetInnerFilter(v IValueFilterContext) { s.InnerFilter = v }

func (s *NotValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotValueFilterContext) NotOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserNotOperator, 0)
}

func (s *NotValueFilterContext) RxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserRxBracket, 0)
}

func (s *NotValueFilterContext) ValueFilter() IValueFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueFilterContext)
}



type AttributeExprValueFilterContext struct {
	*ValueFilterContext
	AttributeExpr IAttributeExpressionContext 
}

func NewAttributeExprValueFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeExprValueFilterContext {
	var p = new(AttributeExprValueFilterContext)

	p.ValueFilterContext = NewEmptyValueFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueFilterContext))

	return p
}


func (s *AttributeExprValueFilterContext) GetAttributeExpr() IAttributeExpressionContext { return s.AttributeExpr }


func (s *AttributeExprValueFilterContext) SetAttributeExpr(v IAttributeExpressionContext) { s.AttributeExpr = v }

func (s *AttributeExprValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeExprValueFilterContext) AttributeExpression() IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}



type AndValueFilterContext struct {
	*ValueFilterContext
	Left IAttributeExpressionContext 
	Op antlr.Token
	Right IAttributeExpressionContext 
}

func NewAndValueFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndValueFilterContext {
	var p = new(AndValueFilterContext)

	p.ValueFilterContext = NewEmptyValueFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueFilterContext))

	return p
}


func (s *AndValueFilterContext) GetOp() antlr.Token { return s.Op }


func (s *AndValueFilterContext) SetOp(v antlr.Token) { s.Op = v }


func (s *AndValueFilterContext) GetLeft() IAttributeExpressionContext { return s.Left }

func (s *AndValueFilterContext) GetRight() IAttributeExpressionContext { return s.Right }


func (s *AndValueFilterContext) SetLeft(v IAttributeExpressionContext) { s.Left = v }

func (s *AndValueFilterContext) SetRight(v IAttributeExpressionContext) { s.Right = v }

func (s *AndValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndValueFilterContext) AllAttributeExpression() []IAttributeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem())
	var tst = make([]IAttributeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeExpressionContext)
		}
	}

	return tst
}

func (s *AndValueFilterContext) AttributeExpression(i int) IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}

func (s *AndValueFilterContext) AndOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserAndOperator, 0)
}



type OrValueFilterContext struct {
	*ValueFilterContext
	Left IAttributeExpressionContext 
	Op antlr.Token
	Right IAttributeExpressionContext 
}

func NewOrValueFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrValueFilterContext {
	var p = new(OrValueFilterContext)

	p.ValueFilterContext = NewEmptyValueFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueFilterContext))

	return p
}


func (s *OrValueFilterContext) GetOp() antlr.Token { return s.Op }


func (s *OrValueFilterContext) SetOp(v antlr.Token) { s.Op = v }


func (s *OrValueFilterContext) GetLeft() IAttributeExpressionContext { return s.Left }

func (s *OrValueFilterContext) GetRight() IAttributeExpressionContext { return s.Right }


func (s *OrValueFilterContext) SetLeft(v IAttributeExpressionContext) { s.Left = v }

func (s *OrValueFilterContext) SetRight(v IAttributeExpressionContext) { s.Right = v }

func (s *OrValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrValueFilterContext) AllAttributeExpression() []IAttributeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem())
	var tst = make([]IAttributeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeExpressionContext)
		}
	}

	return tst
}

func (s *OrValueFilterContext) AttributeExpression(i int) IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}

func (s *OrValueFilterContext) OrOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserOrOperator, 0)
}



type GroupValueFilterContext struct {
	*ValueFilterContext
	InnerFilter IValueFilterContext 
}

func NewGroupValueFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupValueFilterContext {
	var p = new(GroupValueFilterContext)

	p.ValueFilterContext = NewEmptyValueFilterContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueFilterContext))

	return p
}


func (s *GroupValueFilterContext) GetInnerFilter() IValueFilterContext { return s.InnerFilter }


func (s *GroupValueFilterContext) SetInnerFilter(v IValueFilterContext) { s.InnerFilter = v }

func (s *GroupValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupValueFilterContext) LxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserLxBracket, 0)
}

func (s *GroupValueFilterContext) RxBracket() antlr.TerminalNode {
	return s.GetToken(FilterParserRxBracket, 0)
}

func (s *GroupValueFilterContext) ValueFilter() IValueFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueFilterContext)
}




func (p *FilterParser) ValueFilter() (localctx IValueFilterContext) {
	localctx = NewValueFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterParserRULE_valueFilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttributeExprValueFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)

			var _x = p.AttributeExpression()


			localctx.(*AttributeExprValueFilterContext).AttributeExpr = _x
		}


	case 2:
		localctx = NewAndValueFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)

			var _x = p.AttributeExpression()


			localctx.(*AndValueFilterContext).Left = _x
		}
		{
			p.SetState(68)

			var _m = p.Match(FilterParserAndOperator)

			localctx.(*AndValueFilterContext).Op = _m
		}
		{
			p.SetState(69)

			var _x = p.AttributeExpression()


			localctx.(*AndValueFilterContext).Right = _x
		}


	case 3:
		localctx = NewOrValueFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)

			var _x = p.AttributeExpression()


			localctx.(*OrValueFilterContext).Left = _x
		}
		{
			p.SetState(72)

			var _m = p.Match(FilterParserOrOperator)

			localctx.(*OrValueFilterContext).Op = _m
		}
		{
			p.SetState(73)

			var _x = p.AttributeExpression()


			localctx.(*OrValueFilterContext).Right = _x
		}


	case 4:
		localctx = NewNotValueFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(FilterParserNotOperator)
		}
		{
			p.SetState(76)

			var _x = p.ValueFilter()


			localctx.(*NotValueFilterContext).InnerFilter = _x
		}
		{
			p.SetState(77)
			p.Match(FilterParserRxBracket)
		}


	case 5:
		localctx = NewGroupValueFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(79)
			p.Match(FilterParserLxBracket)
		}
		{
			p.SetState(80)

			var _x = p.ValueFilter()


			localctx.(*GroupValueFilterContext).InnerFilter = _x
		}
		{
			p.SetState(81)
			p.Match(FilterParserRxBracket)
		}

	}


	return localctx
}


func (p *FilterParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *FilterContext = nil
			if localctx != nil { t = localctx.(*FilterContext) }
			return p.Filter_Sempred(t, predIndex)

	case 3:
			var t *AttributePathContext = nil
			if localctx != nil { t = localctx.(*AttributePathContext) }
			return p.AttributePath_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FilterParser) Filter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FilterParser) AttributePath_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return !localctx.(*AttributePathContext).path.Undefined()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

