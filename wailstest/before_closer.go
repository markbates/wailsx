package wailstest

import (
	"context"
	"fmt"
)

type BeforeCloser func(ctx context.Context) error

func (bc BeforeCloser) BeforeClose(ctx context.Context) error {
	return bc(ctx)
}

func (bc BeforeCloser) PluginName() string {
	return fmt.Sprintf("%T", bc)
}
