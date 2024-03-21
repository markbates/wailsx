package wailsx

import (
	"context"
	"fmt"
)

type APIStateDataProvider interface {
	StateData(ctx context.Context) (*APIData, error)
}

type APIStateDataProviderFn func(ctx context.Context) (*APIData, error)

func (f APIStateDataProviderFn) StateData(ctx context.Context) (*APIData, error) {
	return f(ctx)
}

func (f APIStateDataProviderFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
