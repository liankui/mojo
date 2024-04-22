package printer

import (
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

func (p *Printer) PrintGenericParameter(ctx context.Context, parameter *lang.GenericParameter) *Printer {
	if parameter == nil || p.GetError() != nil {
		return p
	}

	p.PrintRaw(parameter.Name)

	if parameter.Constraint != nil {
		p.PrintRaw(": ").PrintNominalType(ctx, parameter.Constraint)
	}

	return p
}
