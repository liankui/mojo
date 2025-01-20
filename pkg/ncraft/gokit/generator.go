package gokit

import (
	"github.com/mojo-lang/mojo/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/pkg/ncraft/gokit/generator"
	"github.com/mojo-lang/mojo/pkg/util"
)

type Options = generator.Options

func GenerateClient(ds *data.Service, options Options) error {
	files, err := options.GenerateClient(ds)
	if err != nil {
		return err
	}

	return generateFiles(files, options.Output)
}

func GenerateService(ds *data.Service, options Options) error {
	files, err := options.GenerateService(ds)
	if err != nil {
		return err
	}

	return generateFiles(files, options.Output)
}

func generateFiles(files []*util.GeneratedFile, output string) error {
	guard := &util.PathGuard{
		OnlyClearGenerated: true,
		Suffixes:           []string{".go", ".mod", ".md", ".sh", ".yaml"},
	}
	for _, file := range files {
		if err := file.WriteTo(output, guard); err != nil {
			return err
		}
	}
	return nil
}
