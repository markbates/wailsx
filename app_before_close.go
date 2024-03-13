package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
)

func (app *App) BeforeClose(ctx context.Context) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	var wg safe.Group

	if app.BeforeCloseFn != nil {
		wg.Go(func() error {
			return app.BeforeCloseFn(ctx)
		})
	}

	beforeclosers := plugins.ByType[BeforeCloser](app.Plugins)

	for _, p := range beforeclosers {
		wg.Go(func() error {
			return p.BeforeClose(ctx)
		})
	}

	return wg.Wait()
}
