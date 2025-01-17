package gokit

import (
	"os"
	"path"

	"github.com/mojo-lang/core/go/pkg/logs"

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

		filepath := path.Join(output, file.Name)
		if err := grantExecutePermission(filepath); err != nil {
			logs.Warnw("failed to grant permissions", "filepath", filepath, "error", err)
		}
	}
	return nil
}

func grantExecutePermission(filepath string) error {
	if path.Ext(filepath) == ".sh" {
		info, err := os.Stat(filepath)
		if err != nil {
			return err
		}

		mode := info.Mode()
		newMode := mode | 0111

		if err := os.Chmod(filepath, newMode); err != nil {
			return err
		}
	}

	return nil
}
