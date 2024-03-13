package wailstest

import (
	"context"
	"fmt"
)

type Shutdowner func(ctx context.Context) error

func (sd Shutdowner) Shutdown(ctx context.Context) error {
	return sd(ctx)
}

func (sd Shutdowner) PluginName() string {
	return fmt.Sprintf("%T", sd)
}
