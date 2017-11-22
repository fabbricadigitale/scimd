// Generated from /home/leodido/workspaces/go/src/github.com/fabbricadigitale/scimd/api/filter/Filter.g4 by ANTLR 4.6.

package filter // Filter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFilterListener is a complete listener for a parse tree produced by FilterParser.
type BaseFilterListener struct{}

var _ FilterListener = &BaseFilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseFilterListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseFilterListener) ExitRoot(ctx *RootContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseFilterListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseFilterListener) ExitFilter(ctx *FilterContext) {}

// EnterAttributeExpression is called when production attributeExpression is entered.
func (s *BaseFilterListener) EnterAttributeExpression(ctx *AttributeExpressionContext) {}

// ExitAttributeExpression is called when production attributeExpression is exited.
func (s *BaseFilterListener) ExitAttributeExpression(ctx *AttributeExpressionContext) {}

// EnterAttributePath is called when production attributePath is entered.
func (s *BaseFilterListener) EnterAttributePath(ctx *AttributePathContext) {}

// ExitAttributePath is called when production attributePath is exited.
func (s *BaseFilterListener) ExitAttributePath(ctx *AttributePathContext) {}

// EnterValueExpression is called when production valueExpression is entered.
func (s *BaseFilterListener) EnterValueExpression(ctx *ValueExpressionContext) {}

// ExitValueExpression is called when production valueExpression is exited.
func (s *BaseFilterListener) ExitValueExpression(ctx *ValueExpressionContext) {}

// EnterValueFilter is called when production valueFilter is entered.
func (s *BaseFilterListener) EnterValueFilter(ctx *ValueFilterContext) {}

// ExitValueFilter is called when production valueFilter is exited.
func (s *BaseFilterListener) ExitValueFilter(ctx *ValueFilterContext) {}
