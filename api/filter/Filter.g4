/**
 * This grammar is ...
 */
grammar Filter;

@parser::header { 
}

@parser::members {
var debug = false

func log(format string, a ...interface{}) {
    if debug {
        fmt.Printf(format + "\n", a...)
    }
}
}

root
        : filter? EOF
        ;

filter
        : attributeExpression
        | LxFilter=filter LogicalOperator=AndOperator RxFilter=filter { log("%s => (%s) AND (%s)", $text, $LxFilter.text, $RxFilter.text); }
        | LxFilter=filter LogicalOperator=OrOperator RxFilter=filter { log("%s => (%s) OR (%s)", $text, $LxFilter.text, $RxFilter.text); }
        | valueExpression
        | NotOperator filter ')'
        | '(' filter ')'
        ;

attributeExpression
        : attributePath PrOperator
        | attributePath ComparisonOperator=(EqOperator|NeOperator|CoOperator|SwOperator|EwOperator|GtOperator|LtOperator|GeOperator|LeOperator) ComparisonValue
        ;

attributePath
        : (Urn ':')? AttributeName ('.' SubAttributeName=AttributeName)?
        ;

valueExpression
        : attributePath '[' valueFilter ']'
        ;

valueFilter
        : attributeExpression
        | LxAttributeExpression=attributeExpression LogicalOperator=AndOperator RxAttributeExpression=attributeExpression { log("%s => (%s) AND (%s)", $text, $LxAttributeExpression.text, $RxAttributeExpression.text); }
        | LxAttributeExpression=attributeExpression LogicalOperator=OrOperator RxAttributeExpression=attributeExpression { log("%s => (%s) AND (%s)", $text, $LxAttributeExpression.text, $RxAttributeExpression.text); }
        | NotOperator valueFilter ')'
        | '(' valueFilter ')'
        ;

/*
 * LEXER
 */
AttributeName
        : Alpha Char*
        ;

Urn
        : 'urn'
        ;

ComparisonValue
        : ' ' STRING
        | ' ' NUMBER
        | ' true'
        | ' false'
        | ' null'
        ;

Space
        : ' ' 
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
        : 'not ('
        | 'not('
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
