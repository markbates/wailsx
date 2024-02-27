package wailsx

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Positioner_WindowGetPosition(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  Positioner
		err  bool
	}{
		{
			name: "good",
			pos: Positioner{
				GetPositionFn: func(ctx context.Context) (int, int, error) {
					return 1, 2, nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			pos: Positioner{
				GetPositionFn: func(ctx context.Context) (int, int, error) {
					return 0, 0, io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			pos: Positioner{
				GetPositionFn: func(ctx context.Context) (int, int, error) {
					panic(io.EOF)
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()

			x, y, err := tc.pos.WindowGetPosition(ctx)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
			r.Equal(1, x)
			r.Equal(2, y)
		})
	}
}

func Test_Positioner_WindowGetSize(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  Positioner
		err  bool
	}{
		{
			name: "good",
			pos: Positioner{
				GetSizeFn: func(ctx context.Context) (int, int, error) {
					return 3, 4, nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			pos: Positioner{
				GetSizeFn: func(ctx context.Context) (int, int, error) {
					return 0, 0, io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			pos: Positioner{
				GetSizeFn: func(ctx context.Context) (int, int, error) {
					panic(io.EOF)
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()

			w, h, err := tc.pos.WindowGetSize(ctx)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
			r.Equal(3, w)
			r.Equal(4, h)
		})
	}
}

func Test_Positioner_WindowSetPosition(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  Positioner
		err  bool
	}{
		{
			name: "good",
			pos: Positioner{
				SetPositionFn: func(ctx context.Context, x, y int) error {
					return nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			pos: Positioner{
				SetPositionFn: func(ctx context.Context, x, y int) error {
					return io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			pos: Positioner{
				SetPositionFn: func(ctx context.Context, x, y int) error {
					panic(io.EOF)
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()

			err := tc.pos.WindowSetPosition(ctx, 5, 6)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
		})
	}
}

func Test_Positioner_WindowSetSize(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  Positioner
		err  bool
	}{
		{
			name: "good",
			pos: Positioner{
				SetSizeFn: func(ctx context.Context, w, h int) error {
					return nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			pos: Positioner{
				SetSizeFn: func(ctx context.Context, w, h int) error {
					return io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			pos: Positioner{
				SetSizeFn: func(ctx context.Context, w, h int) error {
					panic(io.EOF)
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()

			err := tc.pos.WindowSetSize(ctx, 7, 8)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
		})
	}
}
