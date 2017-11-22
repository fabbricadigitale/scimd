// Generated from /home/leodido/workspaces/go/src/github.com/fabbricadigitale/scimd/api/filter/Filter.g4 by ANTLR 4.6.

package filter // Filter
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)


 


// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 25, 92, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3, 
	2, 5, 2, 16, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46, 11, 
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 55, 10, 4, 3, 5, 3, 
	5, 5, 5, 59, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5, 3, 6, 3, 6, 3, 6, 
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 90, 10, 7, 
	3, 7, 2, 3, 4, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 2, 14, 22, 98, 2, 15, 3, 
	2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 54, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 65, 
	3, 2, 2, 2, 12, 89, 3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 
	15, 16, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 18, 7, 2, 2, 3, 18, 3, 3, 2, 
	2, 2, 19, 20, 8, 3, 1, 2, 20, 31, 5, 6, 4, 2, 21, 31, 5, 10, 6, 2, 22, 
	23, 7, 25, 2, 2, 23, 24, 5, 4, 3, 2, 24, 25, 7, 3, 2, 2, 25, 31, 3, 2, 
	2, 2, 26, 27, 7, 4, 2, 2, 27, 28, 5, 4, 3, 2, 28, 29, 7, 3, 2, 2, 29, 31, 
	3, 2, 2, 2, 30, 19, 3, 2, 2, 2, 30, 21, 3, 2, 2, 2, 30, 22, 3, 2, 2, 2, 
	30, 26, 3, 2, 2, 2, 31, 44, 3, 2, 2, 2, 32, 33, 12, 7, 2, 2, 33, 34, 7, 
	23, 2, 2, 34, 35, 5, 4, 3, 8, 35, 36, 8, 3, 1, 2, 36, 43, 3, 2, 2, 2, 37, 
	38, 12, 6, 2, 2, 38, 39, 7, 24, 2, 2, 39, 40, 5, 4, 3, 7, 40, 41, 8, 3, 
	1, 2, 41, 43, 3, 2, 2, 2, 42, 32, 3, 2, 2, 2, 42, 37, 3, 2, 2, 2, 43, 46, 
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 5, 3, 2, 2, 2, 
	46, 44, 3, 2, 2, 2, 47, 48, 5, 8, 5, 2, 48, 49, 7, 13, 2, 2, 49, 55, 3, 
	2, 2, 2, 50, 51, 5, 8, 5, 2, 51, 52, 9, 2, 2, 2, 52, 53, 7, 11, 2, 2, 53, 
	55, 3, 2, 2, 2, 54, 47, 3, 2, 2, 2, 54, 50, 3, 2, 2, 2, 55, 7, 3, 2, 2, 
	2, 56, 57, 7, 10, 2, 2, 57, 59, 7, 5, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 63, 7, 9, 2, 2, 61, 62, 7, 6, 2, 2, 
	62, 64, 7, 9, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 9, 3, 2, 
	2, 2, 65, 66, 5, 8, 5, 2, 66, 67, 7, 7, 2, 2, 67, 68, 5, 12, 7, 2, 68, 
	69, 7, 8, 2, 2, 69, 11, 3, 2, 2, 2, 70, 90, 5, 6, 4, 2, 71, 72, 5, 6, 4, 
	2, 72, 73, 7, 23, 2, 2, 73, 74, 5, 6, 4, 2, 74, 75, 8, 7, 1, 2, 75, 90, 
	3, 2, 2, 2, 76, 77, 5, 6, 4, 2, 77, 78, 7, 24, 2, 2, 78, 79, 5, 6, 4, 2, 
	79, 80, 8, 7, 1, 2, 80, 90, 3, 2, 2, 2, 81, 82, 7, 25, 2, 2, 82, 83, 5, 
	12, 7, 2, 83, 84, 7, 3, 2, 2, 84, 90, 3, 2, 2, 2, 85, 86, 7, 4, 2, 2, 86, 
	87, 5, 12, 7, 2, 87, 88, 7, 3, 2, 2, 88, 90, 3, 2, 2, 2, 89, 70, 3, 2, 
	2, 2, 89, 71, 3, 2, 2, 2, 89, 76, 3, 2, 2, 2, 89, 81, 3, 2, 2, 2, 89, 85, 
	3, 2, 2, 2, 90, 13, 3, 2, 2, 2, 10, 15, 30, 42, 44, 54, 58, 63, 89,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "')'", "'('", "':'", "'.'", "'['", "']'", "", "'urn'", "", "' '", "' pr'", 
	"' eq'", "' ne'", "' co'", "' sw'", "' ew'", "' gt'", "' ge'", "' lt'", 
	"' le'", "' and '", "' or '",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "AttributeName", "Urn", "ComparisonValue", 
	"Space", "PrOperator", "EqOperator", "NeOperator", "CoOperator", "SwOperator", 
	"EwOperator", "GtOperator", "GeOperator", "LtOperator", "LeOperator", "AndOperator", 
	"OrOperator", "NotOperator",
}

