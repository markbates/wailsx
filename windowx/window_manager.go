package windowx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

type WindowManager interface {
	Hide(ctx context.Context) error
	ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error)
	Show(ctx context.Context) error
	WindowCenter(ctx context.Context) error
	WindowExecJS(ctx context.Context, js string) error
	WindowFullscreen(ctx context.Context) error
	WindowGetPosition(ctx context.Context) (int, int, error)
	WindowGetSize(ctx context.Context) (int, int, error)
	WindowHide(ctx context.Context) error
	WindowIsFullscreen(ctx context.Context) (bool, error)
	WindowIsMaximised(ctx context.Context) bool
	WindowIsMinimised(ctx context.Context) bool
	WindowIsNormal(ctx context.Context) bool
	WindowMaximise(ctx context.Context) error
	WindowMinimise(ctx context.Context) error
	WindowPrint(ctx context.Context) error
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error
	WindowSetDarkTheme(ctx context.Context) error
	WindowSetLightTheme(ctx context.Context) error
	WindowSetMaxSize(ctx context.Context, width int, height int) error
	WindowSetMinSize(ctx context.Context, width int, height int) error
	WindowSetPosition(ctx context.Context, x int, y int) error
	WindowSetSize(ctx context.Context, width int, height int) error
	WindowSetSystemDefaultTheme(ctx context.Context) error
	WindowSetTitle(ctx context.Context, title string) error
	WindowUnfullscreen(ctx context.Context) error
	WindowUnmaximise(ctx context.Context) error
	WindowUnminimise(ctx context.Context) error
}
