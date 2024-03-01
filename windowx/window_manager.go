package windowx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

type WindowManager interface {
	Themer
	Toggler
	Maximiser
	Positioner
	ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error)
	WindowExecJS(ctx context.Context, js string) error
	WindowPrint(ctx context.Context) error
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetTitle(ctx context.Context, title string) error
}
