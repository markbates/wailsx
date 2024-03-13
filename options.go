package wailsx

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
)

func (app *App) Options() (*options.App, error) {
	opts := &options.App{
		Title:            app.Name,
		Width:            1024,
		Height:           768,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup: func(ctx context.Context) {
			if err := app.Startup(ctx); err != nil {
				println("Error: ", err.Error())
			}
		},
		OnShutdown: func(ctx context.Context) {
			if err := app.Shutdown(ctx); err != nil {
				println("Error: ", err.Error())
			}
		},
		OnDomReady: func(ctx context.Context) {
			if err := app.DomReady(ctx); err != nil {
				println("Error: ", err.Error())
			}
		},
		OnBeforeClose: func(ctx context.Context) bool {
			err := app.BeforeClose(ctx)
			return err != nil
		},
		Bind: []interface{}{
			app,
		},
	}
	return opts, nil
}
