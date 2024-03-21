package windowx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ThemeData_SetDarkTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var th *ThemeData
	r.Error(th.SetDarkTheme())

	th = &ThemeData{}
	r.NoError(th.SetDarkTheme())

	r.Equal(THEME_DARK, th.Theme)
}

func Test_ThemeData_SetLightTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var th *ThemeData
	r.Error(th.SetLightTheme())

	th = &ThemeData{}
	r.NoError(th.SetLightTheme())

	r.Equal(THEME_LIGHT, th.Theme)
}

func Test_ThemeData_SetSystemTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var th *ThemeData
	r.Error(th.SetSystemTheme())

	th = &ThemeData{}
	r.NoError(th.SetSystemTheme())

	r.Equal(THEME_SYSTEM, th.Theme)
}

func Test_ThemeData_SetBackgroundColour(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var R, G, B, A uint8 = 1, 2, 3, 4

	var th *ThemeData
	r.Error(th.SetBackgroundColour(R, G, B, A))

	th = &ThemeData{}
	r.NoError(th.SetBackgroundColour(R, G, B, A))

	r.Equal(R, th.BackgroundColour.R)
	r.Equal(G, th.BackgroundColour.G)
	r.Equal(B, th.BackgroundColour.B)
	r.Equal(A, th.BackgroundColour.A)
}

func Test_ThemeData_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var th *ThemeData
	r.Equal("*windowx.ThemeData", th.PluginName())
}

func Test_ThemeData_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	th := &ThemeData{
		Theme: THEME_DARK,
		BackgroundColour: Colour{
			R: 1,
			G: 2,
			B: 3,
			A: 4,
		},
	}

	ctx := context.Background()

	sd, err := th.StateData(ctx)
	r.NoError(err)

	r.Equal(th, sd.Data)
}
