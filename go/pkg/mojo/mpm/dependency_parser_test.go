package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mojo-lang/mojo/go/pkg/mojo/compiler"
	"github.com/mojo-lang/mojo/go/pkg/mojo/context"
	_ "github.com/mojo-lang/mojo/go/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
	"github.com/mojo-lang/mojo/go/pkg/mojo/testdata"
)

func TestDependencyParser_ParsePackagePath(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePackagePath(context.Empty(), "mojo-alias", testdata.AliasCaseFiles)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
