package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type ExpressionCompiler interface {
	CompileExpression(ctx context.Context, decl *lang.Expression) error
}

func IsExpressionCompiler(c interface{}) bool {
	if _, ok := c.(ExpressionCompiler); ok {
		return true
	}
	return false
}
