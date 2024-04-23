package converter

import (
	"sync"

	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/pkg/context"
)

type NominalPlugin interface {
	Name() string
	ConvertTo(ctx context.Context, t *lang.NominalType) (string, string, error)
}

var (
	once           sync.Once
	nominalPlugins map[string][]NominalPlugin
)

func RegisterNominalPlugin(plugin NominalPlugin) {
	once.Do(func() {
		nominalPlugins = make(map[string][]NominalPlugin)
	})

	if plugin != nil {
		pgs := nominalPlugins[plugin.Name()]
		for _, p := range pgs {
			if p.Name() == plugin.Name() {
				logs.Warnw("the plugin already registered", "name", plugin.Name())
				return
			}
		}
		pgs = append(pgs, plugin)
		nominalPlugins[plugin.Name()] = pgs
	}
}

func NominalPlugins() map[string][]NominalPlugin {
	once.Do(func() {
		nominalPlugins = make(map[string][]NominalPlugin)
	})
	return nominalPlugins
}
