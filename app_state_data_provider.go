package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

const AppStateDataProviderName = "app"

type AppStateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[AppData], error)
}

type AppStateDataProviderFn func(ctx context.Context) (statedata.Data[AppData], error)

func (f AppStateDataProviderFn) StateData(ctx context.Context) (statedata.Data[AppData], error) {
	return f(ctx)
}

func (f AppStateDataProviderFn) PluginName() string {
	return AppStateDataProviderName
}
