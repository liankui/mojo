package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type EnumPreHook interface {
	PreEnum(ctx context.Context, decl *lang.EnumDecl) error
}

type EnumPostHook interface {
	PostEnum(ctx context.Context, decl *lang.EnumDecl)
}
