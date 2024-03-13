package wailsx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_App_BeforeClose(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	r.Error(app.BeforeClose(ctx))

	app, err := NopApp("test")
	r.NoError(err)

	var called bool
	app.Plugins = append(app.Plugins, wailstest.BeforeCloser(func(ctx context.Context) error {
		called = true
		return nil
	}))

	r.NoError(app.BeforeClose(ctx))
	r.True(called)
}
