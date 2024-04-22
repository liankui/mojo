package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type FunctionPreHook interface {
	PreFunction(ctx context.Context, pkg *lang.FunctionDecl) error
}

type FunctionPostHook interface {
	PostFunction(ctx context.Context, pkg *lang.FunctionDecl)
}
