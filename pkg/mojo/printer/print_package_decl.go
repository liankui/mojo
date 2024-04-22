package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

func (p *Printer) PrintPackageDecl(ctx context.Context, decl *lang.PackageDecl) *Printer {
	if decl == nil || p.GetError() != nil {
		return p
	}

	breaker := &OnceLineBreaker{}
	p.printPreDecl(ctx, decl, breaker).
		Break(p).
		PrintLine("package", " ", decl.Name, " ").
		PrintObjectLiteralExpr(ctx, decl.PackageLiteralExpr)

	return p
}
