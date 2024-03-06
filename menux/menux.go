package menux

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

const (
	ErrNilMenu = es("menu is nil")
)

type es string

func (e es) Error() string {
	return string(e)
}

func NopManager() Manager {
	return Manager{
		MenuSetApplicationMenuFn:    func(ctx context.Context, me *menu.Menu) error { return nil },
		MenuUpdateApplicationMenuFn: func(ctx context.Context) error { return nil },
	}
}
