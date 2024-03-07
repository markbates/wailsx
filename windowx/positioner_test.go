package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_PositionManager_WindowCenter(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	err := pm.WindowCenter(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowCenter")))

	var called bool
	pm.WindowCenterFn = func(ctx context.Context) error {
		called = true
		return nil
	}

	err = pm.WindowCenter(ctx)
	r.NoError(err)
	r.True(called)

	pm.WindowCenterFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}

	err = pm.WindowCenter(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowGetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	_, _, err := pm.WindowGetPosition(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowGetPosition")))

	ex := 1
	ey := 2
	pm.WindowGetPositionFn = func(ctx context.Context) (int, int, error) {
		return ex, ey, nil
	}

	x, y, err := pm.WindowGetPosition(ctx)
	r.NoError(err)
	r.Equal(ex, x)
	r.Equal(ey, y)

	pm.WindowGetPositionFn = func(ctx context.Context) (int, int, error) {
		return 0, 0, wailstest.ErrTest
	}

	_, _, err = pm.WindowGetPosition(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowGetSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	_, _, err := pm.WindowGetSize(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowGetSize")))

	ew := 1
	eh := 2
	pm.WindowGetSizeFn = func(ctx context.Context) (int, int, error) {
		return ew, eh, nil
	}

	w, h, err := pm.WindowGetSize(ctx)
	r.NoError(err)
	r.Equal(ew, w)
	r.Equal(eh, h)

	pm.WindowGetSizeFn = func(ctx context.Context) (int, int, error) {
		return 0, 0, wailstest.ErrTest
	}

	_, _, err = pm.WindowGetSize(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowSetMaxSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	err := pm.WindowSetMaxSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetMaxSize")))

	var aw, ah int
	pm.WindowSetMaxSizeFn = func(ctx context.Context, width int, height int) error {
		aw = width
		ah = height
		return nil
	}

	err = pm.WindowSetMaxSize(ctx, 1, 2)
	r.NoError(err)
	r.Equal(1, aw)
	r.Equal(2, ah)

	pm.WindowSetMaxSizeFn = func(ctx context.Context, width int, height int) error {
		return wailstest.ErrTest
	}

	err = pm.WindowSetMaxSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowSetMinSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	err := pm.WindowSetMinSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetMinSize")))

	var aw, ah int
	pm.WindowSetMinSizeFn = func(ctx context.Context, width int, height int) error {
		aw = width
		ah = height
		return nil
	}

	err = pm.WindowSetMinSize(ctx, 1, 2)
	r.NoError(err)
	r.Equal(1, aw)
	r.Equal(2, ah)

	pm.WindowSetMinSizeFn = func(ctx context.Context, width int, height int) error {
		return wailstest.ErrTest
	}

	err = pm.WindowSetMinSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowSetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	err := pm.WindowSetPosition(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetPosition")))

	var ax, ay int
	pm.WindowSetPositionFn = func(ctx context.Context, x int, y int) error {
		ax = x
		ay = y
		return nil
	}

	err = pm.WindowSetPosition(ctx, 1, 2)
	r.NoError(err)
	r.Equal(1, ax)
	r.Equal(2, ay)

	pm.WindowSetPositionFn = func(ctx context.Context, x int, y int) error {
		return wailstest.ErrTest
	}

	err = pm.WindowSetPosition(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_PositionManager_WindowSetSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := Positioner{}

	ctx := context.Background()

	err := pm.WindowSetSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailsrun.ErrNotAvailable("WindowSetSize")))

	var aw, ah int
	pm.WindowSetSizeFn = func(ctx context.Context, width int, height int) error {
		aw = width
		ah = height
		return nil
	}

	err = pm.WindowSetSize(ctx, 1, 2)
	r.NoError(err)
	r.Equal(1, aw)
	r.Equal(2, ah)

	pm.WindowSetSizeFn = func(ctx context.Context, width int, height int) error {
		return wailstest.ErrTest
	}

	err = pm.WindowSetSize(ctx, 1, 2)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))
}
