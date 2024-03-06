package menux

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

type MenuManager interface {
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
	MenuUpdateApplicationMenu(ctx context.Context) error
}
