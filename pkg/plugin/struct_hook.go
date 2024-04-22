package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type StructPreHook interface {
	PreStruct(ctx context.Context, pkg *lang.StructDecl) error
}

type StructPostHook interface {
	PostStruct(ctx context.Context, pkg *lang.StructDecl)
}
