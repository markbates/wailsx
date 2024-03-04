package windowx

import (
	"context"
)

type WindowManager interface {
	Maximiser
	Positioner
	Reloader
	Themer
	Toggler
	ScreenGetAll(ctx context.Context) ([]Screen, error)
	WindowExecJS(ctx context.Context, js string) error
	WindowPrint(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetTitle(ctx context.Context, title string) error
}
