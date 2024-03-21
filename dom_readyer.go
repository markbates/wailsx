package wailsx

import (
	"context"
	"fmt"
)

type DomReadyer interface {
	DomReady(ctx context.Context) error
}

var _ DomReadyer = DomReadyerFn(nil)

type DomReadyerFn func(ctx context.Context) error

func (f DomReadyerFn) DomReady(ctx context.Context) error {
	return f(ctx)
}

func (f DomReadyerFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
