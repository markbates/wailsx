package windowx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

var _ WindowManager = &Manager{}

type Manager struct {
	ThemeManager
	HideFn                      func(ctx context.Context) error
	ScreenGetAllFn              func(ctx context.Context) ([]wailsrun.Screen, error)
	ShowFn                      func(ctx context.Context) error
	WindowCenterFn              func(ctx context.Context) error
	WindowExecJSFn              func(ctx context.Context, js string) error
	WindowFullscreenFn          func(ctx context.Context) error
	WindowGetPositionFn         func(ctx context.Context) (int, int, error)
	WindowGetSizeFn             func(ctx context.Context) (int, int, error)
	WindowHideFn                func(ctx context.Context) error
	WindowIsFullscreenFn        func(ctx context.Context) (bool, error)
	WindowIsMaximisedFn         func(ctx context.Context) bool
	WindowIsMinimisedFn         func(ctx context.Context) bool
	WindowIsNormalFn            func(ctx context.Context) bool
	WindowMaximiseFn            func(ctx context.Context) error
	WindowMinimiseFn            func(ctx context.Context) error
	WindowPrintFn               func(ctx context.Context) error
	WindowReloadFn              func(ctx context.Context) error
	WindowReloadAppFn           func(ctx context.Context) error
	WindowSetAlwaysOnTopFn      func(ctx context.Context, b bool) error
	WindowSetBackgroundColourFn func(ctx context.Context, R, G, B, A uint8) error
	WindowSetMaxSizeFn          func(ctx context.Context, width int, height int) error
	WindowSetMinSizeFn          func(ctx context.Context, width int, height int) error
	WindowSetPositionFn         func(ctx context.Context, x int, y int) error
	WindowSetSizeFn             func(ctx context.Context, width int, height int) error
	WindowSetTitleFn            func(ctx context.Context, title string) error
	WindowUnfullscreenFn        func(ctx context.Context) error
	WindowUnmaximiseFn          func(ctx context.Context) error
	WindowUnminimiseFn          func(ctx context.Context) error
}

func (wm *Manager) Hide(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error) {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) Show(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowCenter(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowExecJS(ctx context.Context, js string) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowFullscreen(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowGetPosition(ctx context.Context) (int, int, error) {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowGetSize(ctx context.Context) (int, int, error) {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowHide(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowIsFullscreen(ctx context.Context) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowIsMaximised(ctx context.Context) bool {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowIsMinimised(ctx context.Context) bool {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowIsNormal(ctx context.Context) bool {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowMaximise(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowMinimise(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowPrint(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowReload(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowReloadApp(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetBackgroundColour(ctx context.Context, R uint8, G uint8, B uint8, A uint8) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetMinSize(ctx context.Context, width int, height int) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetPosition(ctx context.Context, x int, y int) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetSize(ctx context.Context, width int, height int) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowSetTitle(ctx context.Context, title string) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowUnfullscreen(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowUnmaximise(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (wm *Manager) WindowUnminimise(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}
