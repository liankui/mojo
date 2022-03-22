package compiler

import (
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
)

const irregularCaseRuleName = "compiler.irregular-case-rule"

func init() {
    plugin.RegisterPlugin(NewIrregularCaseRuleCompiler(nil))
}

type IrregularCaseRuleCompiler struct {
    plugin.BasicPlugin
}

func NewIrregularCaseRuleCompiler(options core.Options) *IrregularCaseRuleCompiler {
    return &IrregularCaseRuleCompiler{
        BasicPlugin: plugin.BasicPlugin{
            Name:          irregularCaseRuleName,
            Group:         "compiler",
            GroupPriority: 10,
            Priority:      2,
            Creator: func(options core.Options) plugin.Plugin {
                return NewIrregularCaseRuleCompiler(options)
            },
        },
    }
}

func (c *IrregularCaseRuleCompiler) CompileSourceFile(ctx context.Context, sourceFile *lang.SourceFile) error {
    pkg := context.Package(ctx)
    logs.Infow("enter the plugin", "plugin", c.Name, "method", "CompilePackage", "pkg", pkg.GetFullName(), "file", sourceFile.GetName())

    for _, statement := range sourceFile.Statements {
        if decl := statement.GetDeclaration(); decl != nil {
            return c.ApplyAttribute(ctx, lang.GetAttributes(decl).GetAttribute(core.IrregularCaseRuleAttributeName))
        }
    }
    return nil
}

func (c *IrregularCaseRuleCompiler) ApplyAttribute(ctx context.Context, attribute *lang.Attribute) (err error) {
    if attribute != nil && len(attribute.Arguments) > 0 {
        if len(attribute.Arguments) == 2 {
            rule, replacement := "", ""
            if rule, err = attribute.Arguments[0].GetString(); err != nil {
                return err
            }
            if replacement, err = attribute.Arguments[1].GetString(); err != nil {
                return err
            }
            err = core.ApplyIrregularCaseRuleAttribute(&core.Regex{Expression: rule}, replacement)
        }
    }
    return
}