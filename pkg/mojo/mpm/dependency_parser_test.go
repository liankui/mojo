package mpm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/pkg/context"
	_ "github.com/mojo-lang/mojo/pkg/mojo/compiler"
	_ "github.com/mojo-lang/mojo/pkg/mojo/parser"
	"github.com/mojo-lang/mojo/pkg/plugin"
)

func TestDependencyParser_ParsePath(t *testing.T) {
	plugins := plugin.NewPlugins("mpm", "syntax", "semantic", "compiler")
	pkg, err := plugins.ParsePath(context.Empty(), "../testdata/mojo-alias")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}
