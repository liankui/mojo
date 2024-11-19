package model

import (
	"bytes"
	"io"
	"strings"

	"github.com/mojo-lang/mojo/pkg/ncraft/data"
	"github.com/mojo-lang/mojo/pkg/ncraft/gokit/generator/model/templates"
	"github.com/mojo-lang/mojo/pkg/util"

	_go "github.com/mojo-lang/mojo/pkg/ncraft/go"
)

const (
	// TemplatePath is the path to the entity go template file.
	TemplatePath = "pkg/model/ENTITY_model.go.tmpl"
	// ExampleTemplatePath is the path to the go example model file.
	ExampleTemplatePath = "pkg/model/example_model.go.tmpl"
)

type Model struct{}

func (m Model) Generate(path string, service *data.Service) ([]*util.GeneratedFile, error) {
	var files []*util.GeneratedFile

	// for _, v := range service.Entities {
	for _, v := range service.ImportedMessages {
		if v.Name == "Null" {
			continue
		}

		if len(service.Go.ImportedTypePaths) > 0 {
			v.Go.ImportPath = service.Go.ImportedTypePaths[0]
		}
		reader, err := util.ApplyTemplate("Model", templates.EntityModel, v, service.FuncMap)
		if err != nil {
			return nil, err
		}

		codeBytes, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		formattedCode := _go.FormatCodeBytes(codeBytes)
		files = append(files, &util.GeneratedFile{
			Name:                strings.TrimSuffix(strings.Replace(path, "ENTITY", strings.ToLower(v.Name), 1), ".tmpl"),
			Reader:              bytes.NewReader(formattedCode),
			SkipIfExist:         true,
			SkipIfUserCodeMixed: true,
		})
	}

	return files, nil
}

func (m Model) GenerateExample(path string, service *data.Service) ([]*util.GeneratedFile, error) {
	var files []*util.GeneratedFile

	formattedCode := _go.FormatCodeBytes([]byte(templates.ExampleEntityModel))
	files = append(files, &util.GeneratedFile{
		Name:                path,
		Reader:              bytes.NewReader(formattedCode),
		SkipIfExist:         true,
		SkipIfUserCodeMixed: true,
	})

	return files, nil
}
