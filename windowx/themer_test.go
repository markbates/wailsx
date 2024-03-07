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
	r.Equal(ThemeStataDataName, sd.Name)

	bg := sd.Data.BackgroundColour
	r.Equal(1, int(bg.R))
	r.Equal(2, int(bg.G))
	r.Equal(3, int(bg.B))
	r.Equal(4, int(bg.A))

	r.NoError(th.WindowSetDarkTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(ThemeStataDataName, sd.Name)
	r.True(sd.Data.IsDarkTheme)
	r.False(sd.Data.IsLightTheme)
	r.False(sd.Data.IsSystemTheme)

	r.NoError(th.WindowSetLightTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(ThemeStataDataName, sd.Name)
	r.False(sd.Data.IsDarkTheme)
	r.True(sd.Data.IsLightTheme)
	r.False(sd.Data.IsSystemTheme)

	r.NoError(th.WindowSetSystemDefaultTheme(ctx))

	sd, err = th.StateData(ctx)
	r.NoError(err)

	r.Equal(ThemeStataDataName, sd.Name)
	r.False(sd.Data.IsDarkTheme)
	r.False(sd.Data.IsLightTheme)
	r.True(sd.Data.IsSystemTheme)
}
