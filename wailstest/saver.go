package wailstest

import (
	"context"
	"fmt"
)

type Saver func(ctx context.Context) error

func (s Saver) Save(ctx context.Context) error {
	return s(ctx)
}

func (s Saver) PluginName() string {
	return fmt.Sprintf("%T", s)
}
