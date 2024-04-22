package mpm

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/plugin"
	"github.com/mojo-lang/mojo/pkg/util"
)

const pluginName = "mpm.dependency-parser"

func init() {
	plugin.RegisterPlugin(NewDependencyParser(nil))
}

type DependencyParser struct {
	plugin.BasicPlugin

	parsedPackages map[string]*lang.Package
}

func NewDependencyParser(options core.Options) *DependencyParser {
	return &DependencyParser{
		BasicPlugin: plugin.BasicPlugin{
			Name:          pluginName,
			Group:         "mpm",
			GroupPriority: 0,
			Priority:      1,
			Creator: func(options core.Options) plugin.Plugin {
				return NewDependencyParser(options)
			},
		},
		parsedPackages: make(map[string]*lang.Package),
	}
}

// ParseFile
// TODO implement the imports
func (p *DependencyParser) ParseFile(ctx context.Context, fileName string) (*lang.SourceFile, error) {
	_ = ctx
	return nil, nil
}

func (p *DependencyParser) ParsePackage(ctx context.Context, pkg *lang.Package) error {
	if util.IsPackageProcessed(pkg, pluginName) {
		return nil
	}
	logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackage", "pkg", pkg.FullName)

	if plugins := plugin.ContextPlugins(ctx); plugins != nil {
		for _, dep := range pkg.ResolvedDependencies {
			logs.Infow("begin to parse mojo dependency", "dependency", dep.FullName)

			cloned := plugins.Copy()
			if err := cloned.ParsePackage(plugin.WithPlugins(ctx, cloned), dep); err != nil && !core.IsSkipError(err) {
				return err
			}
		}
		if plugins.Next() != nil {
			if err := plugins.ParsePackage(ctx, pkg); err != nil {
				return err
			}
		}
	}

	util.SetPackageProcessed(pkg, pluginName)
	return nil
}

func (p *DependencyParser) ParsePath(ctx context.Context, pkgPath string) (*lang.Package, error) {
	workingDir := plugin.ContextWorkingDir(ctx)
	logs.Infow("enter the plugin", "plugin", p.Name, "method", "ParsePackagePath", "workingDir", workingDir, "path", pkgPath)

	if strings.HasPrefix(pkgPath, workingDir) {
		pkgPath = strings.TrimPrefix(pkgPath, workingDir)
	}

	fullPath := filepath.Join(workingDir, pkgPath)
	if pkg, ok := p.parsedPackages[fullPath]; ok {
		logs.Infow("skip when already parsed the package", "plugin", p.Name, "method", "ParsePackagePath", "fullPath", fullPath)
		return pkg, nil
	}

	// parse the mojo package
	pkg, err := p.parsePackageFile(ctx, fullPath)
	if err != nil {
		return nil, err
	}

	pkg.SetExtraString("path", pkgPath)
	pkg.SetExtraString("workingDir", workingDir)

	// parse the dependency
	includedMojoPkg := false
	for name, d := range pkg.Dependencies {
		if strings.HasPrefix(name, "mojo.") {
			depPkg := GetMojoPackage(name)
			if depPkg == nil {
				return nil, fmt.Errorf("failed to found the required package %s", name)
			}
			includedMojoPkg = true
			pkg.ResolvedDependencies[depPkg.FullName] = depPkg
			continue
		}

		depPath := d.Path
		if len(depPath) == 0 {
			depPath, err = GetPackageCenter().Get(name, d)
			if err != nil {
				return nil, err
			}
		}

		depPath = util.GetAbsolutePath(filepath.Join(workingDir, pkgPath), depPath)
		depPkg, err := p.ParsePath(ctx, depPath)
		if err != nil {
			return nil, err
		}

		pkg.ResolvedDependencies[depPkg.FullName] = depPkg
	}

	// add mojo default packages
	if !strings.HasPrefix(pkg.FullName, "mojo.") {
		if pkg.ResolvedDependencies == nil {
			pkg.ResolvedDependencies = make(map[string]*lang.Package)
		}

		if includedMojoPkg {
			if _, ok := pkg.ResolvedDependencies["mojo.core"]; !ok {
				corePkg := GetMojoPackage("mojo.core")
				pkg.ResolvedDependencies[corePkg.FullName] = corePkg
			}
		} else {
			mojoPkgs := GetMojoPackages()
			for _, mp := range mojoPkgs {
				pkg.ResolvedDependencies[mp.FullName] = mp
			}
		}
	}

	p.parsedPackages[fullPath] = pkg
	return pkg, nil
}

func (p *DependencyParser) parsePackageFile(ctx context.Context, pkgPath string) (*lang.Package, error) {
	plugins := plugin.NewPlugins("syntax")
	packageFile := path.Join(pkgPath, "package.mojo")
	file, err := plugins.ParseFile(ctx, packageFile)
	if err != nil {
		logs.Errorw("failed to parse package file", "file", packageFile, "error", err.Error())
		return nil, err
	}

	if len(file.Statements) == 0 {
		logs.Errorw("not a valid package.mojo file, has no statement include", "file", file.FullName)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	decl := file.Statements[0].GetDeclaration()
	if decl == nil {
		logs.Errorw("not a valid package.mojo file, has no declaration statement include", "file", file.FullName)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	pkgDecl := decl.GetPackageDecl()
	if pkgDecl == nil {
		logs.Errorw("not a valid package.mojo file, has no declaration statement include", "file", file.FullName)
		return nil, errors.New("there is no package declaration in mojo file")
	}

	pkg := pkgDecl.Package
	if len(pkg.Dependencies) > 0 {
		pkg.ResolvedDependencies = make(map[string]*lang.Package)
	}

	return pkgDecl.Package, nil
}
