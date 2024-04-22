package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type NominalTypePreHook interface {
	PreNominalType(ctx context.Context, typ *lang.NominalType) error
}

type NominalTypePostHook interface {
	PostNominalType(ctx context.Context, typ *lang.NominalType)
}
