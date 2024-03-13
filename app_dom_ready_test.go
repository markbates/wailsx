package wailsx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_App_DomReady(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	r.Error(app.DomReady(ctx))

	app, err := NopApp("test")
	r.NoError(err)

	var called bool
	app.Plugins = append(app.Plugins, DomReadyerFn(func(ctx context.Context) error {
		called = true
		return nil
	}))

	r.NoError(app.DomReady(ctx))
	r.True(called)

}
