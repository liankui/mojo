package printer

import (
	"bytes"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/printer"
)

const needIndent = "need-indent"

// A Config node controls the output.
type (
	Config = printer.Config
	Cursor = printer.Cursor
)

type Printer struct {
	P      *printer.Printer
	Buffer *bytes.Buffer
}

func New(config *Config) *Printer {
	buffer := bytes.NewBuffer(nil)
	p := &Printer{
		P:      printer.New(config, buffer),
		Buffer: buffer,
	}
	return p
}

func (p *Printer) GetIndent() int {
	return p.P.GetIndent()
}

func (p *Printer) IsNewLine() bool {
	return p.P.IsNewLine()
}

func (p *Printer) GetError() error {
	return p.P.Error
}

func (p *Printer) SetError(err error) *Printer {
	p.P.Error = err
	return p
}

func (p *Printer) Indent() *Printer {
	p.P.Indent()
	return p
}

func (p *Printer) Outdent() *Printer {
	p.P.Outdent()
	return p
}

func (p *Printer) BreakLine() *Printer {
	p.P.BreakLine()
	return p
}

func (p *Printer) PrintRaw(values ...interface{}) *Printer {
	p.P.PrintRaw(values...)
	return p
}

func (p *Printer) PrintIndent() *Printer {
	p.P.PrintIndent()
	return p
}

func (p *Printer) PrintLine(values ...interface{}) *Printer {
	p.P.PrintLine(values...)
	return p
}

func (p *Printer) PrintTo(cursor Cursor) *Printer {
	p.P.PrintTo(cursor)
	return p
}

func (p *Printer) PrintTerm(ctx context.Context, term *lang.Term) *Printer {
	p.P.PrintTerm(ctx, term)
	return p
}

func (p *Printer) PrintDocument(ctx context.Context, document *lang.Document) *Printer {
	p.P.PrintDocument(ctx, document)
	return p
}

func (p *Printer) PrintComments(ctx context.Context, comments ...*lang.Comment) *Printer {
	p.P.PrintComments(ctx, comments...)
	return p
}
