package menux

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

var _ MenuManager = Manager{}

type Manager struct {
	MenuSetApplicationMenuFn    func(ctx context.Context, menu *menu.Menu) error `json:"-"`
	MenuUpdateApplicationMenuFn func(ctx context.Context) error                  `json:"-"`
}

func (m Manager) MenuSetApplicationMenu(ctx context.Context, me *menu.Menu) error {
	if me == nil {
		return ErrNilMenu
	}

	return safe.Run(func() error {
		fn := m.MenuSetApplicationMenuFn
		if fn == nil {
			fn = wailsrun.MenuSetApplicationMenu
		}

		return fn(ctx, me)
	})
}

func (m Manager) MenuUpdateApplicationMenu(ctx context.Context) error {
	return safe.Run(func() error {
		fn := m.MenuUpdateApplicationMenuFn
		if fn == nil {
			fn = wailsrun.MenuUpdateApplicationMenu
		}

		return fn(ctx)
	})
}
