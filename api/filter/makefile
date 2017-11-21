SHELL = /bin/bash

.PHONY = clean

filter_{parser,lexer,base_listener,listener}.go Filter{,Lexer}.tokens: Filter.g4
	$(SHELL) -c "./antlr.sh -Dlanguage=Go $?"

clean:
	rm -f filter_{parser,lexer,base_listener,listener}.go Filter{,Lexer}.tokens