package wailstest

import (
	"context"
	"fmt"
)

type Startuper func(ctx context.Context) error

func (s Startuper) Startup(ctx context.Context) error {
	return s(ctx)
}

func (s Startuper) PluginName() string {
	return fmt.Sprintf("%T", s)
}
