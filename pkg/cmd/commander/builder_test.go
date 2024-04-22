package commander

import (
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/plugin"
)

var b *Builder

func init() {
	b = &Builder{
		Package:              nil,
		Files:                nil,
		OpenAPIs:             nil,
		PackageName:          "",
		Targets:              "",
		Engine:               "gokit",
		Output:               "",
		Pwd:                  "./testdata/data",
		Path:                 "",
		APIEnabled:           true,
		NcraftAllEnabled:     false,
		NcraftServiceEnabled: true,
		NcraftClientEnabled:  false,
		NcraftSidecarEnabled: false,
		Repository:           "",
	}

	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")

	if strings.HasPrefix(b.Path, b.Pwd) {
		b.Path = strings.TrimPrefix(b.Path, b.Pwd)
	}
	pkg, err := plugins.ParsePath(context.Empty(), path.Join(b.Pwd, b.Path))
	if err != nil {
		log.Fatalf("failed to parsePath, error=%v", err)
	}
	b.Package = pkg
}

func clearServiceGoFiles() error {
	if err := os.RemoveAll(path.Join(b.Pwd, "service-go")); err != nil {
		log.Printf("failed to clear service-go files, error=%v\n", err)
		return err
	}
	return nil
}

func Test_clearFiles(t *testing.T) {
	if err := clearServiceGoFiles(); err != nil {
		t.Errorf("test errror=%v", err)
	}
}

func TestBuilder_buildGokit(t *testing.T) {
	if err := b.buildGokit("service"); err != nil {
		t.Errorf("buildGokit() error = %v", err)
	}
}
