package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_MaximiserManager_WindowFullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := MaximiserManager{}

	var called bool
	mm.WindowFullscreenFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowFullscreen(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowIsFullscreenFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	ctx := context.Background()

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowIsMaximisedFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	ctx := context.Background()

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowIsMinimisedFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	ctx := context.Background()

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowMaximiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowMaximise(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowMinimiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowMinimise(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowUnfullscreenFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowUnfullscreen(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowUnmaximiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowUnmaximise(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowUnminimiseFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	ctx := context.Background()

	err := mm.WindowUnminimise(ctx)
	r.NoError(err)
	r.True(called)

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

	mm := MaximiserManager{}

	var called bool
	mm.WindowIsNormalFn = func(ctx context.Context) (bool, error) {
		called = true
		return true, nil
	}

	ctx := context.Background()

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
