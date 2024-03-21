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
	r.Equal(WINDOW_FULLSCREEN, md.Layout)
}

func Test_MaximiserData_SetMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetMaximised())

	md = &MaximiserData{}
	r.NoError(md.SetMaximised())
	r.Equal(WINDOW_MAXIMISED, md.Layout)
}

func Test_MaximiserData_SetMinimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetMinimised())

	md = &MaximiserData{}
	r.NoError(md.SetMinimised())
	r.Equal(WINDOW_MINIMISED, md.Layout)
}

func Test_MaximiserData_SetNormal(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetNormal())

	md = &MaximiserData{}
	r.NoError(md.SetNormal())
	r.Equal(WINDOW_NORMAL, md.Layout)
}

func Test_MaximiserData_SetUnfullscreen(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnfullscreen())

	md = &MaximiserData{}
	r.NoError(md.SetUnfullscreen())
	r.Equal(WINDOW_NORMAL, md.Layout)
}

func Test_MaximiserData_SetUnmaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnmaximised())

	md = &MaximiserData{}
	r.NoError(md.SetUnmaximised())
	r.Equal(WINDOW_NORMAL, md.Layout)
}

func Test_MaximiserData_SetUnminimised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.SetUnminimised())

	md = &MaximiserData{}
	r.NoError(md.SetUnminimised())
	r.Equal(WINDOW_NORMAL, md.Layout)
}

func Test_MaximiserData_ToggleMaximised(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var md *MaximiserData
	r.Error(md.ToggleMaximised())

	md = &MaximiserData{}
	r.NoError(md.ToggleMaximised())

	r.Equal(WINDOW_MAXIMISED, md.Layout)

	r.NoError(md.ToggleMaximised())
	r.Equal(WINDOW_NORMAL, md.Layout)
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

	r.Equal(md, sd)
}

func Test_Maximiser_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	wm := &MaximiserData{}
	r.Equal("*windowx.MaximiserData", wm.PluginName())
}
