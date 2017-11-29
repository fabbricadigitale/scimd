lexer grammar FilterLexer;

AttributeName
        : Alpha Char*
        ;

ComparisonValue
        : ' ' STRING
        | ' ' NUMBER
        | ' true'
        | ' false'
        | ' null'
        ;

Space
        : [ ]
        ;

Urn
        : [uU][rR][nN] Colon (. | ~[ ])+ Colon
        ;

Colon
        : ':'
        ;

PrOperator
        : ' pr'
        ;

EqOperator
        : ' eq'
        ;

NeOperator
        : ' ne'
        ;

CoOperator
        : ' co'
        ;

SwOperator
        : ' sw'
        ;

EwOperator
        : ' ew'
        ;

GtOperator
        : ' gt'
        ;

GeOperator
        : ' ge'
        ;

LtOperator
        : ' lt'
        ;

LeOperator
        : ' le'
        ;

AndOperator
        : ' and '
        ;

OrOperator
        : ' or '
        ;

NotOperator
        : 'not' Space? LxBracket
        ;

RxBracket
        : ')'
        ;

LxBracket
        : '('
        ;

RxSquareBracket
        : ']'
        ;

Dot
        : '.'
        ;

LxSquareBracket
        : '['
        ;

fragment Alpha
        : 'a'..'z'
        | 'A'..'Z'
        ;

fragment Char
        : '-' | '_' | '0'..'9' | Alpha
        ;

/*
 * Part of the JSON grammar.
 * Source: "The Definitive ANTLR 4 Reference", Terence Parr
 */
fragment STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;


fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;


fragment HEX
   : [0-9a-fA-F]
   ;


fragment NUMBER
   : '-'? INT ('.' [0-9] +)? EXP?
   ;


fragment INT
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   : [Ee] [+\-]? INT
   ;