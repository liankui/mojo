package compiler

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

const labelFormatName = "compiler.label-format"

func init() {
    plugin.RegisterPlugin(NewLabelFormatCompiler(nil))
}

// LabelFormatCompiler used for union, tuple
type LabelFormatCompiler struct {
    plugin.BasicPlugin
}

func NewLabelFormatCompiler(options core.Options) *LabelFormatCompiler {
    return &LabelFormatCompiler{
        BasicPlugin: plugin.BasicPlugin{
            Name:          labelFormatName,
            Group:         "compiler",
            GroupPriority: 10,
            Priority:      13,
            Creator: func(options core.Options) plugin.Plugin {
                return NewLabelFormatCompiler(options)
            },
        },
    }
}

func (c *LabelFormatCompiler) CompileTypeAlias(ctx context.Context, decl *lang.TypeAliasDecl) error {
    pkg := context.Package(ctx)
    logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompileTypeAlias", "pkg", pkg.FullName, "decl", decl.Name)

    if decl.Type.GetFullName() == core.UnionTypeFullName {
        if lang.HasAttribute(decl.Attributes, core.LabelFormatAttributeName) {
        }
    }
    return nil
}
