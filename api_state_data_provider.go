package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

const APIStateDataProviderName = "api"

type APIStateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[*APIData], error)
}

type APIStateDataProviderFn func(ctx context.Context) (statedata.Data[*APIData], error)

func (f APIStateDataProviderFn) StateData(ctx context.Context) (statedata.Data[*APIData], error) {
	return f(ctx)
}

func (f APIStateDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
