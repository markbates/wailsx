package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Reload_WindowReload(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	rd := Reloader{}

	ctx := context.Background()

	var called bool
	rd.WindowReloadFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err := rd.WindowReload(ctx)
	r.NoError(err)
	r.True(called)

	rd.WindowReloadFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = rd.WindowReload(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	rd.WindowReloadFn = nil
	err = rd.WindowReload(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowReload")))
}

func Test_Reload_WindowReloadApp(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	rd := Reloader{}

	ctx := context.Background()

	var called bool
	rd.WindowReloadAppFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err := rd.WindowReloadApp(ctx)
	r.NoError(err)
	r.True(called)

	rd.WindowReloadAppFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = rd.WindowReloadApp(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	rd.WindowReloadAppFn = nil
	err = rd.WindowReloadApp(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowReloadApp")))
}
