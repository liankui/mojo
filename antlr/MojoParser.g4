/*
 * [The "BSD license"]
 *  Copyright (c) 2014 Terence Parr
 *  All rights reserved.
 *
 *  Redistribution and use in source and binary forms, with or without
 *  modification, are permitted provided that the following conditions
 *  are met:
 *
 *  1. Redistributions of source code must retain the above copyright
 *     notice, this list of conditions and the following disclaimer.
 *  2. Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *  3. The name of the author may not be used to endorse or promote products
 *     derived from this software without specific prior written permission.
 *
 *  THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 *  IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 *  OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 *  IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 *  INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 *  NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *  DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *  THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 *  THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE
 *
 * Converted from Apple's doc, http://tinyurl.com/n8rkoue, to ANTLR's
 * meta-language.
 */

parser grammar MojoParser;

options {
    tokenVocab=MojoLexer;
}

mojoFile : EOL* (statements)? EOL* EOF;

// Statements

// GRAMMAR OF A STATEMENT
statement
 : declaration
 | expression
 | loopStatement
 | branchStatement
 | controlTransferStatement
 | freeDocument
 ;

freeDocument : document;

statements
    : statement (eos EOL* statement)* SEMI?
    ;

// GRAMMAR OF A LOOP STATEMENT

loopStatement
    : forInStatement
    | whileStatement
  //| repeatWhileStatement
    ;

// GRAMMAR OF A FOR_IN STATEMENT

forInStatement : KEYWORD_FOR pattern EOL* KEYWORD_IN expression EOL* codeBlock ;

// GRAMMAR OF A WHILE STATEMENT

whileStatement : KEYWORD_WHILE conditions EOL* codeBlock ;

conditions : condition (eov EOL* condition)* ;

condition
 : expression
 | optionalBindingCondition
 ;
 
optionalBindingCondition
 : KEYWORD_VAR pattern EOL* initializer
 ;

// GRAMMAR OF A REPEAT-WHILE STATEMENT

//repeatWhileStatement : KEYWORD_REPEATE codeBlock KEYWORD_WHILE expression ;

// GRAMMAR OF A BRANCH STATEMENT

branchStatement
 : ifStatement
 | matchStatement
 ;

// GRAMMAR OF AN IF STATEMENT

ifStatement
  : KEYWORD_IF conditions EOL* codeBlock EOL* elseClause?
  ;

elseClause
  : KEYWORD_ELSE EOL* codeBlock
  | KEYWORD_ELSE EOL* ifStatement
  ;

// GRAMMAR OF A MATCH STATEMENT

matchStatement : KEYWORD_MATCH expression EOL* LCURLY (EOL* matchCases)? EOL* RCURLY  ;
matchCases : matchCase (eos EOL* matchCase)* eos;
matchCase : pattern EOL* RIGHT_RIGHT_ARROWS EOL* ( codeBlock | expression )  ;

// GRAMMAR OF A CONTROL TRANSFER STATEMENT

controlTransferStatement
 : breakStatement
 | continueStatement
 | returnStatement
 ;

// GRAMMAR OF A BREAK STATEMENT

breakStatement : KEYWORD_BREAK;

// GRAMMAR OF A CONTINUE STATEMENT

continueStatement : KEYWORD_CONTINUE;

// GRAMMAR OF A RETURN STATEMENT

returnStatement : KEYWORD_RETURN expression? ;

// Generic Parameters and Arguments

// GRAMMAR OF A GENERIC PARAMETER CLAUSE

genericParameterClause : LT EOL* genericParameters EOL* GT;
genericParameters : genericParameter (eovWithDocument EOL* genericParameter)* eovWithDocument?;
genericParameter
 : typeName
 | typeName ELLIPSIS
 | typeName typeAnnotation
 ;

// GRAMMAR OF A GENERIC ARGUMENT CLAUSE
genericArgumentClause : LT EOL* genericArguments EOL* GT;
genericArguments : genericArgument (eov EOL* genericArgument)*;
genericArgument : type_ attributes?;

// context-sensitive. Allow < as pre, post, or binary op
//lt : {_input.LT(1).getText().equals("<")}? operator ;
//gt : {_input.LT(1).getText().equals(">")}? operator ;

