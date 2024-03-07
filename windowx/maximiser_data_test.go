package windowx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MaximiserData_Fullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetFullscreen())

	md = &MaximiserData{}
	r.NoError(md.SetFullscreen())

	r.True(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.False(md.IsNormal)
}

func Test_MaximiserData_SetMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetMaximised())

	md = &MaximiserData{}
	r.NoError(md.SetMaximised())

	r.False(md.IsFullscreen)
	r.True(md.IsMaximised)
	r.False(md.IsMinimised)
	r.False(md.IsNormal)
}

func Test_MaximiserData_SetMinimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetMinimised())

	md = &MaximiserData{}
	r.NoError(md.SetMinimised())

	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.True(md.IsMinimised)
	r.False(md.IsNormal)
}

func Test_MaximiserData_SetNormal(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetNormal())

	md = &MaximiserData{}
	r.NoError(md.SetNormal())

	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.True(md.IsNormal)
}

func Test_MaximiserData_SetUnfullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnfullscreen())

	md = &MaximiserData{}
	r.NoError(md.SetUnfullscreen())

	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.True(md.IsNormal)
}

func Test_MaximiserData_SetUnmaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnmaximised())

	md = &MaximiserData{}
	r.NoError(md.SetUnmaximised())

	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.True(md.IsNormal)
}

func Test_MaximiserData_SetUnminimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnminimised())

	md = &MaximiserData{}
	r.NoError(md.SetUnminimised())

	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.True(md.IsNormal)
}

func Test_MaximiserData_ToggleMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.ToggleMaximised())

	md = &MaximiserData{}
	r.NoError(md.ToggleMaximised())

	r.False(md.IsFullscreen)
	r.True(md.IsMaximised)
	r.False(md.IsMinimised)
	r.False(md.IsNormal)

	r.NoError(md.ToggleMaximised())
	r.False(md.IsFullscreen)
	r.False(md.IsMaximised)
	r.False(md.IsMinimised)
	r.True(md.IsNormal)
}

func Test_MaximiserData_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var md *MaximiserData

	_, err := md.StateData(ctx)
	r.Error(err)

	md = &MaximiserData{}

	sd, err := md.StateData(ctx)
	r.NoError(err)

	r.Equal(MaximiserStateDataName, sd.Name)
	r.Equal(md, sd.Data)
}

func Test_Maximiser_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := &MaximiserData{}
	r.Equal("*windowx.MaximiserData", wm.PluginName())
}
