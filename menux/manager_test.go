package menux

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func Test_Manager_MenuSetApplicationMenu(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	me := &menu.Menu{}

	tcs := []struct {
		name string
		menu *menu.Menu
		fn   func(ctx context.Context, me *menu.Menu) error
		err  error
	}{
		{
			name: "with function",
			menu: me,
			fn: func(ctx context.Context, me *menu.Menu) error {
				return nil
			},
		},
		{
			name: "with nil function",
			menu: me,
			err:  wailsrun.ErrNotAvailable("MenuSetApplicationMenu"),
		},
		{
			name: "with error",
			menu: me,
			fn: func(ctx context.Context, me *menu.Menu) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			menu: me,
			fn: func(ctx context.Context, me *menu.Menu) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "nil menu",
			err:  ErrNilMenu,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			m := Manager{
				MenuSetApplicationMenuFn: tc.fn,
			}

			err := m.MenuSetApplicationMenu(ctx, tc.menu)
			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
		})
	}
}

func Test_Manager_MenuUpdateApplicationMenu(t *testing.T) {
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
			name: "with nil function",
			err:  wailsrun.ErrNotAvailable("MenuUpdateApplicationMenu"),
		},
		{
			name: "with error",
			fn: func(ctx context.Context) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			m := Manager{
				MenuUpdateApplicationMenuFn: tc.fn,
			}

			err := m.MenuUpdateApplicationMenu(ctx)
			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
		})
	}
}
