package printer

import (
	"github.com/mojo-lang/mojo/pkg/printer"
	"github.com/mojo-lang/mojo/pkg/sql/printer/ansi"
)

type SqlitePrinter struct {
	ansi.SQLPrinter
}

func New(printer *printer.Printer) *SqlitePrinter {
	return &SqlitePrinter{ansi.SQLPrinter{Printer: printer}}
}
