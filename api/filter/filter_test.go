package filter

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestX(t *testing.T) {
	inFile, _ := os.Open("testdata/ok.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()

		stream := antlr.NewInputStream(input)
		lexer := NewFilterLexer(stream)
		tokens := antlr.NewCommonTokenStream(lexer, 0)

		parser := NewFilterParser(tokens)
		parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		parser.BuildParseTrees = true

		tree := parser.Root()

		symbols := lexer.GetSymbolicNames()

		for _, tkn := range tokens.GetAllTokens() {
			sym := "//"
			if t := tkn.GetTokenType(); t >= 0 {
				sym = symbols[t]
			}
			fmt.Printf("%+v \t\t => %s\n", tkn, sym)
		}

		fmt.Println()

		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	}
}
