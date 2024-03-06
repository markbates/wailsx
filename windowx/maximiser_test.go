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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
	r.True(errors.Is(err, wailsrun.ErrNotAvailable))

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
