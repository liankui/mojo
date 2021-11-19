// Package generator generates a gokit service based on a service definition.
package gokit

import (
	"bytes"
	"github.com/iancoleman/strcase"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/handlers"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/render"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/template"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
	"github.com/mojo-lang/mojo/go/pkg/util"
	"github.com/pkg/errors"
	"go/format"
	"io"
	"io/ioutil"
	path2 "path"
	"strings"
)

func GenerateClient(sd *types.Service, conf render.Config) error {
	return GenerateTemplatedService(sd, conf, template.ClientNames(), template.Client)
}

func GenerateService(sd *types.Service, conf render.Config) error {
	return GenerateTemplatedService(sd, conf, template.ServiceNames(), template.Service)
}

// GenerateService returns a gokit service generated from a service definition (svcdef),
// the package to the root of the generated service goPackage, the package
// to the .pb.go service struct files (goPBPackage) and any prevously generated files.
func GenerateTemplatedService(sd *types.Service, conf render.Config, tmplPaths []string, getter template.FileGetter) error {
	data, err := render.NewData(sd, conf)
	if err != nil {
		return errors.Wrap(err, "cannot create template data")
	}

	if tmpl := conf.HandlerTemplate; tmpl != nil {
		if len(tmpl.Interface) > 0 {
			handlers.ResetHandlerInterface(tmpl.Interface)
		}
		if len(tmpl.Methods) > 0 {
			handlers.ResetHandlerMethods(tmpl.Methods)
		}
		if len(tmpl.Extension) > 0 {
			handlers.ResetHandlerExtension(tmpl.Extension)
		}
	}

	codeGenFiles := make(map[string]io.Reader)

	// Remove the suffix "-service" since it's added back in by templatePathToActual
	svcName := strcase.ToKebab(sd.Interface.Name)

	for _, tmplPath := range tmplPaths {
		// Re-derive the actual path for this file based on the service output
		// path provided by the ncraft main.go
		actualPath := templatePathToActual(tmplPath, sd.PkgName, svcName)
		file, err := generateTemplateFile(tmplPath, data, conf.PreviousFiles[actualPath], getter)
		if err != nil {
			return errors.Wrap(err, "cannot render template")
		}

		codeGenFiles[actualPath] = file
	}

	for name, file := range codeGenFiles {
		//if util.IsExist(dir) {
		//	os.RemoveAll(dir)
		//}

		name := path2.Join(conf.Output, name)
		path := path2.Dir(name)
		util.CreateDir(path)

		if bytes, err := io.ReadAll(file); err != nil {
			return err
		} else {
			ioutil.WriteFile(name, bytes, 0666)
		}
	}

	return nil
}

// generateTemplateFile contains logic to choose how to render a template file
// based on path and if that file was generated previously. It accepts a
// template path to render, a templateExecutor to apply to the template, and a
// map of paths to files for the previous generation. It returns a
// io.Reader representing the generated file.
func generateTemplateFile(tmplPath string, data *render.Data, prevFile io.Reader, getter template.FileGetter) (io.Reader, error) {
	var genCode io.Reader
	var err error

	// Get the actual path to the file rather than the template file path
	actualFP := templatePathToActual(tmplPath, data.PackageName, data.Interface.Name)

	switch tmplPath {
	case handlers.ServerHandlerPath:
		h, err := handlers.New(data.Interface, prevFile)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse previous handler: %q", actualFP)
		}

		if genCode, err = h.Render(tmplPath, data); err != nil {
			return nil, errors.Wrapf(err, "cannot render template: %s", tmplPath)
		}
	case handlers.HookPath:
		hook := handlers.NewHook(prevFile)
		if genCode, err = hook.Render(tmplPath, data); err != nil {
			return nil, errors.Wrapf(err, "cannot render template: %s", tmplPath)
		}
	case handlers.MiddlewaresPath:
		m := handlers.NewMiddlewares()
		m.Load(prevFile)
		if genCode, err = m.Render(tmplPath, data); err != nil {
			return nil, errors.Wrapf(err, "cannot render template: %s", tmplPath)
		}
	default:
		if genCode, err = applyTemplateFromPath(tmplPath, data, getter); err != nil {
			return nil, errors.Wrapf(err, "cannot render template: %s", tmplPath)
		}
	}

	codeBytes, err := ioutil.ReadAll(genCode)
	if err != nil {
		return nil, err
	}

	// ignore error as we want to write the code either way to inspect after
	// writing to disk
	formattedCode := formatCode(codeBytes)

	return bytes.NewReader(formattedCode), nil
}

// templatePathToActual accepts a templateFilePath and the svcName of the
// service and returns what the relative file path of what should be written to
// disk
func templatePathToActual(tmplPath, pkgName string, svcName string) string {
	// Switch "NAME" in path with svcName.
	// i.e. for svcName = addsvc; /NAME-server -> /addsvc-service/addsvc-server
	actual := strings.Replace(tmplPath, "PKGNAME", pkgName, -1)
	actual = strings.Replace(actual, "NAME", svcName, -1)
	actual = strings.TrimSuffix(actual, ".tmpl")
	return actual
}

// applyTemplateFromPath calls applyTemplate with the template at tmplPath
func applyTemplateFromPath(tmplPath string, data *render.Data, getter template.FileGetter) (io.Reader, error) {
	tmplBytes, err := getter(tmplPath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to find template file: %v", tmplPath)
	}

	return data.ApplyTemplate(string(tmplBytes), tmplPath)
}

// formatCode takes a string representing golang code and attempts to return a
// formatted copy of that code.  If formatting fails, a warning is logged and
// the original code is returned.
func formatCode(code []byte) []byte {
	formatted, err := format.Source(code)
	if err != nil {
		//log.WithError(err).Warn("Code formatting error, generated service will not build, outputting unformatted code")
		// return code so at least we get something to examine
		return code
	}

	return formatted
}