// GRAMMAR OF A DECLARATION
declaration : document? (attributes EOL)?
      ( packageDeclaration
      | importDeclaration
      | constantDeclaration
      | variableDeclaration
      | typeAliasDeclaration
      | functionDeclaration
      | enumDeclaration
      | structDeclaration
      | interfaceDeclaration
      | attributeDeclaration
      );

// GRAMMAR OF A CODE BLOCK
codeBlock : LCURLY (EOL* statements)? EOL* RCURLY ;

// GRAMMAR OF A PACKAGE DECLARATION
packageDeclaration : KEYWORD_PACKAGE packageIdentifier (EOL* objectLiteral)?;

packageIdentifier : packageName (DOT packageName)* ;
packageName : VALUE_IDENTIFIER;

// GRAMMAR OF AN IMPORT DECLARATION
importDeclaration
  : KEYWORD_IMPORT importPath (importAllClause | importValueAsClause | importTypeClause | importGroupClause)?
  ;
importPath : importPathIdentifier (DOT importPathIdentifier)*  ;
importPathIdentifier : declarationIdentifier;

importAllClause : DOT STAR;

importValueAsClause : KEYWORD_AS declarationIdentifier;

importTypeClause : DOT typeName importTypeAsClause?;
importTypeAsClause : KEYWORD_AS typeName;

importGroupClause : DOT LCURLY EOL* importGroup EOL* RCURLY;
importGroup : (importValue | importType) (eov EOL* (importValue | importType))* eov?;
importValue : declarationIdentifier importValueAsClause?;
importType : typeName importTypeAsClause?;

// GRAMMAR OF A CONSTANT DECLARATION

constantDeclaration
  : KEYWORD_CONST patternInitializers
  ;
patternInitializers
  : patternInitializer (eov EOL*  patternInitializer)*
  | LCURLY EOL* documentedPatternInitializer (eov EOL* documentedPatternInitializer)* eov? EOL* RCURLY
  ;

documentedPatternInitializer : document? (attributes EOL)? patternInitializer;

/** rule is ambiguous. can match "var x = 1" with x as pattern
 *  OR with x as expression_pattern.
 *  ANTLR resolves in favor or first choice: pattern is x, 1 is initializer.
 */
patternInitializer : pattern initializer? ;
initializer : assignmentOperator EOL* expression  ;

variableDeclaration
  : KEYWORD_VAR patternInitializers
  | identifierPattern COLON_EQUAL expression
  ;

// GRAMMAR OF A TYPE ALIAS DECLARATION
typeAliasDeclaration
  : KEYWORD_TYPE typeAliasName genericParameterClause? EOL* typeAliasAssignment
  ;

typeAliasName : typeName ;
typeAliasAssignment : assignmentOperator EOL* type_ attributes?;

// GRAMMAR OF A FUNCTION DECLARATION
functionDeclaration : functionHead functionName genericParameterClause? functionSignature (EOL* functionBody)? ;

functionHead : KEYWORD_FUNC ;

functionName : declarationIdentifier | operator ;

functionSignature
 : functionParameterClause followingDocument? (EOL* functionResult)?
 ;
 
functionResult : arrowOperator (labelIdentifier COLON)? type_ attributes? (EOL* followingDocument)? ;

functionBody : LCURLY followingDocument? (EOL* statements)? EOL* RCURLY ;

functionParameterClause
  : LPAREN RPAREN
  | LPAREN EOL* functionParameters EOL* RPAREN
  ;

functionParameters
  : functionParameter (eovWithDocument EOL* functionParameter)* eovWithDocument?
  ;

functionParameter
 : labelIdentifier typeAnnotation (EOL* initializer)?
 | labelIdentifier COLON type_ ELLIPSIS attributes?
 ;

// GRAMMAR OF AN ENUMERATION DECLARATION

enumDeclaration
  : KEYWORD_ENUM enumName genericParameterClause? (EOL* typeInheritanceClause)? EOL* enumBody
  ;

enumBody : LCURLY (followingDocument)? (EOL* enumMembers)? EOL* RCURLY ;

