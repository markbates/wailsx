package wailsx

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LayoutManager_WindowGetPosition(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ly   LayoutManager
		err  bool
	}{
		{
			name: "good",
			ly: LayoutManager{
				GetPositionFn: func(ctx context.Context) (int, int, error) {
					return 1, 2, nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			ly: LayoutManager{
				GetPositionFn: func(ctx context.Context) (int, int, error) {
					return 0, 0, io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			ly: LayoutManager{
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

			x, y, err := tc.ly.WindowGetPosition(ctx)
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

func Test_LayoutManager_WindowGetSize(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ly   LayoutManager
		err  bool
	}{
		{
			name: "good",
			ly: LayoutManager{
				GetSizeFn: func(ctx context.Context) (int, int, error) {
					return 3, 4, nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			ly: LayoutManager{
				GetSizeFn: func(ctx context.Context) (int, int, error) {
					return 0, 0, io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			ly: LayoutManager{
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

			w, h, err := tc.ly.WindowGetSize(ctx)
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

func Test_LayoutManager_WindowSetPosition(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ly   LayoutManager
		err  bool
	}{
		{
			name: "good",
			ly: LayoutManager{
				SetPositionFn: func(ctx context.Context, x, y int) error {
					return nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			ly: LayoutManager{
				SetPositionFn: func(ctx context.Context, x, y int) error {
					return io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			ly: LayoutManager{
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

			err := tc.ly.WindowSetPosition(ctx, 5, 6)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
		})
	}
}

func Test_LayoutManager_WindowSetSize(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ly   LayoutManager
		err  bool
	}{
		{
			name: "good",
			ly: LayoutManager{
				SetSizeFn: func(ctx context.Context, w, h int) error {
					return nil
				},
			},
		},
		{
			name: "error",
			err:  true,
			ly: LayoutManager{
				SetSizeFn: func(ctx context.Context, w, h int) error {
					return io.EOF
				},
			},
		},
		{
			name: "panic",
			err:  true,
			ly: LayoutManager{
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

			err := tc.ly.WindowSetSize(ctx, 7, 8)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)
		})
	}
}
