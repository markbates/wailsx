package windowx

import (
	"context"

	"github.com/markbates/safe"
)

var _ Reloader = Reload{}

type Reload struct {
	WindowReloadAppFn func(ctx context.Context) error
	WindowReloadFn    func(ctx context.Context) error
}

func (r Reload) WindowReload(ctx context.Context) error {
	return safe.Run(func() error {
		if r.WindowReloadFn == nil {
			return r.WindowReloadAppFn(ctx)
		}
		return r.WindowReloadFn(ctx)
	})
}

func (r Reload) WindowReloadApp(ctx context.Context) error {
	return safe.Run(func() error {
		return r.WindowReloadAppFn(ctx)
	})
}
