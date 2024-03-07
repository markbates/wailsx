package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_MaximiserManager_WindowFullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowFullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowFullscreen")))

	var called bool
	mm.WindowFullscreenFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowFullscreen(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.True(sd.Data.IsFullscreen)
	r.False(sd.Data.IsMaximised)
	r.False(sd.Data.IsMinimised)
	r.False(sd.Data.IsNormal)

	mm.WindowFullscreenFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowFullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

}

func Test_MaximiserManager_WindowIsFullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	_, err := mm.WindowIsFullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowIsFullscreen")))

	var called bool
	mm.WindowIsFullscreenFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	b, err := mm.WindowIsFullscreen(ctx)
	r.NoError(err)
	r.True(called)
	r.True(b)

	mm.WindowIsFullscreenFn = func(ctx context.Context) (bool, error) {
		return false, wailstest.ErrTest
	}

	b, err = mm.WindowIsFullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
	r.False(b)

}

func Test_MaximiserManager_WindowIsMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	_, err := mm.WindowIsMaximised(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowIsMaximised")))

	var called bool
	mm.WindowIsMaximisedFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	b, err := mm.WindowIsMaximised(ctx)
	r.NoError(err)
	r.True(called)
	r.True(b)

	mm.WindowIsMaximisedFn = func(ctx context.Context) (bool, error) {
		return false, wailstest.ErrTest
	}

	b, err = mm.WindowIsMaximised(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
	r.False(b)
}

func Test_MaximiserManager_WindowIsMinimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	_, err := mm.WindowIsMinimised(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowIsMinimised")))

	var called bool
	mm.WindowIsMinimisedFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	b, err := mm.WindowIsMinimised(ctx)
	r.NoError(err)
	r.True(called)
	r.True(b)

	mm.WindowIsMinimisedFn = func(ctx context.Context) (bool, error) {
		return false, wailstest.ErrTest
	}

	b, err = mm.WindowIsMinimised(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
	r.False(b)
}

func Test_MaximiserManager_WindowMaximise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowMaximise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowMaximise")))

	var called bool
	mm.WindowMaximiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowMaximise(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.False(sd.Data.IsFullscreen)
	r.True(sd.Data.IsMaximised)
	r.False(sd.Data.IsMinimised)
	r.False(sd.Data.IsNormal)

	mm.WindowMaximiseFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowMaximise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_MaximiserManager_WindowMinimise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowMinimise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowMinimise")))

	var called bool
	mm.WindowMinimiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowMinimise(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.False(sd.Data.IsFullscreen)
	r.False(sd.Data.IsMaximised)
	r.True(sd.Data.IsMinimised)
	r.False(sd.Data.IsNormal)

	mm.WindowMinimiseFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowMinimise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_MaximiserManager_WindowUnfullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowUnfullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowUnfullscreen")))

	var called bool
	mm.WindowUnfullscreenFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowUnfullscreen(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.False(sd.Data.IsFullscreen)
	r.False(sd.Data.IsMaximised)
	r.False(sd.Data.IsMinimised)
	r.True(sd.Data.IsNormal)

	mm.WindowUnfullscreenFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowUnfullscreen(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_MaximiserManager_WindowUnmaximise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowUnmaximise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowUnmaximise")))

	var called bool
	mm.WindowUnmaximiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowUnmaximise(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.False(sd.Data.IsFullscreen)
	r.False(sd.Data.IsMaximised)
	r.False(sd.Data.IsMinimised)
	r.True(sd.Data.IsNormal)

	mm.WindowUnmaximiseFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowUnmaximise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_MaximiserManager_WindowUnminimise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	err := mm.WindowUnminimise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowUnminimise")))

	var called bool
	mm.WindowUnminimiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = mm.WindowUnminimise(ctx)
	r.NoError(err)
	r.True(called)

	sd, err := mm.StateData(ctx)
	r.NoError(err)
	r.False(sd.Data.IsFullscreen)
	r.False(sd.Data.IsMaximised)
	r.False(sd.Data.IsMinimised)
	r.True(sd.Data.IsNormal)

	mm.WindowUnminimiseFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = mm.WindowUnminimise(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_MaximiserManager_WindowIsNormal(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := Maximiser{}

	ctx := context.Background()

	_, err := mm.WindowIsNormal(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowIsNormal")))

	var called bool
	mm.WindowIsNormalFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	b, err := mm.WindowIsNormal(ctx)
	r.NoError(err)
	r.True(called)
	r.True(b)

	mm.WindowIsNormalFn = func(ctx context.Context) (bool, error) {
		return false, wailstest.ErrTest
	}

	b, err = mm.WindowIsNormal(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
	r.False(b)
}

func Test_MaximiseManager_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &Maximiser{}

	sd, err := mm.StateData(context.Background())
	r.NoError(err)
	r.NotNil(sd)
	r.Equal("maximiser", sd.Name)

	mj := sd.Data

	r.False(mj.IsFullscreen)
	r.False(mj.IsMaximised)
	r.False(mj.IsMinimised)
	r.False(mj.IsNormal)

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
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			mm := Maximiser{
				WindowToggleMaximiseFn: tc.fn,
			}

			r.False(mm.data.IsMaximised)

			err := mm.WindowToggleMaximise(ctx)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.True(mm.data.IsMaximised)

		})
	}

	t.Run("toggle data", func(t *testing.T) {
		r := require.New(t)

		mm := Maximiser{
			WindowToggleMaximiseFn: func(ctx context.Context) error {
				return nil
			},
		}

		r.False(mm.data.IsMaximised)
		r.False(mm.data.IsMinimised)

		err := mm.WindowToggleMaximise(ctx)
		r.NoError(err)

		r.True(mm.data.IsMaximised)
		r.False(mm.data.IsMinimised)

		err = mm.WindowToggleMaximise(ctx)
		r.NoError(err)

		r.False(mm.data.IsMaximised)
		r.False(mm.data.IsMinimised)

	})

}
