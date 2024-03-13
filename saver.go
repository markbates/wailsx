package wailsx

import (
	"context"
	"fmt"
)

type Saver interface {
	Save(ctx context.Context) error
}

type SaverFn func(ctx context.Context) error

func (f SaverFn) Save(ctx context.Context) error {
	return f(ctx)
}

func (f SaverFn) PluginName() string {
	return fmt.Sprintf("%T", f)
}
