/**
 * This grammar is ...
 */
parser grammar FilterParser;

@header {
import "github.com/fabbricadigitale/scimd/api/attr"
}

options {
        tokenVocab=FilterLexer;
}

root
        : filter? EOF
        ;

filter
        : AttributeExpr=attributeExpression # AttributeExprFilter
        | Left=filter AndOperator Right=filter # AndFilter
        | Left=filter OrOperator Right=filter # OrFilter
        | ValueExpr=valueExpression # ValueExprFilter
        | NotOperator InnerFilter=filter RxBracket # NotFilter
        | LxBracket InnerFilter=filter RxBracket # GroupFilter
        ;

attributeExpression
        : Path=attributePath Op=PrOperator
        | Path=attributePath Op=(EqOperator|NeOperator|CoOperator|SwOperator|EwOperator|GtOperator|LtOperator|GeOperator|LeOperator) Value=ComparisonValue
        ;

attributePath
        returns [*attr.Path path]
        : Urn? AttributeName (Dot AttributeName)?
        {$path = attr.Parse($ctx.GetText())}
        {$path.Valid()}?<fail={"is not a valid URN"}>
        ;

valueExpression
        : Path=attributePath LxSquareBracket InnerFilter=valueFilter RxSquareBracket
        ;

valueFilter
        : AttributeExpr=attributeExpression # AttributeExprValueFilter
        | Left=attributeExpression Op=AndOperator Right=attributeExpression # AndValueFilter
        | Left=attributeExpression Op=OrOperator Right=attributeExpression # OrValueFilter
        | NotOperator InnerFilter=valueFilter RxBracket # NotValueFilter
        | LxBracket InnerFilter=valueFilter RxBracket # GroupValueFilter
        ;
