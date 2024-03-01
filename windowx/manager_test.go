package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_ScreenGetAll(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := Manager{}

	wm.ScreenGetAllFn = func(ctx context.Context) ([]wailsrun.Screen, error) {
		return []wailsrun.Screen{
			{
				Width: 100,
			},
		}, nil
	}

	ctx := context.Background()

	screens, err := wm.ScreenGetAll(ctx)
	r.NoError(err)
	r.Len(screens, 1)
	r.Equal(100, screens[0].Width)

	wm.ScreenGetAllFn = func(ctx context.Context) ([]wailsrun.Screen, error) {
		return nil, wailstest.ErrTest
	}

	_, err = wm.ScreenGetAll(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Manager_WindowExecJS(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := Manager{}

	var act string
	wm.WindowExecJSFn = func(ctx context.Context, js string) error {
		if len(js) == 0 {
			return wailstest.ErrTest
		}
		act = js
		return nil
	}

	ctx := context.Background()

	exp := "alert('hello')"
	err := wm.WindowExecJS(ctx, exp)
	r.NoError(err)
	r.Equal(exp, act)

	err = wm.WindowExecJS(ctx, "")
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Manager_WindowPrint(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := Manager{}

	var called bool
	wm.WindowPrintFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := wm.WindowPrint(ctx)
	r.NoError(err)
	r.True(called)

	wm.WindowPrintFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = wm.WindowPrint(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Manager_WindowSetAlwaysOnTop(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := Manager{}

	var act bool
	wm.WindowSetAlwaysOnTopFn = func(ctx context.Context, b bool) error {
		act = b
		return nil
	}

	ctx := context.Background()

	err := wm.WindowSetAlwaysOnTop(ctx, true)
	r.NoError(err)
	r.True(act)

	wm.WindowSetAlwaysOnTopFn = func(ctx context.Context, b bool) error {
		return wailstest.ErrTest
	}

	err = wm.WindowSetAlwaysOnTop(ctx, true)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_Manager_WindowSetTitle(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := Manager{}

	var act string
	wm.WindowSetTitleFn = func(ctx context.Context, title string) error {
		act = title
		return nil
	}

	ctx := context.Background()

	exp := "hello"

	err := wm.WindowSetTitle(ctx, exp)
	r.NoError(err)
	r.Equal(exp, act)

	wm.WindowSetTitleFn = func(ctx context.Context, title string) error {
		return wailstest.ErrTest
	}

	err = wm.WindowSetTitle(ctx, exp)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}
