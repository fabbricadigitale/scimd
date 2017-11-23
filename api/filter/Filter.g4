/**
 * This grammar is ...
 */
grammar Filter;

@parser::header { 
}

@parser::members {

}

root
        : filter? EOF
        ;

filter
        : AttributeExpr=attributeExpression # AttributeExprFilter
        | Left=filter AndOperator Right=filter # AndFilter
        | Left=filter OrOperator Right=filter # OrFilter
        | ValueExpr=valueExpression # ValueExprFilter
        | NotOperator InnerFilter=filter ')' # NotFilter
        | '(' InnerFilter=filter ')' # GroupFilter
        ;

attributeExpression
        : Path=attributePath Op=PrOperator
        | Path=attributePath Op=(EqOperator|NeOperator|CoOperator|SwOperator|EwOperator|GtOperator|LtOperator|GeOperator|LeOperator) Value=ComparisonValue
        ;

attributePath
        : (URI=Urn ':')? Name=AttributeName ('.' Sub=AttributeName)?
        ;

valueExpression
        : Path=attributePath '[' InnerFilter=valueFilter ']'
        ;

valueFilter
        : AttributeExpr=attributeExpression # AttributeExprValueFilter
        | Left=attributeExpression Op=AndOperator Right=attributeExpression # AndValueFilter
        | Left=attributeExpression Op=OrOperator Right=attributeExpression # OrValueFilter
        | NotOperator InnerFilter=valueFilter ')' # NotValueFilter
        | '(' InnerFilter=valueFilter ')' # GroupValueFilter
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
