package filter

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// TreeShapeListener is ...
type TreeShapeListener struct {
	*BaseFilterListener
}

// NewTreeShapeListener is ...
func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// EnterEveryRule is ...
func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Printf("text: %s\n", ctx.GetText())
}

func (s *TreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("%s => \x1b[31;1merror\x1b[0m\n", node.GetSymbol())
}
