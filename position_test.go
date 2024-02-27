package wailsx

import (
	"context"
	"encoding/json"
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
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			ctx := context.Background()
			tc.pos.Update(ctx)

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

	r := require.New(t)

	catcher := &wailstest.PositionCatcher{}
	pos := NewPosition()
	pos.Positioner = Positioner{
		SetPositionFn: catcher.WindowSetPosition,
		SetSizeFn:     catcher.WindowSetSize,
	}

	pos.Layout(context.Background())

	r.Equal(pos.PosX(), catcher.X)
	r.Equal(pos.PosY(), catcher.Y)
	r.Equal(pos.Width(), catcher.W)
	r.Equal(pos.Height(), catcher.H)

}
