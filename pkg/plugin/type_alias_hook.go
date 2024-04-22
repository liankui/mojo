package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type TypeAliasPreHook interface {
	PreTypeAlias(ctx context.Context, pkg *lang.TypeAliasDecl) error
}

type TypeAliasPostHook interface {
	PostTypeAlias(ctx context.Context, pkg *lang.TypeAliasDecl)
}
