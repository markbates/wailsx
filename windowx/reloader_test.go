package windowx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Reload_WindowReload(t *testing.T) {
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
			err:  wailsrun.ErrNotAvailable("WindowReload"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			rld := Reloader{
				WindowReloadFn: tc.fn,
			}

			err := rld.WindowReload(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_Reload_WindowReloadApp(t *testing.T) {
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
			err:  wailsrun.ErrNotAvailable("WindowReloadApp"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			rld := Reloader{
				WindowReloadAppFn: tc.fn,
			}

			err := rld.WindowReloadApp(ctx)
			r.Equal(tc.err, err)
		})
	}
}
