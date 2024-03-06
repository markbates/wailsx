package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ WindowManager = &Manager{}

type Manager struct {
	MaximiseManager
	PositionerManager
	ReloadManager
	ThemeManager
	Toggler

	ScreenGetAllFn         func(ctx context.Context) ([]Screen, error)
	WindowExecJSFn         func(ctx context.Context, js string) error
	WindowPrintFn          func(ctx context.Context) error
	WindowSetAlwaysOnTopFn func(ctx context.Context, b bool) error
	WindowSetTitleFn       func(ctx context.Context, title string) error
}

func NewManager() *Manager {
	return &Manager{
		MaximiseManager:   &Maximiser{},
		PositionerManager: &Positioner{},
		ReloadManager:     &Reloader{},
		ThemeManager:      &Themer{},
		Toggler:           &Toggle{},
	}
}

func NewNOOPManager() *Manager {
	return &Manager{
		MaximiseManager:   NewNOOPMaximiser(),
		PositionerManager: NewNOOPPositioner(),
		ReloadManager:     NewNOOPReloader(),
		ThemeManager:      NewNOOPThemer(),
		Toggler:           NewNOOPToggle(),

		ScreenGetAllFn: func(ctx context.Context) ([]Screen, error) {
			return nil, nil
		},
		WindowExecJSFn: func(ctx context.Context, js string) error {
			return nil
		},
		WindowPrintFn: func(ctx context.Context) error {
			return nil
		},
		WindowSetAlwaysOnTopFn: func(ctx context.Context, b bool) error {
			return nil
		},
		WindowSetTitleFn: func(ctx context.Context, title string) error {
			return nil
		},
	}
}

func (wm Manager) ScreenGetAll(ctx context.Context) ([]Screen, error) {
	var screens []Screen

	err := safe.Run(func() error {
		if wm.ScreenGetAllFn == nil {
			wm.ScreenGetAllFn = func(ctx context.Context) ([]Screen, error) {
				wsc, err := wailsrun.ScreenGetAll(ctx)
				if err != nil {
					return nil, err
				}

				screens = make([]Screen, 0, len(wsc))
				for _, sc := range wsc {
					screens = append(screens, Screen{
						Size: ScreenSize{
							Width:  sc.Size.Width,
							Height: sc.Size.Height,
						},
					})
				}
				return screens, nil
			}
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
