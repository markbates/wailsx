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

func Test_Position(t *testing.T) {
	t.Parallel()

	dp := &Position{X: PosX, Y: PosY, W: PosW, H: PosH}

	tcs := []struct {
		name string
		pos  *Position
		exp  *Position
	}{
		{
			name: "default",
			pos:  &Position{},
			exp:  dp,
		},
		{
			name: "nil",
			pos:  nil,
			exp:  dp,
		},
		{
			name: "with values",
			pos: &Position{
				X: 1,
				Y: 2,
				W: 3,
				H: 4,
			},
			exp: &Position{X: 1, Y: 2, W: 3, H: 4},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			r.Equal(tc.exp.X, tc.pos.PosX())
			r.Equal(tc.exp.Y, tc.pos.PosY())
			r.Equal(tc.exp.W, tc.pos.Width())
			r.Equal(tc.exp.H, tc.pos.Height())
		})
	}

}

func Test_Position_Update(t *testing.T) {
	t.Parallel()

	dp := NewPosition()

	tcs := []struct {
		name string
		pos  *Position
		exp  *Position
		err  bool
	}{
		{
			name: "default",
			pos: &Position{
				Positioner: Positioner{
					GetPositionFn: wailstest.PositionGet(0, 0),
					GetSizeFn:     wailstest.PositionGet(0, 0),
				},
			},
			exp: dp,
		},
		{
			name: "with values",
			pos: &Position{
				Positioner: Positioner{
					GetPositionFn: wailstest.PositionGet(1, 2),
					GetSizeFn:     wailstest.PositionGet(3, 4),
				},
			},
			exp: &Position{X: 1, Y: 2, W: 3, H: 4},
		},
		{
			name: "error",
			err:  true,
			pos: &Position{
				Positioner: Positioner{
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
			pos: &Position{
				Positioner: Positioner{
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

			err := tc.pos.Update(ctx)
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)

			r.Equal(tc.exp.PosX(), tc.pos.PosX())
			r.Equal(tc.exp.PosY(), tc.pos.PosY())
			r.Equal(tc.exp.Width(), tc.pos.Width())
			r.Equal(tc.exp.Height(), tc.pos.Height())
		})
	}

}

func Test_Position_MarshalJSON(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  *Position
		err  bool
	}{
		{
			name: "empty",
			pos:  &Position{},
		},
		{
			name: "default",
			pos:  NewPosition(),
		},
		{
			name: "with values",
			pos: &Position{
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

			assertJSON(t, filepath.Join("positions", tc.name), tc.pos)
		})
	}

}

func Test_Position_Layout(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		pos  *Position
		err  bool
	}{
		{
			name: "default",
			pos:  NewPosition(),
		},
		{
			name: "set position error",
			pos: &Position{
				Positioner: Positioner{
					SetPositionFn: func(ctx context.Context, x int, y int) error {
						return io.EOF
					},
				},
			},
			err: true,
		},
		{
			name: "set position panic",
			pos: &Position{
				Positioner: Positioner{
					SetPositionFn: func(ctx context.Context, x int, y int) error {
						panic(io.EOF)
					},
				},
			},
			err: true,
		},
		{
			name: "set size error",
			pos: &Position{
				Positioner: Positioner{
					SetSizeFn: func(ctx context.Context, w int, h int) error {
						return io.EOF
					},
				},
			},
			err: true,
		},
		{
			name: "set size panic",
			pos: &Position{
				Positioner: Positioner{
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

			pos := tc.pos

			ec := &wailstest.PositionCatcher{}

			if pos.SetPositionFn == nil {
				pos.SetPositionFn = ec.WindowSetPosition
			}

			if pos.SetSizeFn == nil {
				pos.SetSizeFn = ec.WindowSetSize
			}

			err := pos.Layout(context.Background())
			if tc.err {
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
				return
			}

			r.NoError(err)

			r.Equal(pos.PosX(), ec.X)
			r.Equal(pos.PosY(), ec.Y)
			r.Equal(pos.Width(), ec.W)
			r.Equal(pos.Height(), ec.H)
		})
	}

	r := require.New(t)

	catcher := &wailstest.PositionCatcher{}
	pos := NewPosition()
	pos.Positioner = Positioner{
		SetPositionFn: catcher.WindowSetPosition,
		SetSizeFn:     catcher.WindowSetSize,
	}

	ctx := context.Background()

	err := pos.Layout(ctx)
	r.NoError(err)

	r.Equal(pos.PosX(), catcher.X)
	r.Equal(pos.PosY(), catcher.Y)
	r.Equal(pos.Width(), catcher.W)
	r.Equal(pos.Height(), catcher.H)
}
