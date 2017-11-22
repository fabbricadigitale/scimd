// Generated from /home/leodido/workspaces/go/src/github.com/fabbricadigitale/scimd/api/filter/Filter.g4 by ANTLR 4.6.

package filter // Filter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FilterListener is a complete listener for a parse tree produced by FilterParser.
type FilterListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterAttributeExpression is called when entering the attributeExpression production.
	EnterAttributeExpression(c *AttributeExpressionContext)

	// EnterAttributePath is called when entering the attributePath production.
	EnterAttributePath(c *AttributePathContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterValueFilter is called when entering the valueFilter production.
	EnterValueFilter(c *ValueFilterContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitAttributeExpression is called when exiting the attributeExpression production.
	ExitAttributeExpression(c *AttributeExpressionContext)

	// ExitAttributePath is called when exiting the attributePath production.
	ExitAttributePath(c *AttributePathContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitValueFilter is called when exiting the valueFilter production.
	ExitValueFilter(c *ValueFilterContext)
}
