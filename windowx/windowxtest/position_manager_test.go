package windowxtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PositionManager_WindowCenter(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.False(pm.IsCentered)

	err := pm.WindowCenter(context.Background())
	r.NoError(err)

	r.True(pm.IsCentered)
}

func Test_PositionManager_WindowGetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	pm.X = 10
	pm.Y = 20

	x, y, err := pm.WindowGetPosition(context.Background())
	r.NoError(err)
	r.Equal(10, x)
	r.Equal(20, y)
}

func Test_PositionManager_WindowGetSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	pm.W = 10
	pm.H = 20

	w, h, err := pm.WindowGetSize(context.Background())
	r.NoError(err)
	r.Equal(10, w)
	r.Equal(20, h)
}

func Test_PositionManager_WindowSetMaxSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.MaxW)
	r.Equal(0, pm.MaxH)

	ctx := context.Background()

	err := pm.WindowSetMaxSize(ctx, 10, 20)
	r.NoError(err)

	r.Equal(10, pm.MaxW)
	r.Equal(20, pm.MaxH)

	err = pm.WindowSetMaxSize(ctx, -1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetMinSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.MinW)
	r.Equal(0, pm.MinH)

	ctx := context.Background()

	err := pm.WindowSetMinSize(ctx, 10, 20)
	r.NoError(err)

	r.Equal(10, pm.MinW)
	r.Equal(20, pm.MinH)

	err = pm.WindowSetMinSize(ctx, -1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.X)
	r.Equal(0, pm.Y)

	ctx := context.Background()

	err := pm.WindowSetPosition(ctx, 10, 20)
	r.NoError(err)

	r.Equal(10, pm.X)
	r.Equal(20, pm.Y)

	err = pm.WindowSetPosition(ctx, -1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.W)
	r.Equal(0, pm.H)

	ctx := context.Background()

	err := pm.WindowSetSize(ctx, 10, 20)
	r.NoError(err)

	r.Equal(10, pm.W)
	r.Equal(20, pm.H)

	err = pm.WindowSetSize(ctx, -1, -1)
	r.Error(err)
}
