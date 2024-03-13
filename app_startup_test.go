package wailsx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_App_Startup(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	r.Error(app.Startup(ctx))

	app, err := NopApp("test")
	r.NoError(err)

	var called bool
	app.Plugins = append(app.Plugins, wailstest.Startuper(func(ctx context.Context) error {
		called = true
		return nil
	}))

	r.NoError(app.Startup(ctx))
	r.True(called)
}
