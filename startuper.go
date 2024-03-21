package wailsx

import (
	"context"
	"fmt"
)

type Startuper interface {
	Startup(ctx context.Context) error
}

var _ Startuper = StartuperFn(nil)

type StartuperFn func(ctx context.Context) error

func (f StartuperFn) Startup(ctx context.Context) error {
	return f(ctx)
}

func (f StartuperFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