enumName: typeName;
enumMembers : enumMember (eovWithDocument EOL* enumMember)* eovWithDocument?;

enumMember
 : document? (attributes EOL)? declarationIdentifier attributes? (EOL* initializer)?
 //| freeDocument
 ;

// GRAMMAR OF A STRUCTURE DECLARATION
structDeclaration
 : KEYWORD_TYPE structName genericParameterClause? structType
 ;

structName : typeName;
structType : (EOL* typeInheritanceClause)? (EOL* structBody)?;

structBody
   : LCURLY followingDocument? (EOL* structMembers)? EOL* RCURLY
   ;

structMembers
  : structMember (eosWithDocument EOL* structMember)* eosWithDocument?
  //: structMember eosWithDocument?
  //| structMember eosWithDocument EOL* structMembers
  ;

structMember
 : document? (attributes EOL)?
 ( structDeclaration
 | enumDeclaration
 | constantDeclaration
 | typeAliasDeclaration
 | structMemberDeclaration
 //| freeDocument
 )
 ;

structMemberDeclaration
  : declarationIdentifier typeAnnotation (EOL* initializer )?
  ;

// GRAMMAR OF A INTERFACE DECLARATION

interfaceDeclaration
  : KEYWORD_INTERFACE interfaceName genericParameterClause? (EOL* typeInheritanceClause)? EOL* interfaceBody ;

interfaceName : typeName ;
interfaceBody : LCURLY followingDocument? (EOL* interfaceMembers)? EOL* RCURLY ;

interfaceMembers
  : interfaceMember (eosWithDocument EOL* interfaceMember)* eosWithDocument?
  ;

interfaceMember
 : document? (attributes EOL)?
 ( typeAliasDeclaration
 | interfaceMethodDeclaration
 //| freeDocument
 )
 ;

// GRAMMAR OF A INTERFACE METHOD DECLARATION
interfaceMethodDeclaration : functionName genericParameterClause? EOL* functionSignature;

// GRAMMAR OF A ATTRIBUTE DECLARATION
attributeDeclaration
  : document? (EOL* attribute)? KEYWORD_ATTRIBUTE attributeName genericParameterClause? (structType | typeAnnotation);

// Patterns

// GRAMMAR OF A PATTERN
pattern
 : wildcardPattern typeAnnotation?
 | identifierPattern typeAnnotation?
 | tuplePattern typeAnnotation?
 //| enum_value_pattern
 | optionalPattern
 | KEYWORD_IS type_
 | pattern KEYWORD_AS type_
 | expressionPattern
 ;

// GRAMMAR OF A WILDCARD PATTERN

wildcardPattern : UNDERSCORE  ;

// GRAMMAR OF AN IDENTIFIER PATTERN

identifierPattern : declarationIdentifier ;

// GRAMMAR OF A TUPLE PATTERN

tuplePattern : LPAREN tuplePatternElementList? RPAREN  ;
tuplePatternElementList
	:	tuplePatternElement (COMMA tuplePatternElement)*
	;
tuplePatternElement : pattern  ;

// GRAMMAR OF AN ENUMERATION CASE PATTERN

//enum_value_pattern : typeIdentifier? DOT enum_case_name tuple_pattern? ;

// GRAMMAR OF AN OPTIONAL PATTERN
optionalPattern : identifierPattern QUESTION ;

// GRAMMAR OF AN EXPRESSION PATTERN
expressionPattern : expression  ;

// Attributes

// GRAMMAR OF AN ATTRIBUTE
attribute
 : '@' DECIMAL_LITERAL
 | '@' attributeIdentifier genericArgumentClause? attributeArgumentClause?
 ;

attributeIdentifier : ((packageIdentifier DOT)? attributeName);
attributeName : VALUE_IDENTIFIER | keywordAsIdentifierInLabels;

attributeArgumentClause : LPAREN (EOL*  expressions)? EOL* RPAREN  ;

// GRAMMAR OF AN ATTRIBUTES
attributes : attribute (EOL* attribute)* ;

// Expressions

// GRAMMAR OF AN EXPRESSION
expression : prefixExpression binaryExpressions? ;

expressions : expression (eov EOL* expression)* eov?;

