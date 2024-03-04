package windowxtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PositionManager_WindowCenter(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.False(pm.IsCentered)

	err := pm.WindowCenter()
	r.NoError(err)

	r.True(pm.IsCentered)
}

func Test_PositionManager_WindowGetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	pm.X = 10
	pm.Y = 20

	x, y, err := pm.WindowGetPosition()
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

	w, h, err := pm.WindowGetSize()
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

	err := pm.WindowSetMaxSize(10, 20)
	r.NoError(err)

	r.Equal(10, pm.MaxW)
	r.Equal(20, pm.MaxH)

	err = pm.WindowSetMaxSize(-1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetMinSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.MinW)
	r.Equal(0, pm.MinH)

	err := pm.WindowSetMinSize(10, 20)
	r.NoError(err)

	r.Equal(10, pm.MinW)
	r.Equal(20, pm.MinH)

	err = pm.WindowSetMinSize(-1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetPosition(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.X)
	r.Equal(0, pm.Y)

	err := pm.WindowSetPosition(10, 20)
	r.NoError(err)

	r.Equal(10, pm.X)
	r.Equal(20, pm.Y)

	err = pm.WindowSetPosition(-1, -1)
	r.Error(err)
}

func Test_PositionManager_WindowSetSize(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	pm := &PositionManger{}
	r.Equal(0, pm.W)
	r.Equal(0, pm.H)

	err := pm.WindowSetSize(10, 20)
	r.NoError(err)

	r.Equal(10, pm.W)
	r.Equal(20, pm.H)

	err = pm.WindowSetSize(-1, -1)
	r.Error(err)
}
