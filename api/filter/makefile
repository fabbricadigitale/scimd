SHELL = /bin/bash

.PHONY = clean

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

filter_parser.go FilterParser.tokens: FilterParser.g4 filter_lexer.go FilterLexer.tokens
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"
	@$(SHELL) -c "tail -q -n +3 filter_parser.go | tee filter_parser.go > /dev/null"

filter_lexer.go FilterLexer.tokens: FilterLexer.g4
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"
	@$(SHELL) -c "tail -q -n +3 filter_lexer.go | tee filter_lexer.go > /dev/null"

test: filter_{parser,lexer}.go Filter{Parser,Lexer}.tokens *_test.go
	go test

clean:
	rm -f filter_{parser,lexer}.go Filter{Parser,Lexer}.tokens
