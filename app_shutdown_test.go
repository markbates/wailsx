package wailsx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_App_Shutdown(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	r.Error(app.Shutdown(ctx))

	app, err := NopApp("test")
	r.NoError(err)

	var called bool
	app.Plugins = append(app.Plugins, ShutdownerFn(func(ctx context.Context) error {
		called = true
		return nil
	}))

	r.NoError(app.Shutdown(ctx))
	r.True(called)
}