var ruleNames = []string{
	"root", "filter", "attributeExpression", "attributePath", "valueExpression", 
	"valueFilter",
}

type FilterParser struct {
	*antlr.BaseParser
}

func NewFilterParser(input antlr.TokenStream) *FilterParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(FilterParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Filter.g4"

	return this
}


var debug = false

func log(format string, a ...interface{}) {
    if debug {
        fmt.Printf(format + "\n", a...)
    }
}



// FilterParser tokens.
const (
	FilterParserEOF = antlr.TokenEOF
	FilterParserT__0 = 1
	FilterParserT__1 = 2
	FilterParserT__2 = 3
	FilterParserT__3 = 4
	FilterParserT__4 = 5
	FilterParserT__5 = 6
	FilterParserAttributeName = 7
	FilterParserUrn = 8
	FilterParserComparisonValue = 9
	FilterParserSpace = 10
	FilterParserPrOperator = 11
	FilterParserEqOperator = 12
	FilterParserNeOperator = 13
	FilterParserCoOperator = 14
	FilterParserSwOperator = 15
	FilterParserEwOperator = 16
	FilterParserGtOperator = 17
	FilterParserGeOperator = 18
	FilterParserLtOperator = 19
	FilterParserLeOperator = 20
	FilterParserAndOperator = 21
	FilterParserOrOperator = 22
	FilterParserNotOperator = 23
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


func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitRoot(s)
	}
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


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FilterParserT__1) | (1 << FilterParserAttributeName) | (1 << FilterParserUrn) | (1 << FilterParserNotOperator))) != 0) {
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

	// GetLogicalOperator returns the LogicalOperator token.
	GetLogicalOperator() antlr.Token 


	// SetLogicalOperator sets the LogicalOperator token.
	SetLogicalOperator(antlr.Token) 


	// GetLxFilter returns the LxFilter rule contexts.
	GetLxFilter() IFilterContext

	// GetRxFilter returns the RxFilter rule contexts.
	GetRxFilter() IFilterContext


	// SetLxFilter sets the LxFilter rule contexts.
	SetLxFilter(IFilterContext)

	// SetRxFilter sets the RxFilter rule contexts.
	SetRxFilter(IFilterContext)


	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	LxFilter IFilterContext 
	LogicalOperator antlr.Token
	RxFilter IFilterContext 
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

func (s *FilterContext) GetLogicalOperator() antlr.Token { return s.LogicalOperator }


func (s *FilterContext) SetLogicalOperator(v antlr.Token) { s.LogicalOperator = v }


func (s *FilterContext) GetLxFilter() IFilterContext { return s.LxFilter }

func (s *FilterContext) GetRxFilter() IFilterContext { return s.RxFilter }


func (s *FilterContext) SetLxFilter(v IFilterContext) { s.LxFilter = v }

func (s *FilterContext) SetRxFilter(v IFilterContext) { s.RxFilter = v }


func (s *FilterContext) AttributeExpression() IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}

