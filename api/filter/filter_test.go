package filter

import (
	"bufio"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/stretchr/testify/require"
)

type testErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func (l *testErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.t.Fail()
}

func TestParser(t *testing.T) {
	inFile, _ := os.Open("testdata/ok.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	l := new(testErrorListener)
	l.t = t

	for scanner.Scan() {
		input := scanner.Text()

		stream := antlr.NewInputStream(input)
		lexer := NewFilterLexer(stream)
		lexer.AddErrorListener(l)
		tokens := antlr.NewCommonTokenStream(lexer, 0)

		parser := NewFilterParser(tokens)
		parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		parser.AddErrorListener(l)

		tree := parser.Root()
		/*
			fmt.Println(tree)

			symbols := lexer.GetSymbolicNames()


			for _, tkn := range tokens.GetAllTokens() {
				sym := "//"
				if t := tkn.GetTokenType(); t >= 0 {
					sym = symbols[t]
				}
				fmt.Printf("%+v \t\t => %s\n", tkn, sym)
			}
		*/

		require.Equal(t, 2, tree.GetChildCount()) // filter <EOF>
	}
}

func TestParserError(t *testing.T) {
	inFile, _ := os.Open("testdata/wrong.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		errListener := new(parserErrorListener)
		input := scanner.Text()

		stream := antlr.NewInputStream(input)
		lexer := NewFilterLexer(stream)
		lexer.RemoveErrorListeners()
		tokens := antlr.NewCommonTokenStream(lexer, 0)

		parser := NewFilterParser(tokens)
		parser.RemoveErrorListeners()
		parser.AddErrorListener(errListener)

		require.Panics(t, func() {
			parser.Root()
		})
	}
}
