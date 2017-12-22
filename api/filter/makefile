SHELL = /bin/bash

.PHONY = clean

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

filter_parser.go FilterParser.tokens: FilterParser.g4 filter_lexer.go FilterLexer.tokens antlr
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"
	@sed -i -e 1,2d filter_parser.go

filter_lexer.go FilterLexer.tokens: FilterLexer.g4 antlr
	$(SHELL) -c "./antlr -Dlanguage=Go -package $(CURRENT_DIR) -no-listener -no-visitor $<"
	@sed -i -e 1,2d filter_lexer.go

antlr:
	@docker pull leodido/antlr
	@docker create --name antlr leodido/antlr
	@docker cp antlr:antlr .
	@docker rm antlr

test: filter_parser.go filter_lexer.go FilterParser.tokens FilterLexer.tokens *_test.go
	go test -v

clean:
	rm -f filter_{parser,lexer}.go Filter{Parser,Lexer}.tokens
