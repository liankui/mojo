package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type FileParser interface {
	ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error)
}
