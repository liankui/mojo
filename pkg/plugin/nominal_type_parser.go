package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type NominalTypeParser interface {
	ParseNominalType(ctx context.Context, typ *lang.NominalType) error
}
