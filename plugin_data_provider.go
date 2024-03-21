package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/statedata"
)

type PluginDataProvider interface {
	plugins.Plugin
	statedata.DataProvider[any]
}

var _ PluginDataProvider = PluginDataProviderFn(nil)

type PluginDataProviderFn func(ctx context.Context) (any, error)

func (f PluginDataProviderFn) StateData(ctx context.Context) (any, error) {
	return f(ctx)
}

func (f PluginDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}

type RestorablePlugin interface {
	plugins.Plugin
	RestorePlugin(ctx context.Context, data any) error
}
