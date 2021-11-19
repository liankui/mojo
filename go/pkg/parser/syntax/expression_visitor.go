package syntax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
)

type ExpressionVisitor struct {
	*BaseMojoParserVisitor
}

func NewExpressionVisitor() *ExpressionVisitor {
	return &ExpressionVisitor{}
}

func GetExpression(ctx IExpressionContext) *lang.Expression {
	if ctx != nil {
		visitor := NewExpressionVisitor()
		switch expr := ctx.Accept(visitor).(type) {
		case *lang.Expression:
			return expr
		case *lang.ArrayLiteralExpr:
			return lang.NewArrayLiteralExpression(expr)
		case *lang.ObjectLiteralExpr:
			return lang.NewObjectLiteralExpression(expr)
		case *lang.MapLiteralExpr:
			return lang.NewMapLiteralExpression(expr)
		case *lang.IdentifierExpr:
			return lang.NewIdentifierExpression(expr)
		default:
			fmt.Print("===> error")
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	if prefixCtx := ctx.PrefixExpression(); prefixCtx != nil {
		if expr, ok := prefixCtx.Accept(e).(*lang.Expression); ok {
			if binaryCtx := ctx.BinaryExpressions(); binaryCtx != nil {
				if binaryExprs, ok := binaryCtx.Accept(e).([]*lang.BinaryExpr); ok {
					return BinaryExprParser{}.Parse(expr, binaryExprs)
				}
			}
			return expr
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitBinaryExpressions(ctx *BinaryExpressionsContext) interface{} {
	binaryCtxs := ctx.AllBinaryExpression()
	var binaries []*lang.BinaryExpr
	for _, binaryCtx := range binaryCtxs {
		if expr, ok := binaryCtx.Accept(e).(*lang.BinaryExpr); ok {
			binaries = append(binaries, expr)
		}
	}
	return binaries
}

func (e *ExpressionVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	if prefixExprCtx := ctx.PrefixExpression(); prefixExprCtx != nil {
		if binaryOperator := ctx.BinaryOperator(); binaryOperator != nil {
			if expression, ok := prefixExprCtx.Accept(e).(*lang.Expression); ok {
				return &lang.BinaryExpr{
					Operator:          &lang.Operator{Symbol: binaryOperator.GetText()},
					LeftHandArgument:  nil,
					RightHandArgument: expression,
				}
			}
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitPrefixExpression(ctx *PrefixExpressionContext) interface{} {
	postfixCtx := ctx.PostfixExpression()
	if postfixCtx != nil {
		if expression, ok := postfixCtx.Accept(e).(*lang.Expression); ok {
			operator := ctx.PrefixOperator()
			if operator != nil {
				return lang.NewPrefixUnaryExpression(&lang.PrefixUnaryExpr{
					Operator:   operator.GetText(),
					Expression: expression,
				})
			}
			return expression
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	primaryCtx := ctx.PrimaryExpression()
	if primaryCtx != nil {
		if expression, ok := primaryCtx.Accept(e).(*lang.Expression); ok {
			suffixExprs := ctx.AllSuffixExpression()

			for _, suffixExpr := range suffixExprs {
				visitor := &SuffixExpressionVisitor{
					PrimaryExpression: expression,
				}
				var expr *lang.Expression
				if expr, ok = visitor.Visit(suffixExpr).(*lang.Expression); ok {
					expression = expr
				}
			}

			operator := ctx.PostfixOperator()
			if operator != nil {
				return lang.NewPostfixUnaryExpression(&lang.PostfixUnaryExpr{
					Operator:   operator.GetText(),
					Expression: expression,
				})
			}

			return expression
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	if literalCtx := ctx.LiteralExpression(); literalCtx != nil {
		return literalCtx.Accept(e)
	}

	if identifierCtx := ctx.DeclarationIdentifier(); identifierCtx != nil {
		if identifier, ok := identifierCtx.Accept(e).(*lang.Identifier); ok {
			arguments := GetGenericArguments(ctx.GenericArgumentClause())
			return lang.NewIdentifierExpression(&lang.IdentifierExpr{
				Identifier:       identifier,
				GenericArguments: arguments,
			})
		}
	}

	if parenthesizedCtx := ctx.ParenthesizedExpression(); parenthesizedCtx != nil {
		return parenthesizedCtx.Accept(e)
	}

	if closureCtx := ctx.ClosureExpression(); closureCtx != nil {
		return closureCtx.Accept(&ClosureExprVisitor{})
	}

	if wildcardCtx := ctx.WildcardExpression(); wildcardCtx != nil {
		return wildcardCtx.Accept(e)
	}

	if tupleCtx := ctx.TupleExpression(); tupleCtx != nil {
		return tupleCtx.Accept(e)
	}

	return nil
}

func (e *ExpressionVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	literalCtx := ctx.Literal()
	if literalCtx != nil {
		return literalCtx.Accept(e)
	}

	arrayCtx := ctx.ArrayLiteral()
	if arrayCtx != nil {
		return arrayCtx.Accept(NewArrayLiteralVisitor())
	}

	objectCtx := ctx.ObjectLiteral()
	if objectCtx != nil {
		return objectCtx.Accept(NewObjectLiteralVisitor())
	}

	mapCtx := ctx.MapLiteral()
	if mapCtx != nil {
		return mapCtx.Accept(NewMapLiteralVisitor())
	}

	return nil
}

func (e *ExpressionVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	nullCtx := ctx.NullLiteral()
	if nullCtx != nil {
		return lang.NewNullLiteralExpression(&lang.NullLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
		})
	}

	boolCtx := ctx.BoolLiteral()
	if boolCtx != nil {
		return lang.NewBoolLiteralExpression(&lang.BoolLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
			Kind:          0,
			Implicit:      false,
			Value:         boolCtx.GetText() == "true",
		})
	}

	numericCtx := ctx.NumericLiteral()
	if numericCtx != nil {
		return numericCtx.Accept(e)
	}

	strCtx := ctx.StringLiteral()
	if strCtx != nil {
		return strCtx.Accept(e)
	}

	return nil
}

func (e *ExpressionVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	isNegative := false
	negatePrefix := ctx.NegatePrefixOperator()
	if negatePrefix != nil {
		isNegative = true
	}

	integerCtx := ctx.IntegerLiteral()
	if integerCtx != nil {
		v := integerCtx.Accept(e).(*lang.IntegerLiteralExpr)
		if v != nil {
			v.IsNegative = isNegative
			return lang.NewIntegerLiteralExpression(v)
		}
	}

	floatCtx := ctx.FLOAT_LITERAL()
	if floatCtx != nil {
		v, err := strconv.ParseFloat(floatCtx.GetText(), 64)
		if err != nil {
			return lang.NewFloatLiteralExpression(&lang.FloatLiteralExpr{
				StartPosition: nil,
				EndPosition:   nil,
				Kind:          0,
				Implicit:      false,
				IsNegative:    isNegative,
				Value:         v,
			})
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	decimal := ctx.DECIMAL_LITERAL()
	if decimal != nil {
		v, err := strconv.ParseUint(decimal.GetText(), 10, 64)
		if err != nil {
			//FIXME log error
			return nil
		}

		return &lang.IntegerLiteralExpr{
			StartPosition: nil,
			EndPosition:   nil,
			Kind:          0,
			Implicit:      false,
			IsNegative:    false,
			Value:         int64(v),
		}
	}

	return nil
}

func (e *ExpressionVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	static := ctx.STATIC_STRING_LITERAL()
	if static != nil {
		value := static.GetText()
		if strings.HasPrefix(value, "'") {
			value = value[1 : len(value)-1]
		} else if strings.HasPrefix(value, "\"") {
			value = value[1 : len(value)-1]
		}

		return lang.NewStringLiteralExpression(&lang.StringLiteralExpr{
			//StartPosition:        nil,
			//EndPosition:          nil,
			//Kind:                 0,
			Implicit: false,
			Value:    value,
		})
	}

	return nil
}

func (e *ExpressionVisitor) VisitDeclarationIdentifier(ctx *DeclarationIdentifierContext) interface{} {
	identifier := ctx.GetText()
	if len(identifier) > 0 {
		return &lang.Identifier{
			Name: identifier,
		}
	}
	return nil
}

func (e *ExpressionVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	if exprCtx := ctx.Expression(); exprCtx != nil {
		return exprCtx.Accept(e)
	}
	return nil
}

func (e *ExpressionVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) interface{} {
	if ctx.UNDERSCORE().GetText() == "_" {
		return lang.NewWildcardExpression(&lang.WildcardExpr{
			StartPosition: nil,
			EndPosition:   nil,
		})
	}
	return nil
}

func (e *ExpressionVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	elementCtxes := ctx.AllTupleElement()
	tupleExpr := &lang.TupleExpr{
		StartPosition: GetPosition(ctx.GetStart()),
		EndPosition:   GetPosition(ctx.GetStop()),
	}

	for _, elementCtx := range elementCtxes {
		if element, ok := elementCtx.Accept(e).(*lang.Argument); ok {
			tupleExpr.Elements = append(tupleExpr.Elements, element)
			if len(element.Label) > 0 {
				tupleExpr.HasElementLabels = true
			}
		}
	}

	if len(tupleExpr.Elements) > 0 {
		return lang.NewTupleExpression(tupleExpr)
	}
	return nil
}

func (e *ExpressionVisitor) VisitTupleElement(ctx *TupleElementContext) interface{} {
	element := &lang.Argument{}
	if exprCtx := ctx.Expression(); exprCtx != nil {
		element.Value = GetExpression(exprCtx)

		element.StartPosition = GetPosition(ctx.GetStart())
		element.EndPosition = GetPosition(ctx.GetStop())
	}

	if identifierCtx := ctx.LabelIdentifier(); identifierCtx != nil {
		element.Label = identifierCtx.GetText()
	}

	return element
}