// GRAMMAR OF A PREFIX EXPRESSION
prefixExpression
  : prefixOperator postfixExpression
  | postfixExpression
  ;

// GRAMMAR OF A BINARY EXPRESSION
binaryExpression
  : binaryOperator prefixExpression
  | assignmentOperator prefixExpression
  | conditionalOperator prefixExpression
  | typeCastingOperator
  ;

binaryExpressions : binaryExpression+ ;

// GRAMMAR OF A CONDITIONAL OPERATOR
conditionalOperator : QUESTION expression COLON ;

// GRAMMAR OF A TYPE_CASTING OPERATOR
typeCastingOperator
  : KEYWORD_IS type_
  | KEYWORD_AS type_
  ;

// GRAMMAR OF A PRIMARY EXPRESSION
primaryExpression
 : declarationIdentifier genericArgumentClause?
 | literalExpression
 | typeIdentifier DOT declarationIdentifier genericArgumentClause?
 | closureExpression
 | parenthesizedExpression
 | tupleExpression
 | implicitMemberExpression
 | wildcardExpression
 | structConstructionExpression
 //| key_path_expression
 ;

// GRAMMAR OF A LITERAL EXPRESSION
literalExpression
 : numericOperatorLiteral
 | stringOperatorLiteral
 | literal
 | arrayLiteral
 | mapLiteral
 | objectLiteral
 | structLiteral
 ;

numericOperatorLiteral : numericLiteral postfixLiteralOperator ;
stringOperatorLiteral : prefixLiteralOperator stringLiteral;

postfixLiteralOperator : TYPE_IDENTIFIER | VALUE_IDENTIFIER ;
prefixLiteralOperator : VALUE_IDENTIFIER;

arrayLiteral : LBRACK (EOL* arrayLiteralItems)? EOL* RBRACK ;

arrayLiteralItems
  : arrayLiteralItem (eov EOL* arrayLiteralItem)* eov?
 ;
 
arrayLiteralItem : expression ;

mapLiteral
   : LCURLY (EOL* mapLiteralItems)? EOL* RCURLY
   ;

mapLiteralItems
    : (mapLiteralItem (eov EOL* mapLiteralItem)* eov?)
    //| (mapLiteralIntegerItem (eov EOL* mapLiteralIntegerItem)* eov?)
    ;

mapLiteralItem
    : (stringLiteral | integerLiteral) COLON expression;

//mapLiteralStringItem
//    : stringLiteral COLON expression;
//
//mapLiteralIntegerItem
//    : integerLiteral COLON expression
//    ;

objectLiteral
 : LCURLY (EOL* objectLiteralItems)? EOL* RCURLY
 ;

objectLiteralItems
  : objectLiteralItem (eov EOL* objectLiteralItem)* eov?
  ;

objectLiteralItem
    : pathIdentifier (COLON expression)?
    ;

structLiteral
    : typeIdentifier objectLiteral
    ;

structConstructionExpression
    : typeIdentifier functionCallSuffix
    ;

// GRAMMAR OF A CLOSURE EXPRESSION
closureExpression
    : LCURLY statements RCURLY
    | LCURLY closureParameters EOL* RIGHT_ARROW (EOL* type_)? EOL* statements RCURLY
    ;

closureParameters
    : closureParameter (eov EOL* closureParameter)* eov?
    ;

closureParameter
    : functionParameter
    | labelIdentifier
    ;

// GRAMMAR OF A IMPLICIT MEMBER EXPRESSION

implicitMemberExpression : DOT labelIdentifier ; // let a: MyType = .default; static let `default` = MyType()

// GRAMMAR OF A PARENTHESIZED EXPRESSION

parenthesizedExpression
    : LPAREN EOL* (
        expression //|
        //expression GRAPH_RIGHT_PATH expression
      ) EOL* RPAREN ;
// GRAMMAR OF A TUPLE EXPRESSION

tupleExpression
 : LPAREN RPAREN
 | LPAREN tupleElement (COMMA tupleElement)+ RPAREN
 ;

tupleElement
 : expression
 | labelIdentifier COLON expression
 ;

// GRAMMAR OF A WILDCARD EXPRESSION

wildcardExpression : UNDERSCORE ;

