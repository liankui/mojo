package render

import (
	"bytes"
	"io"
	"strings"
	"text/template"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/compiler"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/httptransport"
	"github.com/mojo-lang/mojo/go/pkg/ncraft/gokit/generator/types"
	"github.com/pkg/errors"
)

// Data is passed to templates as the executing struct; its fields
// and methods are used to modify the template
type Data struct {
	// repository path for the directory containing the definition service files
	RepositoryPath string

	// import path for .pb.go files containing service structs
	ApiImportPath string

	ExternalMessageImports []string
	ExternalStructs        []string
	ExternalEnums          []string

	// PackageName is the name of the package containing the service definition
	PackageName     string
	FullPackageName string

	AllInterfaceNames []string

	// GRPC/Protobuf service, with all parameters and return values accessible
	Interface  *types.Interface
	ClientArgs *types.ClientArguments

	// A helper struct for generating http transport functionality.
	HTTPHelper *httptransport.Helper

	FuncMap template.FuncMap

	Version     string
	VersionDate string

	Extension map[string]interface{}
}

func unifyStringArray(array []string) []string {
	index := make(map[string]bool)
	for _, item := range array {
		if _, ok := index[item]; !ok {
			index[item] = true
		}
	}
	var result []string
	for key, _ := range index {
		result = append(result, key)
	}
	return result
}

func NewData(sd *types.Service, conf Config) (*Data, error) {
	return &Data{
		RepositoryPath:         conf.Repository,
		ApiImportPath:          sd.ApiImportPath,
		PackageName:            sd.PkgName,
		FullPackageName:        sd.FullPkgName,
		Interface:              sd.Interface,
		AllInterfaceNames:      sd.AllInterfaceNames,
		ExternalMessageImports: unifyStringArray(sd.ImportPaths),
		ExternalStructs:        unifyStringArray(sd.ImportStructs),
		ExternalEnums:          unifyStringArray(sd.ImportEnums),
		ClientArgs:             types.NewClientArguments(sd.Interface),
		HTTPHelper:             httptransport.NewHelper(sd.Interface),
		FuncMap:                FuncMap,
		Extension:              conf.ExtensionData,
		Version:                conf.Version,
		VersionDate:            conf.VersionDate,
	}, nil
}

// ApplyTemplate applies the passed template with the Data
func (e *Data) ApplyTemplate(tmpl string, tmplName string) (io.Reader, error) {
	return ApplyTemplate(tmpl, tmplName, e, e.FuncMap)
}

// ApplyTemplate is a helper methods that packages can call to render a
// template with any data and func map
func ApplyTemplate(tmpl string, tmplName string, data interface{}, funcMap template.FuncMap) (io.Reader, error) {
	codeTemplate, err := template.New(tmplName).Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create template")
	}

	outputBuffer := bytes.NewBuffer(nil)
	err = codeTemplate.Execute(outputBuffer, data)
	if err != nil {
		return nil, errors.Wrap(err, "template error")
	}

	return outputBuffer, nil
}

// FuncMap contains a series of utility functions to be passed into
// templates and used within those templates.
var FuncMap = template.FuncMap{
	"ToLower":      strings.ToLower,
	"GoName":       strcase.ToCamel,
	"PackageName":  compiler.GetPackageName,
	"ToSnake":      strcase.ToSnake,
	"ToKebab":      strcase.ToKebab,
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}
