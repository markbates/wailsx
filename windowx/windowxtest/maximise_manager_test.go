package windowxtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MaximiserManager_WindowFullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowFullscreen(ctx)
	r.NoError(err)

	r.True(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	err = mm.WindowFullscreen(ctx)
	r.Error(err)
}

func Test_MaximiserManager_WindowIsFullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := MaximiserManager{}
	r.False(mm.IsFullscreen)

	ctx := context.Background()

	b, err := mm.WindowIsFullscreen(ctx)
	r.NoError(err)
	r.False(b)

	mm.IsFullscreen = true
	b, err = mm.WindowIsFullscreen(ctx)
	r.NoError(err)
	r.True(b)
}

func Test_MaximiserManager_WindowIsMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := MaximiserManager{}
	r.False(mm.IsMaximised)

	ctx := context.Background()

	b, err := mm.WindowIsMaximised(ctx)
	r.NoError(err)
	r.False(b)

	mm.IsMaximised = true
	b, err = mm.WindowIsMaximised(ctx)
	r.NoError(err)
	r.True(b)
}

func Test_MaximiserManager_WindowIsMinimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := MaximiserManager{}
	r.False(mm.IsMinimised)

	ctx := context.Background()

	b, err := mm.WindowIsMinimised(ctx)
	r.NoError(err)
	r.False(b)

	mm.IsMinimised = true
	b, err = mm.WindowIsMinimised(ctx)
	r.NoError(err)
	r.True(b)
}

func Test_MaximiserManager_WindowMaximise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowMaximise(ctx)
	r.NoError(err)

	r.False(mm.IsFullscreen)
	r.True(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	err = mm.WindowMaximise(ctx)
	r.Error(err)
}

func Test_MaximiserManager_WindowMinimise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowMinimise(ctx)
	r.NoError(err)

	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.True(mm.IsMinimised)
	r.False(mm.IsNormal)

	err = mm.WindowMinimise(ctx)
	r.Error(err)
}

func Test_MaximiserManager_WindowUnfullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowUnfullscreen(ctx)
	r.Error(err)

	mm.IsFullscreen = true
	err = mm.WindowUnfullscreen(ctx)
	r.NoError(err)

	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.True(mm.IsNormal)
}

func Test_MaximiserManager_WindowUnmaximise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowUnmaximise(ctx)
	r.Error(err)

	mm.IsMaximised = true
	err = mm.WindowUnmaximise(ctx)
	r.NoError(err)

	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.True(mm.IsNormal)
}

func Test_MaximiserManager_WindowUnminimise(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := &MaximiserManager{}
	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.False(mm.IsNormal)

	ctx := context.Background()

	err := mm.WindowUnminimise(ctx)
	r.Error(err)

	mm.IsMinimised = true
	err = mm.WindowUnminimise(ctx)
	r.NoError(err)

	r.False(mm.IsFullscreen)
	r.False(mm.IsMaximised)
	r.False(mm.IsMinimised)
	r.True(mm.IsNormal)
}

func Test_MaximiserManager_WindowIsNormal(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mm := MaximiserManager{}

	ctx := context.Background()

	b, err := mm.WindowIsNormal(ctx)
	r.NoError(err)
	r.False(b)

	mm.IsNormal = false
	b, err = mm.WindowIsNormal(ctx)
	r.NoError(err)
	r.False(b)
}
