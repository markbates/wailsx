package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
)

type AppStateDataProvider interface {
	plugins.Plugin
	StateData(ctx context.Context) (*AppData, error)
}

type AppStateDataProviderFn func(ctx context.Context) (*AppData, error)

func (f AppStateDataProviderFn) StateData(ctx context.Context) (*AppData, error) {
	return f(ctx)
}

func (f AppStateDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
