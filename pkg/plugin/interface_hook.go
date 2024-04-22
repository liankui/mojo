package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type InterfacePreHook interface {
	PreInterface(ctx context.Context, pkg *lang.InterfaceDecl) error
}

type InterfacePostHook interface {
	PostInterface(ctx context.Context, pkg *lang.InterfaceDecl)
}
