package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ WindowManager = Manager{}

type Manager struct {
	MaximiserManager
	PositionManger
	Reload
	ThemeManager
	Toggle

	ScreenGetAllFn         func(ctx context.Context) ([]wailsrun.Screen, error)
	WindowExecJSFn         func(ctx context.Context, js string) error
	WindowPrintFn          func(ctx context.Context) error
	WindowSetAlwaysOnTopFn func(ctx context.Context, b bool) error
	WindowSetTitleFn       func(ctx context.Context, title string) error
}

func (wm Manager) ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error) {
	var screens []wailsrun.Screen

	err := safe.Run(func() error {
		if wm.ScreenGetAllFn == nil {
			wm.ScreenGetAllFn = wailsrun.ScreenGetAll
		}

		var err error
		screens, err = wm.ScreenGetAllFn(ctx)

		return err
	})

	return screens, err
}

func (wm Manager) WindowExecJS(ctx context.Context, js string) error {
	return safe.Run(func() error {
		if wm.WindowExecJSFn == nil {
			return wailsrun.WindowExecJS(ctx, js)
		}
		return wm.WindowExecJSFn(ctx, js)
	})
}

func (wm Manager) WindowPrint(ctx context.Context) error {
	return safe.Run(func() error {
		if wm.WindowPrintFn == nil {
			return wailsrun.WindowPrint(ctx)
		}
		return wm.WindowPrintFn(ctx)
	})
}

func (wm Manager) WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	return safe.Run(func() error {
		if wm.WindowSetAlwaysOnTopFn == nil {
			return wailsrun.WindowSetAlwaysOnTop(ctx, b)
		}
		return wm.WindowSetAlwaysOnTopFn(ctx, b)
	})
}

func (wm Manager) WindowSetTitle(ctx context.Context, title string) error {
	return safe.Run(func() error {
		if wm.WindowSetTitleFn == nil {
			return wailsrun.WindowSetTitle(ctx, title)
		}
		return wm.WindowSetTitleFn(ctx, title)
	})
}
