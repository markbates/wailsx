package windowx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_PositionManager_WindowCenter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	err := pm.WindowCenter(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowCenter"), err)

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
			err:  wailsrun.ErrNotAvailable("WindowCenter"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowCenterFn: tc.fn,
			}

			err := pm.WindowCenter(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_PositionManager_WindowGetPosition(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	_, _, err := pm.WindowGetPosition(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowGetPosition"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (int, int, error)
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (int, int, error) {
				return 0, 0, nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (int, int, error) {
				return 0, 0, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (int, int, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowGetPosition"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowGetPositionFn: tc.fn,
			}

			x, y, err := pm.WindowGetPosition(ctx)
			r.Equal(tc.err, err)
			r.Zero(x)
			r.Zero(y)
		})
	}
}

func Test_PositionManager_WindowGetSize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	_, _, err := pm.WindowGetSize(ctx)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowGetSize"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (int, int, error)
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (int, int, error) {
				return 0, 0, nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (int, int, error) {
				return 0, 0, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (int, int, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowGetSize"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowGetSizeFn: tc.fn,
			}

			w, h, err := pm.WindowGetSize(ctx)
			r.Equal(tc.err, err)
			r.Zero(w)
			r.Zero(h)
		})
	}
}

func Test_PositionManager_WindowSetMaxSize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	err := pm.WindowSetMaxSize(ctx, 0, 0)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowSetMaxSize"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context, width int, height int) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, width int, height int) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, width int, height int) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, width int, height int) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetMaxSize"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowSetMaxSizeFn: tc.fn,
			}

			err := pm.WindowSetMaxSize(ctx, 0, 0)
			r.Equal(tc.err, err)
		})
	}
}

func Test_PositionManager_WindowSetMinSize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	err := pm.WindowSetMinSize(ctx, 0, 0)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowSetMinSize"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context, width int, height int) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, width int, height int) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, width int, height int) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, width int, height int) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetMinSize"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowSetMinSizeFn: tc.fn,
			}

			err := pm.WindowSetMinSize(ctx, 0, 0)
			r.Equal(tc.err, err)
		})
	}
}

func Test_PositionManager_WindowSetPosition(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	err := pm.WindowSetPosition(ctx, 0, 0)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowSetPosition"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context, x int, y int) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, x int, y int) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, x int, y int) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, x int, y int) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetPosition"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowSetPositionFn: tc.fn,
			}

			err := pm.WindowSetPosition(ctx, 0, 0)
			r.Equal(tc.err, err)
		})
	}
}

func Test_PositionManager_WindowSetSize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	r := require.New(t)

	var pm *Positioner
	err := pm.WindowSetSize(ctx, 0, 0)
	r.Error(err)
	r.Equal(wailsrun.ErrNotAvailable("WindowSetSize"), err)

	tcs := []struct {
		name string
		fn   func(ctx context.Context, width int, height int) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, width int, height int) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, width int, height int) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, width int, height int) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetSize"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			pm := &Positioner{
				WindowSetSizeFn: tc.fn,
			}

			err := pm.WindowSetSize(ctx, 0, 0)
			r.Equal(tc.err, err)
		})
	}
}

func Test_NopPositioner(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := NopPositioner()
	r.NotNil(pm)
	r.NotNil(pm.WindowCenterFn)
	r.NotNil(pm.WindowGetPositionFn)
	r.NotNil(pm.WindowGetSizeFn)
	r.NotNil(pm.WindowSetMaxSizeFn)
	r.NotNil(pm.WindowSetMinSizeFn)
	r.NotNil(pm.WindowSetPositionFn)
	r.NotNil(pm.WindowSetSizeFn)
}

func Test_Positioner_RestorePosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var pm *Positioner

	err := pm.RestorePosition(ctx, &PositionData{})
	r.Error(err)

	pm = NopPositioner()
	err = pm.RestorePosition(ctx, nil)
	r.Error(err)

	var ax, ay, aw, ah, maxw, maxh, minw, minh int

	pm.WindowSetPositionFn = func(ctx context.Context, x int, y int) error {
		ax = x
		ay = y
		return nil
	}

	pm.WindowSetSizeFn = func(ctx context.Context, w int, h int) error {
		aw = w
		ah = h
		return nil
	}

	pm.WindowSetMaxSizeFn = func(ctx context.Context, w int, h int) error {
		maxw = w
		maxh = h
		return nil
	}

	pm.WindowSetMinSizeFn = func(ctx context.Context, w int, h int) error {
		minw = w
		minh = h
		return nil
	}

	data := &PositionData{}

	err = pm.RestorePosition(ctx, data)
	r.NoError(err)

	r.Equal(pm.InitPosX(), ax)
	r.Equal(pm.InitPosY(), ay)
	r.Equal(pm.InitWidth(), aw)
	r.Equal(pm.InitHeight(), ah)
	r.Zero(maxw)
	r.Zero(maxh)
	r.Zero(minw)
	r.Zero(minh)

	data.X = 10
	data.Y = 20
	data.W = 30
	data.H = 40
	data.MaxW = 50
	data.MaxH = 60
	data.MinW = 70
	data.MinH = 80

	err = pm.RestorePosition(ctx, data)
	r.NoError(err)

	r.Equal(data.X, ax)
	r.Equal(data.Y, ay)
	r.Equal(data.W, aw)
	r.Equal(data.H, ah)
	r.Equal(data.MaxW, maxw)
	r.Equal(data.MaxH, maxh)
	r.Equal(data.MinW, minw)
	r.Equal(data.MinH, minh)
}