func (s *FilterContext) ValueExpression() IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *FilterContext) NotOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserNotOperator, 0)
}

func (s *FilterContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *FilterContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *FilterContext) AndOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserAndOperator, 0)
}

func (s *FilterContext) OrOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserOrOperator, 0)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitFilter(s)
	}
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
		{
			p.SetState(18)
			p.AttributeExpression()
		}


	case 2:
		{
			p.SetState(19)
			p.ValueExpression()
		}


	case 3:
		{
			p.SetState(20)
			p.Match(FilterParserNotOperator)
		}
		{
			p.SetState(21)
			p.filter(0)
		}
		{
			p.SetState(22)
			p.Match(FilterParserT__0)
		}


	case 4:
		{
			p.SetState(24)
			p.Match(FilterParserT__1)
		}
		{
			p.SetState(25)
			p.filter(0)
		}
		{
			p.SetState(26)
			p.Match(FilterParserT__0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFilterContext(p, _parentctx, _parentState)
				localctx.(*FilterContext).LxFilter = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FilterParserRULE_filter)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(31)

					var _m = p.Match(FilterParserAndOperator)

					localctx.(*FilterContext).LogicalOperator = _m
				}
				{
					p.SetState(32)

					var _x = p.filter(6)

					localctx.(*FilterContext).RxFilter = _x
				}
				 log("%s => (%s) AND (%s)", p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1)), (func() string { if localctx.(*FilterContext).GetLxFilter() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*FilterContext).GetLxFilter().GetStart(), localctx.(*FilterContext).LxFilter.GetStop()) }}()), (func() string { if localctx.(*FilterContext).GetRxFilter() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*FilterContext).GetRxFilter().GetStart(), localctx.(*FilterContext).RxFilter.GetStop()) }}())); 


			case 2:
				localctx = NewFilterContext(p, _parentctx, _parentState)
				localctx.(*FilterContext).LxFilter = _prevctx
				p.PushNewRecursionContext(localctx, _startState, FilterParserRULE_filter)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(36)

					var _m = p.Match(FilterParserOrOperator)

					localctx.(*FilterContext).LogicalOperator = _m
				}
				{
					p.SetState(37)

					var _x = p.filter(5)

					localctx.(*FilterContext).RxFilter = _x
				}
				 log("%s => (%s) OR (%s)", p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1)), (func() string { if localctx.(*FilterContext).GetLxFilter() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*FilterContext).GetLxFilter().GetStart(), localctx.(*FilterContext).LxFilter.GetStop()) }}()), (func() string { if localctx.(*FilterContext).GetRxFilter() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*FilterContext).GetRxFilter().GetStart(), localctx.(*FilterContext).RxFilter.GetStop()) }}())); 

			}

		}
		p.SetState(44)
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

	// GetComparisonOperator returns the ComparisonOperator token.
	GetComparisonOperator() antlr.Token 


	// SetComparisonOperator sets the ComparisonOperator token.
	SetComparisonOperator(antlr.Token) 


	// IsAttributeExpressionContext differentiates from other interfaces.
	IsAttributeExpressionContext()
}

type AttributeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ComparisonOperator antlr.Token
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

func (s *AttributeExpressionContext) GetComparisonOperator() antlr.Token { return s.ComparisonOperator }


func (s *AttributeExpressionContext) SetComparisonOperator(v antlr.Token) { s.ComparisonOperator = v }


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


func (s *AttributeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterAttributeExpression(s)
	}
}

