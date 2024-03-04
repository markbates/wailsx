package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ Reloader = Reload{}

type Reload struct {
	WindowReloadAppFn func(ctx context.Context) error
	WindowReloadFn    func(ctx context.Context) error
}

func (r Reload) WindowReload(ctx context.Context) error {
	return safe.Run(func() error {
		if r.WindowReloadFn != nil {
			return r.WindowReloadFn(ctx)
		}

		return wailsrun.WindowReload(ctx)
	})
}

func (r Reload) WindowReloadApp(ctx context.Context) error {
	return safe.Run(func() error {
		if r.WindowReloadAppFn != nil {
			return r.WindowReloadAppFn(ctx)
		}

		return wailsrun.WindowReloadApp(ctx)
	})
}
