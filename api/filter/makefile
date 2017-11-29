SHELL = /bin/bash

.PHONY = clean

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

filter_{parser,base_listener,listener,base_visitor,visitor}.go: FilterParser.g4 filter_lexer.go FilterLexer.tokens
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"

filter_lexer.go FilterLexer.tokens: FilterLexer.g4
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"

test: filter_{parser,base_listener,listener,base_visitor,visitor}.go filter_lexer.go FilterLexer.tokens *_test.go
	go test

clean:
	rm -f filter_{parser,lexer,base_listener,listener,base_visitor,visitor}.go Filter{Parser,Lexer}.tokens