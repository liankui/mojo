package printer

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db/sql"

	"github.com/mojo-lang/mojo/pkg/context"
)

type SQLPrinter interface {
	PrintExpression(ctx context.Context, expr *sql.Expression)

	PrintNullLiteralExpr(ctx context.Context, expr *sql.NullLiteralExpr)
	PrintIdentifyExpr(ctx context.Context, expr *sql.IdentifierExpr)
}
