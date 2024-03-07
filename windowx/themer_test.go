package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Themer_WindowSetDarkTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tm := Themer{}

	ctx := context.Background()
	err := tm.WindowSetDarkTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetDarkTheme")))

	var called bool
	tm.WindowSetDarkThemeFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = tm.WindowSetDarkTheme(ctx)
	r.NoError(err)
	r.True(called)

	tm.WindowSetDarkThemeFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = tm.WindowSetDarkTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Themer_WindowSetLightTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tm := Themer{}

	ctx := context.Background()
	err := tm.WindowSetLightTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetLightTheme")))

	var called bool
	tm.WindowSetLightThemeFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = tm.WindowSetLightTheme(ctx)
	r.NoError(err)
	r.True(called)

	tm.WindowSetLightThemeFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = tm.WindowSetLightTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Themer_WindowSetSystemDefaultTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tm := Themer{}

	ctx := context.Background()
	err := tm.WindowSetSystemDefaultTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetSystemDefaultTheme")))

	var called bool
	tm.WindowSetSystemDefaultThemeFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = tm.WindowSetSystemDefaultTheme(ctx)
	r.NoError(err)
	r.True(called)

	tm.WindowSetSystemDefaultThemeFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = tm.WindowSetSystemDefaultTheme(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Themer_WindowSetBackgroundColour(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tm := Themer{}

	ctx := context.Background()
	err := tm.WindowSetBackgroundColour(ctx, 0, 0, 0, 0)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetBackgroundColour")))

	var called bool
	tm.WindowSetBackgroundColourFn = func(ctx context.Context, R, G, B, A uint8) error {
		called = true
		return nil
	}

	err = tm.WindowSetBackgroundColour(ctx, 0, 0, 0, 0)
	r.NoError(err)
	r.True(called)

	tm.WindowSetBackgroundColourFn = func(ctx context.Context, R, G, B, A uint8) error {
		return wailstest.ErrTest
	}

	err = tm.WindowSetBackgroundColour(ctx, 0, 0, 0, 0)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}