// GRAMMAR OF A POSTFIX EXPRESSION
postfixExpression: primaryExpression suffixExpression* postfixOperator?;

suffixExpression
    : functionCallSuffix
    | explicitMemberSuffix
    | subscriptSuffix
    ;

explicitMemberSuffix:
	DOT (
		PURE_DECIMAL_DIGITS
		| identifier (
			genericArgumentClause
			| LPAREN argumentNames RPAREN
		)?
	);

subscriptSuffix: LBRACK functionCallArguments RBRACK;

// GRAMMAR OF A FUNCTION CALL EXPRESSION
functionCallSuffix
    : functionCallArgumentClause? trailingClosures
    | functionCallArgumentClause
	;

functionCallArgumentClause
 : LPAREN RPAREN
 | LPAREN functionCallArguments RPAREN
 ;

functionCallArguments : functionCallArgument ( COMMA functionCallArgument )* ;

functionCallArgument
 : expression
 | labelIdentifier COLON expression
 | operator
 | labelIdentifier COLON operator
 ;

trailingClosures:
	closureExpression labeledTrailingClosures?;

labeledTrailingClosures: labeledTrailingClosure+;

labeledTrailingClosure: identifier COLON closureExpression;

// GRAMMAR OF AN EXPLICIT MEMBER EXPRESSION
argumentNames : argumentName (argumentName)* ;
argumentName : labelIdentifier COLON ;

// GRAMMAR OF A TYPE
type_
 : basicType
 | functionType
 | type_ BANG
 | type_ QUESTION
 | type_ ELLIPSIS
 ;

basicType
 : basicType attributes? EOL* PIPE EOL* basicType attributes? #Union
 | basicType attributes? EOL* AND EOL* basicType attributes? #Intersection
 | primeType #Prime
 ;

primeType
 : arrayType
 | mapType
 | tupleType
 | typeIdentifier
 ;

// GRAMMAR OF A TYPE ANNOTATION
typeAnnotation : COLON type_ attributes?;

// GRAMMAR OF A TYPE IDENTIFIER
typeIdentifier : (packageIdentifier DOT)? typeIdentifierClause (DOT typeIdentifierClause)* ;
typeIdentifierClause : typeName genericArgumentClause? ;

typeName : TYPE_IDENTIFIER ;

// GRAMMAR OF A TUPLE TYPE
tupleType : LPAREN (EOL* tupleTypeElements)? EOL* RPAREN ;
tupleTypeElements : tupleTypeElement (eovWithDocument EOL* tupleTypeElement)* eovWithDocument?;
tupleTypeElement : ( declarationIdentifier COLON )? type_ attributes? ;

// GRAMMAR OF A UNION TYPE

// GRAMMAR OF A FUNCTION TYPE
functionType
 : functionParameterClause arrowOperator type_
 ;

// GRAMMAR OF AN ARRAY TYPE

arrayType : LBRACK type_ attributes? RBRACK ;

// GRAMMAR OF A DICTIONARY TYPE

mapType : LCURLY type_ attributes? COLON type_  attributes? RCURLY ;

// GRAMMAR OF AN OPTIONAL TYPE

// The following sets of rules are mutually left-recursive [mojo_type, optional_type, implicitly_unwrapped_optional_type, metatype_type]
//optional_type : type_ QUESTION ;

// GRAMMAR OF A TYPE INHERITANCE CLAUSE

typeInheritanceClause : COLON EOL* typeInheritances ;
typeInheritances : typeInheritance (eovWithDocument EOL* typeInheritance)* eovWithDocument? ;
typeInheritance : primeType attributes? ;


// ---------- Lexical Structure -----------

// GRAMMAR OF AN IDENTIFIER

// identifier : Identifier | context_sensitive_keyword ;
//
// identifier is context sensitive

// var x = 1; funx x() {}; class x {}
declarationIdentifier
     : VALUE_IDENTIFIER
     | keywordAsIdentifierInDeclarations
     ;

// external, internal argument name
labelIdentifier
     : VALUE_IDENTIFIER
     | keywordAsIdentifierInLabels
     ;

pathIdentifier : declarationIdentifier ( DOT declarationIdentifier)*;

