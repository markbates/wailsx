package windowx

import (
	"context"
)

type WindowManager interface {
	MaximiseManager
	PositionManager
	ReloadManager
	ThemeManager
	Toggler

	ScreenGetAll(ctx context.Context) ([]Screen, error)
	WindowExecJS(ctx context.Context, js string) error
	WindowPrint(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetTitle(ctx context.Context, title string) error
}

type WindowManagerDataProvider interface {
	WindowManager
	StateDataProvider
}

type RestorableWindowManager interface {
	WindowManager
	RestoreWindows(ctx context.Context, data *WindowData) error
}
