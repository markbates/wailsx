package wailsx

import (
	"context"
	"fmt"
)

type Shutdowner interface {
	Shutdown(ctx context.Context) error
}

var _ Shutdowner = ShutdownerFn(nil)

type ShutdownerFn func(ctx context.Context) error

func (f ShutdownerFn) Shutdown(ctx context.Context) error {
	return f(ctx)
}

func (f ShutdownerFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
