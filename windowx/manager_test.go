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

	_, err := wm.StateData(ctx)
	r.NoError(err)

	wm = &Manager{}

	sd, err := wm.StateData(ctx)
	r.NoError(err)

	r.Nil(sd.MaximiserData)
	r.Nil(sd.PositionData)
	r.Nil(sd.ThemeData)

	wm = NopManager()

	sd, err = wm.StateData(ctx)
	r.NoError(err)

	r.NotNil(sd.MaximiserData)
	r.NotNil(sd.PositionData)
	r.NotNil(sd.ThemeData)

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

var _ RestorablePositionManager = &restoreablePositioner{}

type restoreablePositioner struct {
	*Positioner
	Data *PositionData
}

func (pm *restoreablePositioner) RestorePosition(ctx context.Context, data *PositionData) error {
	pm.Data = data
	return nil
}

var _ RestorableMaximiseManager = &restoreableMaximiser{}

type restoreableMaximiser struct {
	*Maximiser
	Data *MaximiserData
}

func (mm *restoreableMaximiser) RestoreMaximiser(ctx context.Context, data *MaximiserData) error {
	mm.Data = data
	return nil
}

var _ RestorableThemeManager = &restoreableThemer{}

type restoreableThemer struct {
	*Themer
	Data *ThemeData
}

func (tm *restoreableThemer) RestoreTheme(ctx context.Context, data *ThemeData) error {
	tm.Data = data
	return nil
}

func Test_Manager_RestoreWindows(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var wm *Manager

	data := &WindowData{
		MaximiserData: &MaximiserData{},
		PositionData:  &PositionData{},
		ThemeData:     &ThemeData{},
	}

	err := wm.RestoreWindows(ctx, data)
	r.Error(err)

	wm = NopManager()

	err = wm.RestoreWindows(ctx, nil)
	r.Error(err)

	err = wm.RestoreWindows(ctx, data)
	r.NoError(err)

	data = &WindowData{
		MaximiserData: &MaximiserData{
			Layout: WINDOW_MAXIMISED,
		},
		PositionData: &PositionData{
			X:          1,
			Y:          2,
			W:          3,
			H:          4,
			MaxW:       5,
			MaxH:       6,
			MinW:       7,
			MinH:       8,
			IsCentered: true,
		},
		ThemeData: &ThemeData{
			Theme: THEME_LIGHT,
			BackgroundColour: Colour{
				R: 10,
				G: 11,
				B: 12,
				A: 13,
			},
		},
	}

	pm := &restoreablePositioner{}
	wm.PositionManager = pm

	mm := &restoreableMaximiser{}
	wm.MaximiseManager = mm

	tm := &restoreableThemer{}
	wm.ThemeManager = tm

	err = wm.RestoreWindows(ctx, data)
	r.NoError(err)

	r.Equal(mm.Data, data.MaximiserData)
	r.Equal(pm.Data, data.PositionData)
	r.Equal(tm.Data, data.ThemeData)

}
