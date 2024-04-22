package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type EnumCompiler interface {
	CompileEnum(ctx context.Context, decl *lang.EnumDecl) error
}

func IsEnumCompiler(c interface{}) bool {
	if _, ok := c.(EnumCompiler); ok {
		return true
	}
	return false
}