identifier : VALUE_IDENTIFIER | IMPLICIT_PARAMETER_NAME
 ;

// identifier_list : identifier (COMMA identifier)* ;
// 
// identifier is context sensitive
// See: closure_parameter_clause_identifier_list

// Keywords reserved in particular contexts: associativity, convenience, dynamic, didSet, final, get, infix, indirect, lazy, left, mutating, none, nonmutating, optional, override, postfix, precedence, prefix, Protocol, required, right, set, Type, unowned, weak, and willSet. Outside the context in which they appear in the grammar, they can be used as identifiers.
// context_sensitive_keyword :
//  'associativity' | 'convenience' | 'dynamic' | 'didSet'
//  | 'final' | 'get' | 'infix' | 'indirect'
//  | 'lazy' | 'left' | 'mutating' | 'none'
//  | 'nonmutating' | 'optional' | 'override' | 'postfix'
//  | 'precedence' | 'prefix' | 'Protocol' | 'required'
//  | 'right' | 'set' | 'Type' | 'unowned'
//  | 'weak' | 'willSet'
//
// ^- this does not work. "[10].index(of: 10)". "of" is a keyword. "mojo_type(of: self)"
 
 // Added by myself.
 // Tested all alphanumeric tokens in playground.
 // E.g. "let mutating = 1".
 // E.g. "func mutating() {}".
 //
 // In source code of swift there are multiple cases of error diag::keyword_cant_be_identifier.
 // Maybe it is not even a single error when keyword can't be identifier.
 //
keywordAsIdentifierInDeclarations
    : KEYWORD_AND
    | KEYWORD_AS
    | KEYWORD_ATTRIBUTE
    | 'break'
    | 'const'
    | 'enum'
    | 'func'
    | 'if'
    | 'import'
    | 'in'
    | 'interface'
    | KEYWORD_IS
    | 'match'
    | 'not'
    | KEYWORD_NULL
    | 'package'
    | 'repeat'
    | 'return'
    | 'struct'
    | 'type'
    | 'var'
    | 'xor'
    ;

keywordAsIdentifierInLabels
    : 'and'
    | KEYWORD_AS
    | 'attribute'
    | 'break'
    | 'const'
    | 'continue'
    | 'else'
    | 'enum'
    | 'false'
    | 'for'
    | 'func'
    | 'if'
    | 'import'
    | 'in'
    | 'interface'
    | KEYWORD_IS
    | 'match'
    | 'not'
    | 'null'
    | 'or'
    | 'package'
    | 'repeat'
    | 'return'
    | 'struct'
    | 'true'
    | 'type'
    | 'var'
    | 'while'
    | 'xor'
    ;

// GRAMMAR A DOCUMENT_COMMENT

document : LINE_DOCUMENT (EOL LINE_DOCUMENT)* EOL;

followingDocument : FOLLOWING_LINE_DOCUMENT (EOL FOLLOWING_LINE_DOCUMENT)*;

// GRAMMAR OF OPERATORS

/*
From doc on operators:
 The tokens =, ->, //, /*, *\/ [without the escape on \/], .,
 the prefix operators <, &, and ?, the infix
 operator ?, and the postfix operators >, !, and ? are reserved. These tokens
 can’t be overloaded, nor can they be used as custom operators.

 The whitespace around an operator is used to determine whether an operator
 is used as a prefix operator, a postfix operator, or a binary operator.

	* If an operator has whitespace around both sides or around neither
	  side, it is treated as a binary operator. As an example, the +
	  operator in a+b and a + b is treated as a binary operator.

	* If an operator has whitespace on the left side only, it is treated
	  as a prefix unary operator. As an example, the ++ operator in a ++b
	  is treated as a prefix unary operator.

	* If an operator has whitespace on the right side only, it is treated
	  as a postfix unary operator. As an example, the ++ operator in a++ b
	  is treated as a postfix unary operator.

	* If an operator has no whitespace on the left but is followed
	  immediately by a dot (.), it is treated as a postfix unary
	  operator. As an example, the ++ operator in a++.b is treated as a
	  postfix unary operator (a++ .b rather than a ++ .b).

 For the purposes of these rules, the characters (, [, and { before an operator,
 the characters ), ], and } after an operator, and the characters ,, ;, and :
 are also considered whitespace.

 There is one caveat to the rules above. If the ! or ? predefined operator has
 no whitespace on the left, it is treated as a postfix operator, regardless of
 whether it has whitespace on the right. To use the ? as the optional-chaining
 operator, it must not have whitespace on the left. To use it in the ternary
 conditional (? :) operator, it must have whitespace around both sides.

 In certain constructs, operators with a leading < or > may be split
 into two or more tokens. The remainder is treated the same way and may
 be split again. As a result, there is no need to use whitespace to
 disambiguate between the closing > characters in constructs like
 Map<String, Array<Int>>. In this example, the closing >
 characters are not treated as a single token that may then be
 misinterpreted as a bit shift >> operator.
*/

