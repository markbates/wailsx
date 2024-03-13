package wailstest

import (
	"context"
	"fmt"
)

type DomReadyer func(ctx context.Context) error

func (dr DomReadyer) DomReady(ctx context.Context) error {
	return dr(ctx)
}

func (dr DomReadyer) PluginName() string {
	return fmt.Sprintf("%T", dr)
}
