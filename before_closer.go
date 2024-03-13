package wailsx

import (
	"context"
	"fmt"
)

type BeforeCloser interface {
	BeforeClose(ctx context.Context) error
}

type BeforeCloserFn func(ctx context.Context) error

func (f BeforeCloserFn) BeforeClose(ctx context.Context) error {
	return f(ctx)
}

func (f BeforeCloserFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