func (s *AttributeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitAttributeExpression(s)
	}
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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.AttributePath()
		}
		{
			p.SetState(46)
			p.Match(FilterParserPrOperator)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.AttributePath()
		}
		p.SetState(49)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AttributeExpressionContext).ComparisonOperator = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << FilterParserEqOperator) | (1 << FilterParserNeOperator) | (1 << FilterParserCoOperator) | (1 << FilterParserSwOperator) | (1 << FilterParserEwOperator) | (1 << FilterParserGtOperator) | (1 << FilterParserGeOperator) | (1 << FilterParserLtOperator) | (1 << FilterParserLeOperator))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AttributeExpressionContext).ComparisonOperator = _ri
		} else {
		    p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(50)
			p.Match(FilterParserComparisonValue)
		}

	}


	return localctx
}


// IAttributePathContext is an interface to support dynamic dispatch.
type IAttributePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSubAttributeName returns the SubAttributeName token.
	GetSubAttributeName() antlr.Token 


	// SetSubAttributeName sets the SubAttributeName token.
	SetSubAttributeName(antlr.Token) 


	// IsAttributePathContext differentiates from other interfaces.
	IsAttributePathContext()
}

type AttributePathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	SubAttributeName antlr.Token
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

func (s *AttributePathContext) GetSubAttributeName() antlr.Token { return s.SubAttributeName }


func (s *AttributePathContext) SetSubAttributeName(v antlr.Token) { s.SubAttributeName = v }


func (s *AttributePathContext) AllAttributeName() []antlr.TerminalNode {
	return s.GetTokens(FilterParserAttributeName)
}

func (s *AttributePathContext) AttributeName(i int) antlr.TerminalNode {
	return s.GetToken(FilterParserAttributeName, i)
}

func (s *AttributePathContext) Urn() antlr.TerminalNode {
	return s.GetToken(FilterParserUrn, 0)
}

func (s *AttributePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterAttributePath(s)
	}
}

func (s *AttributePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitAttributePath(s)
	}
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == FilterParserUrn {
		{
			p.SetState(54)
			p.Match(FilterParserUrn)
		}
		{
			p.SetState(55)
			p.Match(FilterParserT__2)
		}

	}
	{
		p.SetState(58)
		p.Match(FilterParserAttributeName)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == FilterParserT__3 {
		{
			p.SetState(59)
			p.Match(FilterParserT__3)
		}
		{
			p.SetState(60)

			var _m = p.Match(FilterParserAttributeName)

			localctx.(*AttributePathContext).SubAttributeName = _m
		}

	}



	return localctx
}


// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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


func (s *ValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterValueExpression(s)
	}
}

func (s *ValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitValueExpression(s)
	}
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
		p.SetState(63)
		p.AttributePath()
	}
	{
		p.SetState(64)
		p.Match(FilterParserT__4)
	}
	{
		p.SetState(65)
		p.ValueFilter()
	}
	{
		p.SetState(66)
		p.Match(FilterParserT__5)
	}



	return localctx
}


// IValueFilterContext is an interface to support dynamic dispatch.
type IValueFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLogicalOperator returns the LogicalOperator token.
	GetLogicalOperator() antlr.Token 


	// SetLogicalOperator sets the LogicalOperator token.
	SetLogicalOperator(antlr.Token) 


	// GetLxAttributeExpression returns the LxAttributeExpression rule contexts.
	GetLxAttributeExpression() IAttributeExpressionContext

	// GetRxAttributeExpression returns the RxAttributeExpression rule contexts.
	GetRxAttributeExpression() IAttributeExpressionContext


	// SetLxAttributeExpression sets the LxAttributeExpression rule contexts.
	SetLxAttributeExpression(IAttributeExpressionContext)

	// SetRxAttributeExpression sets the RxAttributeExpression rule contexts.
	SetRxAttributeExpression(IAttributeExpressionContext)


	// IsValueFilterContext differentiates from other interfaces.
	IsValueFilterContext()
}

type ValueFilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	LxAttributeExpression IAttributeExpressionContext 
	LogicalOperator antlr.Token
	RxAttributeExpression IAttributeExpressionContext 
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

func (s *ValueFilterContext) GetLogicalOperator() antlr.Token { return s.LogicalOperator }


func (s *ValueFilterContext) SetLogicalOperator(v antlr.Token) { s.LogicalOperator = v }


