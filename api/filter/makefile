SHELL = /bin/bash

.PHONY = clean

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MKFILE_PATH))))

filter_{parser,lexer,base_listener,listener}.go Filter{,Lexer}.tokens: Filter.g4
	$(SHELL) -c "./antlr.sh -Dlanguage=Go -package ${CURRENT_DIR} $?"

clean:
	rm -f filter_{parser,lexer,base_listener,listener}.go Filter{,Lexer}.tokens