package windowx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_ScreenGetAll(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	exp := []Screen{
		{
			Size: ScreenSize{
				Width: 100,
			},
		},
	}

	tcs := []struct {
		name string
		fn   func(ctx context.Context) ([]Screen, error)
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) ([]Screen, error) {
				return exp, nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context) ([]Screen, error) {
				return nil, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) ([]Screen, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("ScreenGetAll"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			wm := Manager{
				ScreenGetAllFn: tc.fn,
			}

			act, err := wm.ScreenGetAll(ctx)
			r.Equal(tc.err, err)

			if tc.err == nil {
				r.Equal(exp, act)
			}
		})
	}

}

func Test_Manager_WindowExecJS(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, js string) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, js string) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, js string) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, js string) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowExecJS"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			wm := Manager{
				WindowExecJSFn: tc.fn,
			}

			err := wm.WindowExecJS(ctx, "")
			r.Equal(tc.err, err)
		})
	}

}

func Test_Manager_WindowPrint(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowPrint"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			wm := Manager{
				WindowPrintFn: tc.fn,
			}

			err := wm.WindowPrint(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Manager_WindowSetAlwaysOnTop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, b bool) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, b bool) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, b bool) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, b bool) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetAlwaysOnTop"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			wm := Manager{
				WindowSetAlwaysOnTopFn: tc.fn,
			}

			err := wm.WindowSetAlwaysOnTop(ctx, false)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Manager_WindowSetTitle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, title string) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, title string) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, title string) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, title string) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetTitle"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			wm := Manager{
				WindowSetTitleFn: tc.fn,
			}

			err := wm.WindowSetTitle(ctx, "")
			r.Equal(tc.err, err)
		})
	}
}

func Test_Manager_StateData(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	var wm *Manager

	sd, err := wm.StateData(ctx)
	r.NoError(err)
	r.Equal(ManagerStateDataName, sd.Name)

	wm = &Manager{}

	sd, err = wm.StateData(ctx)
	r.NoError(err)
	r.Equal(ManagerStateDataName, sd.Name)

	r.Nil(sd.Data.MaximiserData)
	r.Nil(sd.Data.PositionData)
	r.Nil(sd.Data.ThemeData)

	wm = NopManager()

	sd, err = wm.StateData(ctx)
	r.NoError(err)
	r.Equal(ManagerStateDataName, sd.Name)

	r.NotNil(sd.Data.MaximiserData)
	r.NotNil(sd.Data.PositionData)
	r.NotNil(sd.Data.ThemeData)

}

func Test_NopManager(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := NopManager()
	r.NotNil(wm)
	r.NotNil(wm.MaximiseManager)
	r.NotNil(wm.PositionManager)
	r.NotNil(wm.ReloadManager)
	r.NotNil(wm.ThemeManager)
	r.NotNil(wm.Toggler)
	r.NotNil(wm.ScreenGetAllFn)
	r.NotNil(wm.WindowExecJSFn)
	r.NotNil(wm.WindowPrintFn)
	r.NotNil(wm.WindowSetAlwaysOnTopFn)
	r.NotNil(wm.WindowSetTitleFn)
}
