package windowx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Toggle_Hide(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "with func",
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
			name: "without func",
			err:  wailsrun.ErrNotAvailable("Hide"),
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			tg := Toggle{}

			ctx := context.Background()

			tg.HideFn = tc.fn

			err := tg.Hide(ctx)
			if tc.err == nil {
				r.NoError(err)
				return
			}

			r.Error(err)
			r.True(errors.Is(err, tc.err))
		})
	}
}

func Test_Toggle_Show(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "with func",
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
			name: "without func",
			err:  wailsrun.ErrNotAvailable("Show"),
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			tg := Toggle{}

			ctx := context.Background()

			tg.ShowFn = tc.fn

			err := tg.Show(ctx)
			if tc.err == nil {
				r.NoError(err)
				return
			}

			r.Error(err)
			r.True(errors.Is(err, tc.err))
		})
	}
}

func Test_Toggle_WindowHide(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "with func",
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
			name: "without func",
			err:  wailsrun.ErrNotAvailable("WindowHide"),
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			tg := Toggle{}

			ctx := context.Background()

			tg.WindowHideFn = tc.fn

			err := tg.WindowHide(ctx)
			if tc.err == nil {
				r.NoError(err)
				return
			}

			r.Error(err)
			r.True(errors.Is(err, tc.err))
		})
	}
}

func Test_Toggle_WindowShow(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "with func",
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
			name: "without func",
			err:  wailsrun.ErrNotAvailable("WindowShow"),
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			tg := Toggle{}

			ctx := context.Background()

			tg.WindowShowFn = tc.fn

			err := tg.WindowShow(ctx)
			if tc.err == nil {
				r.NoError(err)
				return
			}

			r.Error(err)
			r.True(errors.Is(err, tc.err))
		})
	}
}
