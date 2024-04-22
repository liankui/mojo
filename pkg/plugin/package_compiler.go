package plugin

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type PackageCompiler interface {
	CompilePackage(ctx context.Context, pkg *lang.Package) error
}

func IsPackageCompiler(c interface{}) bool {
	if _, ok := c.(PackageCompiler); ok {
		return true
	}
	return false
}
