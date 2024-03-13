package wailsx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_App_Save(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	r.Error(app.Save(ctx))

	app, err := NopApp("test")
	r.NoError(err)

	var called bool
	app.Plugins = append(app.Plugins, SaverFn(func(ctx context.Context) error {
		called = true
		return nil
	}))

	r.NoError(app.Save(ctx))
	r.True(called)

}
