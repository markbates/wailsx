package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
)

func (app *App) DomReady(ctx context.Context) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	var wg safe.Group

	if app.DomReadyFn != nil {
		wg.Go(func() error {
			return app.DomReadyFn(ctx)
		})
	}

	domreadyers := plugins.ByType[DomReadyer](app.Plugins)

	for _, p := range domreadyers {
		wg.Go(func() error {
			return p.DomReady(ctx)
		})
	}

	return wg.Wait()
}
