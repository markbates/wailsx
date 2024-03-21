package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/statedata"
)

type AppStateDataProvider interface {
	plugins.Plugin
	StateData(ctx context.Context) (statedata.Data[AppData], error)
}

type AppStateDataProviderFn func(ctx context.Context) (statedata.Data[AppData], error)

func (f AppStateDataProviderFn) StateData(ctx context.Context) (statedata.Data[AppData], error) {
	return f(ctx)
}

func (f AppStateDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
