SHELL = /bin/bash

.PHONY = clean

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

filter_{parser,lexer,base_listener,listener,base_visitor,visitor}.go: Filter.g4
	$(SHELL) -c "./antlr.sh -Dlanguage=Go -package $(CURRENT_DIR) -listener -no-visitor $?"

test: filter_{parser,lexer,base_listener,listener,base_visitor,visitor}.go *_test.go
	go test

clean:
	rm -f filter_{parser,lexer,base_listener,listener,base_visitor,visitor}.go Filter{,Lexer}.tokens