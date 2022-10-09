package protobuf

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	"github.com/mojo-lang/mojo/go/pkg/protobuf/generator"
	path2 "path"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/cmd/build/builder"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
)

type Builder struct {
	builder.Builder
	Output string
}

func (b Builder) Build() ([]*descriptor.File, error) {
	logs.Infow("begin to build protobuf.", "package", b.Package.FullName, "path", b.Path)

	compiler := generator.NewCompiler()
	err := compiler.CompilePackage(context.Empty(), b.Package)
	if err != nil {
		logs.Errorw("failed to compile protobuf", "package", b.Package.FullName, "error", err.Error())
		return nil, err
	}

	files := compiler.Descriptors.Filter(b.Package.FullName, false)
	if !b.APIEnabled {
		logs.Infow("disable generation, skip to generate protobuf.")
		return files, nil
	}

	output := path2.Join(b.GetAbsolutePath(), "protobuf")
	if len(b.Output) > 0 {
		output = b.Output
	}

	generator := generator.NewGenerator(files)
	return generator.Generate(output)
}
