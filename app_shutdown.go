package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
)

func (app *App) Shutdown(ctx context.Context) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	var wg safe.Group

	wg.Go(func() error {
		return app.Save(ctx)
	})

	if app.ShutdownFn != nil {
		wg.Go(func() error {
			return app.ShutdownFn(ctx)
		})
	}

	shutdowners := plugins.ByType[Shutdowner](app.Plugins)

	for _, p := range shutdowners {
		wg.Go(func() error {
			return p.Shutdown(ctx)
		})
	}

	return wg.Wait()
}
