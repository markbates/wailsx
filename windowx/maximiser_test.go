package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Nil_Maximiser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var mm *Maximiser

	ctx := context.Background()

	err := mm.WindowFullscreen(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowFullscreen"), err)

	_, err = mm.WindowIsFullscreen(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowIsFullscreen"), err)

	_, err = mm.WindowIsMaximised(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowIsMaximised"), err)

	_, err = mm.WindowIsMinimised(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowIsMinimised"), err)

	err = mm.WindowMaximise(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowMaximise"), err)

	err = mm.WindowMinimise(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowMinimise"), err)

	err = mm.WindowUnfullscreen(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowUnfullscreen"), err)

	err = mm.WindowUnmaximise(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowUnmaximise"), err)

	err = mm.WindowUnminimise(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowUnminimise"), err)

	_, err = mm.WindowIsNormal(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowIsNormal"), err)

	err = mm.WindowToggleMaximise(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowToggleMaximise"), err)
}

func Test_Maximiser_WindowFullscreen(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowFullscreen"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowFullscreenFn: tc.fn,
			}

			err := mm.WindowFullscreen(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_FULLSCREEN, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowIsFullscreen(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (bool, error)
		exp  bool
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (bool, error) {
				return true, nil
			},
			exp: true,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowIsFullscreen"),
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (bool, error) {
				return false, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (bool, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowIsFullscreenFn: tc.fn,
			}

			is, err := mm.WindowIsFullscreen(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, is)
		})
	}
}

func Test_Maximiser_WindowIsMaximised(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (bool, error)
		exp  bool
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (bool, error) {
				return true, nil
			},
			exp: true,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowIsMaximised"),
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (bool, error) {
				return false, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (bool, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowIsMaximisedFn: tc.fn,
			}

			is, err := mm.WindowIsMaximised(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, is)
		})
	}
}

func Test_Maximiser_WindowIsMinimised(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (bool, error)
		exp  bool
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (bool, error) {
				return true, nil
			},
			exp: true,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowIsMinimised"),
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (bool, error) {
				return false, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (bool, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowIsMinimisedFn: tc.fn,
			}

			is, err := mm.WindowIsMinimised(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, is)
		})
	}
}

func Test_Maximiser_WindowMaximise(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowMaximise"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowMaximiseFn: tc.fn,
			}

			err := mm.WindowMaximise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_MAXIMISED, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowMinimise(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowMinimise"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowMinimiseFn: tc.fn,
			}

			err := mm.WindowMinimise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_MINIMISED, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowUnfullscreen(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowUnfullscreen"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowUnfullscreenFn: tc.fn,
			}

			err := mm.WindowUnfullscreen(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_NORMAL, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowUnmaximise(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowUnmaximise"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowUnmaximiseFn: tc.fn,
			}

			err := mm.WindowUnmaximise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_NORMAL, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowUnminimise(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowUnminimise"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowUnminimiseFn: tc.fn,
			}

			err := mm.WindowUnminimise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(WINDOW_NORMAL, mm.data.Layout)
		})
	}
}

func Test_Maximiser_WindowIsNormal(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (bool, error)
		exp  bool
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (bool, error) {
				return true, nil
			},
			exp: true,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowIsNormal"),
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (bool, error) {
				return false, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (bool, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowIsNormalFn: tc.fn,
			}

			is, err := mm.WindowIsNormal(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, is)
		})
	}
}

func Test_Maximiser_WindowToggleMaximise(t *testing.T) {
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
			name: "no function",
			err:  wailsrun.ErrNotAvailable("WindowToggleMaximise"),
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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowToggleMaximiseFn: tc.fn,
			}

			r.Equal(WINDOW_NORMAL, mm.data.Layout)

			err := mm.WindowToggleMaximise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)

			r.Equal(WINDOW_MAXIMISED, mm.data.Layout)
		})
	}

	t.Run("toggle data", func(t *testing.T) {
		r := require.New(t)

		mm := Maximiser{
			WindowToggleMaximiseFn: func(ctx context.Context) error {
				return nil
			},
		}

		r.Equal(WINDOW_NORMAL, mm.data.Layout)

		err := mm.WindowToggleMaximise(ctx)
		r.NoError(err)

		r.Equal(WINDOW_MAXIMISED, mm.data.Layout)

		err = mm.WindowToggleMaximise(ctx)
		r.NoError(err)

		r.Equal(WINDOW_NORMAL, mm.data.Layout)

	})
}

func Test_NopMaximiser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := NopMaximiser()
	r.NotNil(mm)
	r.NotNil(mm.WindowFullscreenFn)
	r.NotNil(mm.WindowIsFullscreenFn)
	r.NotNil(mm.WindowIsMaximisedFn)
	r.NotNil(mm.WindowIsMinimisedFn)
	r.NotNil(mm.WindowIsNormalFn)
	r.NotNil(mm.WindowMaximiseFn)
	r.NotNil(mm.WindowMinimiseFn)
	r.NotNil(mm.WindowToggleMaximiseFn)
	r.NotNil(mm.WindowUnfullscreenFn)
	r.NotNil(mm.WindowUnmaximiseFn)
	r.NotNil(mm.WindowUnminimiseFn)

}

func Test_Maximiser_RestoreMaximiser(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var mm *Maximiser

	data := &MaximiserData{}

	err := mm.RestoreMaximiser(ctx, data)
	r.Error(err)

	mm = NopMaximiser()

	err = mm.RestoreMaximiser(ctx, nil)
	r.Error(err)

	var layout string
	mm.WindowFullscreenFn = func(ctx context.Context) error {
		layout = WINDOW_FULLSCREEN
		return nil
	}

	mm.WindowMaximiseFn = func(ctx context.Context) error {
		layout = WINDOW_MAXIMISED
		return nil
	}

	mm.WindowMinimiseFn = func(ctx context.Context) error {
		layout = WINDOW_MINIMISED
		return nil
	}

	err = mm.RestoreMaximiser(ctx, data)
	r.NoError(err)
	r.Equal(WINDOW_NORMAL, layout)

	data.Layout = WINDOW_FULLSCREEN

	err = mm.RestoreMaximiser(ctx, data)
	r.NoError(err)
	r.Equal(WINDOW_FULLSCREEN, layout)

	data.Layout = WINDOW_MAXIMISED

	err = mm.RestoreMaximiser(ctx, data)
	r.NoError(err)
	r.Equal(WINDOW_MAXIMISED, layout)

	data.Layout = WINDOW_MINIMISED

	err = mm.RestoreMaximiser(ctx, data)
	r.NoError(err)
	r.Equal(WINDOW_MINIMISED, layout)

}
