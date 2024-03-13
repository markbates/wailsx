package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

type PluginDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[any], error)
}

type PluginDataProviderFn func(ctx context.Context) (statedata.Data[any], error)

func (f PluginDataProviderFn) StateData(ctx context.Context) (statedata.Data[any], error) {
	return f(ctx)
}

func (f PluginDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
