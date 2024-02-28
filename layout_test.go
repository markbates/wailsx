package wailsx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"path/filepath"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Layout(t *testing.T) {
	t.Parallel()

	dp := &Layout{X: PosX, Y: PosY, W: PosW, H: PosH}

	tcs := []struct {
		name string
		ly   *Layout
		exp  *Layout
	}{
		{
			name: "default",
			ly:   &Layout{},
			exp:  dp,
		},
		{
			name: "nil",
			ly:   nil,
			exp:  dp,
		},
		{
			name: "with values",
			ly: &Layout{
				X: 1,
				Y: 2,
				W: 3,
				H: 4,
			},
			exp: &Layout{X: 1, Y: 2, W: 3, H: 4},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			r.Equal(tc.exp.X, tc.ly.PosX())
			r.Equal(tc.exp.Y, tc.ly.PosY())
			r.Equal(tc.exp.W, tc.ly.Width())
			r.Equal(tc.exp.H, tc.ly.Height())
		})
	}

}

func Test_Layout_Update(t *testing.T) {
	t.Parallel()

	dp := NewLayout()

	tcs := []struct {
		name string
		ly   *Layout
		exp  *Layout
		err  bool
	}{
		{
			name: "default",
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: wailstest.PositionGet(0, 0),
					GetSizeFn:     wailstest.PositionGet(0, 0),
				},
			},
			exp: dp,
		},
		{
			name: "with values",
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: wailstest.PositionGet(1, 2),
					GetSizeFn:     wailstest.PositionGet(3, 4),
				},
			},
			exp: &Layout{X: 1, Y: 2, W: 3, H: 4},
		},
		{
			name: "error",
			err:  true,
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: func(ctx context.Context) (x int, y int, err error) {
						return 0, 0, io.EOF
					},
					GetSizeFn: wailstest.PositionGet(0, 0),
				},
			},
		},
		{
			name: "panic error",
			err:  true,
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: func(ctx context.Context) (x int, y int, err error) {
						panic(io.EOF)
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()

			err := tc.ly.Update(ctx)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)

			r.Equal(tc.exp.PosX(), tc.ly.PosX())
			r.Equal(tc.exp.PosY(), tc.ly.PosY())
			r.Equal(tc.exp.Width(), tc.ly.Width())
			r.Equal(tc.exp.Height(), tc.ly.Height())
		})
	}

}

func Test_Layout_MarshalJSON(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  *Layout
		err  bool
	}{
		{
			name: "empty",
			pos:  &Layout{},
		},
		{
			name: "default",
			pos:  NewLayout(),
		},
		{
			name: "with values",
			pos: &Layout{
				X: 1,
				Y: 2,
				W: 3,
				H: 4,
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			if tc.err {
				_, err := json.Marshal(tc.pos)
				r.Error(err)
				return
			}

			assertJSON(t, filepath.Join("layouts", tc.name), tc.pos)
		})
	}

}

func Test_Layout_Layout(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ly   *Layout
		err  bool
	}{
		{
			name: "default",
			ly:   NewLayout(),
		},
		{
			name: "set position error",
			ly: &Layout{
				LayoutManager: LayoutManager{
					SetPositionFn: func(ctx context.Context, x int, y int) error {
						return io.EOF
					},
				},
			},
			err: true,
		},
		{
			name: "set position panic",
			ly: &Layout{
				LayoutManager: LayoutManager{
					SetPositionFn: func(ctx context.Context, x int, y int) error {
						panic(io.EOF)
					},
				},
			},
			err: true,
		},
		{
			name: "set size error",
			ly: &Layout{
				LayoutManager: LayoutManager{
					SetSizeFn: func(ctx context.Context, w int, h int) error {
						return io.EOF
					},
				},
			},
			err: true,
		},
		{
			name: "set size panic",
			ly: &Layout{
				LayoutManager: LayoutManager{
					SetSizeFn: func(ctx context.Context, w int, h int) error {
						panic(io.EOF)
					},
				},
			},
			err: true,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ly := tc.ly

			ec := &wailstest.LayoutCatcher{}

			if ly.SetPositionFn == nil {
				ly.SetPositionFn = ec.WindowSetPosition
			}

			if ly.SetSizeFn == nil {
				ly.SetSizeFn = ec.WindowSetSize
			}

			err := ly.Layout(context.Background())
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)

			r.Equal(ly.PosX(), ec.X)
			r.Equal(ly.PosY(), ec.Y)
			r.Equal(ly.Width(), ec.W)
			r.Equal(ly.Height(), ec.H)
		})
	}

	r := require.New(t)

	catcher := &wailstest.LayoutCatcher{}
	ly := NewLayout()
	ly.LayoutManager = LayoutManager{
		SetPositionFn: catcher.WindowSetPosition,
		SetSizeFn:     catcher.WindowSetSize,
	}

	ctx := context.Background()

	err := ly.Layout(ctx)
	r.NoError(err)

	r.Equal(ly.PosX(), catcher.X)
	r.Equal(ly.PosY(), catcher.Y)
	r.Equal(ly.Width(), catcher.W)
	r.Equal(ly.Height(), catcher.H)
}
