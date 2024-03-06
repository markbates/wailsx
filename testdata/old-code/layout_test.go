package wailsx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Layout(t *testing.T) {
	t.Parallel()

	dp := &Layout{x: PosX, y: PosY, w: PosW, h: PosH}

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
				x: 1,
				y: 2,
				w: 3,
				h: 4,
			},
			exp: &Layout{x: 1, y: 2, w: 3, h: 4},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			r.Equal(tc.exp.x, tc.ly.X())
			r.Equal(tc.exp.y, tc.ly.Y())
			r.Equal(tc.exp.w, tc.ly.W())
			r.Equal(tc.exp.h, tc.ly.H())
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
					GetPositionFn: wailstest.WindowGetPosition(0, 0),
					GetSizeFn:     wailstest.WindowGetSize(0, 0),
				},
			},
			exp: dp,
		},
		{
			name: "with values",
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: wailstest.WindowGetPosition(1, 2),
					GetSizeFn:     wailstest.WindowGetSize(3, 4),
				},
			},
			exp: &Layout{x: 1, y: 2, w: 3, h: 4},
		},
		{
			name: "error",
			err:  true,
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: wailstest.WindowGetPosition(-1, -1),
					GetSizeFn:     wailstest.WindowGetPosition(-1, -1),
				},
			},
		},
		{
			name: "panic error",
			err:  true,
			ly: &Layout{
				LayoutManager: LayoutManager{
					GetPositionFn: func(ctx context.Context) (x int, y int, err error) {
						panic(wailstest.ErrTest)
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
				r.True(errors.Is(err, wailstest.ErrTest))
				return
			}

			r.NoError(err)

			r.Equal(tc.exp.X(), tc.ly.X())
			r.Equal(tc.exp.Y(), tc.ly.Y())
			r.Equal(tc.exp.W(), tc.ly.W())
			r.Equal(tc.exp.H(), tc.ly.H())
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
				x: 1,
				y: 2,
				w: 3,
				h: 4,
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
						return wailstest.ErrTest
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
						panic(wailstest.ErrTest)
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
						return wailstest.ErrTest
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
						panic(wailstest.ErrTest)
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
				r.True(errors.Is(err, wailstest.ErrTest))
				return
			}

			r.NoError(err)

			r.Equal(ly.X(), ec.X)
			r.Equal(ly.Y(), ec.Y)
			r.Equal(ly.W(), ec.W)
			r.Equal(ly.H(), ec.H)
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

	r.Equal(ly.X(), catcher.X)
	r.Equal(ly.Y(), catcher.Y)
	r.Equal(ly.W(), catcher.W)
	r.Equal(ly.H(), catcher.H)
}

func Test_Layout_Set(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ly := NewLayout()

	ec := &wailstest.LayoutCatcher{}
	ly.SetPositionFn = ec.WindowSetPosition
	ly.SetSizeFn = ec.WindowSetSize
	ly.GetPositionFn = ec.WindowGetPosition
	ly.GetSizeFn = ec.WindowGetSize

	r.Equal(PosX, ly.X())
	r.Equal(PosY, ly.Y())
	r.Equal(PosW, ly.W())
	r.Equal(PosH, ly.H())

	ctx := context.Background()

	err := ly.Set(ctx, 1, 2, 3, 4)
	r.NoError(err)

	r.Equal(1, ly.X())
	r.Equal(2, ly.Y())
	r.Equal(3, ly.W())
	r.Equal(4, ly.H())

}

func Test_Layout_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ly := NewLayout()
	r.Equal(fmt.Sprintf("%T", ly), ly.PluginName())

	sd, err := ly.StateData()
	r.NoError(err)

	r.Equal("position", sd.Name)
	r.Equal(ly, sd.Data)

}
