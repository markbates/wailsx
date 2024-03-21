package windowx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Themer_WindowSetDarkTheme(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

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
			err:  wailsrun.ErrNotAvailable("WindowSetDarkTheme"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			th := Themer{
				WindowSetDarkThemeFn: tc.fn,
			}

			err := th.WindowSetDarkTheme(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Themer_WindowSetLightTheme(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

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
			err:  wailsrun.ErrNotAvailable("WindowSetLightTheme"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			th := Themer{
				WindowSetLightThemeFn: tc.fn,
			}

			err := th.WindowSetLightTheme(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Themer_WindowSetSystemDefaultTheme(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

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
			err:  wailsrun.ErrNotAvailable("WindowSetSystemDefaultTheme"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			th := Themer{
				WindowSetSystemDefaultThemeFn: tc.fn,
			}

			err := th.WindowSetSystemDefaultTheme(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Themer_WindowSetBackgroundColour(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, R, G, B, A uint8) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, R, G, B, A uint8) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, R, G, B, A uint8) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, R, G, B, A uint8) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("WindowSetBackgroundColour"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			th := Themer{
				WindowSetBackgroundColourFn: tc.fn,
			}

			err := th.WindowSetBackgroundColour(ctx, 1, 2, 3, 4)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Themer_StateData(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	th := NopThemer()

	r.NoError(th.WindowSetBackgroundColour(ctx, 1, 2, 3, 4))

	sd, err := th.StateData(ctx)
	r.NoError(err)

	bg := sd.Data.BackgroundColour
	r.Equal(1, int(bg.R))
	r.Equal(2, int(bg.G))
	r.Equal(3, int(bg.B))
	r.Equal(4, int(bg.A))

	r.NoError(th.WindowSetDarkTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(THEME_DARK, sd.Data.Theme)

	r.NoError(th.WindowSetLightTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(THEME_LIGHT, sd.Data.Theme)

	r.NoError(th.WindowSetSystemDefaultTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(THEME_SYSTEM, sd.Data.Theme)
}

func Test_Nil_Themer(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	ctx := context.Background()

	var th *Themer

	err := th.WindowSetDarkTheme(ctx)
	r.Error(err)
	exp := wailsrun.ErrNotAvailable("WindowSetDarkTheme")
	r.Equal(exp, err)

	err = th.WindowSetLightTheme(ctx)
	r.Error(err)
	exp = wailsrun.ErrNotAvailable("WindowSetLightTheme")
	r.Equal(exp, err)

	err = th.WindowSetSystemDefaultTheme(ctx)
	r.Error(err)
	exp = wailsrun.ErrNotAvailable("WindowSetSystemDefaultTheme")
	r.Equal(exp, err)

	err = th.WindowSetBackgroundColour(ctx, 1, 2, 3, 4)
	r.Error(err)
	exp = wailsrun.ErrNotAvailable("WindowSetBackgroundColour")
	r.Equal(exp, err)

	_, err = th.StateData(ctx)
	r.Error(err)

}

func Test_Themer_RestoreTheme(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var th *Themer

	data := &ThemeData{}
	err := th.RestoreTheme(ctx, data)
	r.Error(err)

	th = NopThemer()

	err = th.RestoreTheme(ctx, nil)
	r.Error(err)

	var ar, ag, ab, aa uint8

	th.WindowSetBackgroundColourFn = func(ctx context.Context, R, G, B, A uint8) error {
		ar, ag, ab, aa = R, G, B, A
		return nil
	}

	var theme string

	th.WindowSetDarkThemeFn = func(ctx context.Context) error {
		theme = THEME_DARK
		return nil
	}

	th.WindowSetLightThemeFn = func(ctx context.Context) error {
		theme = THEME_LIGHT
		return nil
	}

	th.WindowSetSystemDefaultThemeFn = func(ctx context.Context) error {
		theme = THEME_SYSTEM
		return nil
	}

	data = &ThemeData{
		BackgroundColour: Colour{
			R: 1,
			G: 2,
			B: 3,
			A: 4,
		},
	}

	err = th.RestoreTheme(ctx, data)
	r.NoError(err)

	r.Equal(THEME_SYSTEM, theme)
	r.Equal(1, int(ar))
	r.Equal(2, int(ag))
	r.Equal(3, int(ab))
	r.Equal(4, int(aa))

	data.Theme = THEME_DARK
	err = th.RestoreTheme(ctx, data)
	r.NoError(err)

	r.Equal(THEME_DARK, theme)

	data.Theme = THEME_LIGHT
	err = th.RestoreTheme(ctx, data)
	r.NoError(err)

	r.Equal(THEME_LIGHT, theme)

	data.Theme = THEME_SYSTEM
	err = th.RestoreTheme(ctx, data)
	r.NoError(err)

	r.Equal(THEME_SYSTEM, theme)

	data.Theme = "unknown"
	err = th.RestoreTheme(ctx, data)
	r.Error(err)
}
