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

	ctx := context.Background()

	_, err := wm.ScreenGetAll(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("ScreenGetAll")))

	wm.ScreenGetAllFn = func(ctx context.Context) ([]Screen, error) {
		return []Screen{
			{
				Size: ScreenSize{
					Width: 100,
				},
			},
		}, nil
	}

	screens, err := wm.ScreenGetAll(ctx)
	r.NoError(err)
	r.Len(screens, 1)

	screen := screens[0]
	w := screen.Size.Width
	r.Equal(100, w)

	wm.ScreenGetAllFn = func(ctx context.Context) ([]Screen, error) {
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

	ctx := context.Background()

	err := wm.WindowExecJS(ctx, "")
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowExecJS")))

	var act string
	wm.WindowExecJSFn = func(ctx context.Context, js string) error {
		if len(js) == 0 {
			return wailstest.ErrTest
		}
		act = js
		return nil
	}

	exp := "alert('hello')"
	err = wm.WindowExecJS(ctx, exp)
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

	ctx := context.Background()
	err := wm.WindowPrint(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowPrint")))

	var called bool
	wm.WindowPrintFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = wm.WindowPrint(ctx)
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
	ctx := context.Background()

	err := wm.WindowSetAlwaysOnTop(ctx, true)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetAlwaysOnTop")))

	var act bool
	wm.WindowSetAlwaysOnTopFn = func(ctx context.Context, b bool) error {
		act = b
		return nil
	}

	err = wm.WindowSetAlwaysOnTop(ctx, true)
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

	ctx := context.Background()

	err := wm.WindowSetTitle(ctx, "")
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetTitle")))

	var act string
	wm.WindowSetTitleFn = func(ctx context.Context, title string) error {
		act = title
		return nil
	}

	exp := "hello"

	err = wm.WindowSetTitle(ctx, exp)
	r.NoError(err)
	r.Equal(exp, act)

	wm.WindowSetTitleFn = func(ctx context.Context, title string) error {
		return wailstest.ErrTest
	}

	err = wm.WindowSetTitle(ctx, exp)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}