func (s *ValueFilterContext) GetLxAttributeExpression() IAttributeExpressionContext { return s.LxAttributeExpression }

func (s *ValueFilterContext) GetRxAttributeExpression() IAttributeExpressionContext { return s.RxAttributeExpression }


func (s *ValueFilterContext) SetLxAttributeExpression(v IAttributeExpressionContext) { s.LxAttributeExpression = v }

func (s *ValueFilterContext) SetRxAttributeExpression(v IAttributeExpressionContext) { s.RxAttributeExpression = v }


func (s *ValueFilterContext) AllAttributeExpression() []IAttributeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem())
	var tst = make([]IAttributeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeExpressionContext)
		}
	}

	return tst
}

func (s *ValueFilterContext) AttributeExpression(i int) IAttributeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeExpressionContext)
}

func (s *ValueFilterContext) AndOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserAndOperator, 0)
}

func (s *ValueFilterContext) OrOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserOrOperator, 0)
}

func (s *ValueFilterContext) NotOperator() antlr.TerminalNode {
	return s.GetToken(FilterParserNotOperator, 0)
}

func (s *ValueFilterContext) ValueFilter() IValueFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueFilterContext)
}

func (s *ValueFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterValueFilter(s)
	}
}

func (s *ValueFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitValueFilter(s)
	}
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.AttributeExpression()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)

			var _x = p.AttributeExpression()


			localctx.(*ValueFilterContext).LxAttributeExpression = _x
		}
		{
			p.SetState(70)

			var _m = p.Match(FilterParserAndOperator)

			localctx.(*ValueFilterContext).LogicalOperator = _m
		}
		{
			p.SetState(71)

			var _x = p.AttributeExpression()


			localctx.(*ValueFilterContext).RxAttributeExpression = _x
		}
		 log("%s => (%s) AND (%s)", p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1)), (func() string { if localctx.(*ValueFilterContext).GetLxAttributeExpression() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*ValueFilterContext).GetLxAttributeExpression().GetStart(), localctx.(*ValueFilterContext).LxAttributeExpression.GetStop()) }}()), (func() string { if localctx.(*ValueFilterContext).GetRxAttributeExpression() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*ValueFilterContext).GetRxAttributeExpression().GetStart(), localctx.(*ValueFilterContext).RxAttributeExpression.GetStop()) }}())); 


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)

			var _x = p.AttributeExpression()


			localctx.(*ValueFilterContext).LxAttributeExpression = _x
		}
		{
			p.SetState(75)

			var _m = p.Match(FilterParserOrOperator)

			localctx.(*ValueFilterContext).LogicalOperator = _m
		}
		{
			p.SetState(76)

			var _x = p.AttributeExpression()


			localctx.(*ValueFilterContext).RxAttributeExpression = _x
		}
		 log("%s => (%s) AND (%s)", p.GetTokenStream().GetTextFromTokens(localctx.GetStart(), p.GetTokenStream().LT(-1)), (func() string { if localctx.(*ValueFilterContext).GetLxAttributeExpression() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*ValueFilterContext).GetLxAttributeExpression().GetStart(), localctx.(*ValueFilterContext).LxAttributeExpression.GetStop()) }}()), (func() string { if localctx.(*ValueFilterContext).GetRxAttributeExpression() == nil { return "" } else { return p.GetTokenStream().GetTextFromTokens(localctx.(*ValueFilterContext).GetRxAttributeExpression().GetStart(), localctx.(*ValueFilterContext).RxAttributeExpression.GetStop()) }}())); 


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Match(FilterParserNotOperator)
		}
		{
			p.SetState(80)
			p.ValueFilter()
		}
		{
			p.SetState(81)
			p.Match(FilterParserT__0)
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)
			p.Match(FilterParserT__1)
		}
		{
			p.SetState(84)
			p.ValueFilter()
		}
		{
			p.SetState(85)
			p.Match(FilterParserT__0)
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

