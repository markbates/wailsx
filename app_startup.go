package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
)

func (app *App) Startup(ctx context.Context) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	var wg safe.Group

	if app.StartupFn != nil {
		wg.Go(func() error {
			return app.StartupFn(ctx)
		})
	}

	startupers := plugins.ByType[Startuper](app.Plugins)

	for _, p := range startupers {
		wg.Go(func() error {
			return p.Startup(ctx)
		})
	}

	return wg.Wait()
}
