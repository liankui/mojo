package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type SourceFilePreHook interface {
	PreSourceFile(ctx context.Context, pkg *lang.SourceFile) error
}

type SourceFilePostHook interface {
	PostSourceFile(ctx context.Context, pkg *lang.SourceFile)
}
