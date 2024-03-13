package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
)

// Save the state of the application
func (app *App) Save(ctx context.Context) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	var wg safe.Group

	if app.SaveFn != nil {
		wg.Go(func() error {
			return app.SaveFn(ctx)
		})
	}

	savers := plugins.ByType[Saver](app.Plugins)

	for _, p := range savers {
		wg.Go(func() error {
			return p.Save(ctx)
		})

	}

	return wg.Wait()
}