/* these following tokens are also a binaryOperator so much come first as special case */

assignmentOperator : EQUAL ;

/** Need to separate this out from prefixOperator as it's referenced in numericLiteral
 *  as specifically a negation prefix op.
 */
negatePrefixOperator : MINUS;

arrowOperator	      : RIGHT_ARROW;
rangeOperator	      : DOT_DOT ;
halfOpenRangeOperator : DOT_DOT_LT;

/**
 "If an operator has whitespace around both sides or around neither side,
 it is treated as a binary operator. As an example, the + operator in a+b
  and a + b is treated as a binary operator."
*/
binaryOperator
    : rangeOperator
    | halfOpenRangeOperator
    | operator;

/**
 "If an operator has whitespace on the left side only, it is treated as a
 prefix unary operator. As an example, the ++ operator in a ++b is treated
 as a prefix unary operator."
*/
prefixOperator :  operator ;

/**
 "If an operator has whitespace on the right side only, it is treated as a
 postfix unary operator. As an example, the ++ operator in a++ b is treated
 as a postfix unary operator."

 "If an operator has no whitespace on the left but is followed immediately
 by a dot (.), it is treated as a postfix unary operator. As an example,
 the ++ operator in a++.b is treated as a postfix unary operator (a++ .b
 rather than a ++ .b)."
 */
postfixOperator: PLUS_PLUS | MINUS_MINUS;

operator
  : operator_head     operator_characters?
  | dot_operator_head (dot_operator_character)*
  ;

operator_characters: (
		//{_input.get(_input.index()-1).getType()!=WS}? operator_character
		{p.GetTokenStream().Get(p.GetTokenStream().Index()-1).GetTokenType() != MojoParserWS}? operator_character
)+;

operator_character
  : operator_head
  | OPERATOR_FOLLOWING_CHARACTER
  ;

operator_head
  : ('/' | '=' | '-' | '+' | '!' | '*' | '%' | '&' | '|' | '<' | '>' | '^' | '~' | QUESTION) // wrapping in (..) makes it a fast set comparison
  | OPERATOR_HEAD_OTHER
  ;

dot_operator_head 		: DOT ;
dot_operator_character  : DOT | operator_character ;

// GRAMMAR OF A LITERAL
literal : numericLiteral | stringLiteral | boolLiteral | nullLiteral  ;

boolLiteral : KEYWORD_TRUE | KEYWORD_FALSE ;

nullLiteral : KEYWORD_NULL ;

numericLiteral
 : negatePrefixOperator? integerLiteral
 | negatePrefixOperator? FLOAT_LITERAL
 ;

// GRAMMAR OF AN INTEGER LITERAL
integerLiteral
 : BINARY_LITERAL
 | OCTAL_LITERAL
 | DECIMAL_LITERAL
 | PURE_DECIMAL_DIGITS
 | HEXADECIMAL_LITERAL
 ;

// GRAMMAR OF A STRING LITERAL
stringLiteral
  : STATIC_STRING_LITERAL
  | INTERPOLATED_STRING_LITERAL
  ;

eos : SEMI | EOL;
eov : COMMA | EOL;

eosWithDocument
  : SEMI (followingDocument EOL)?
  | followingDocument? EOL
  ;

eovWithDocument
  : COMMA (followingDocument EOL)?
  | followingDocument? EOL
  ;