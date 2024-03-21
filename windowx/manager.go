package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ RestorableWindowManager = &Manager{}
var _ WindowManagerDataProvider = &Manager{}
var _ statedata.DataProvider[*WindowData] = &Manager{}

type Manager struct {
	MaximiseManager
	PositionManager
	ReloadManager
	ThemeManager
	Toggler

	ScreenGetAllFn         func(ctx context.Context) ([]Screen, error)   `json:"-"`
	WindowExecJSFn         func(ctx context.Context, js string) error    `json:"-"`
	WindowPrintFn          func(ctx context.Context) error               `json:"-"`
	WindowSetAlwaysOnTopFn func(ctx context.Context, b bool) error       `json:"-"`
	WindowSetTitleFn       func(ctx context.Context, title string) error `json:"-"`
}

func NewManager() *Manager {
	return &Manager{
		MaximiseManager: &Maximiser{},
		PositionManager: &Positioner{},
		ReloadManager:   &Reloader{},
		ThemeManager:    &Themer{},
		Toggler:         &Toggle{},
	}
}

func NopManager() *Manager {
	return &Manager{
		MaximiseManager: NopMaximiser(),
		PositionManager: NopPositioner(),
		ReloadManager:   NopReloader(),
		ThemeManager:    NopThemer(),
		Toggler:         NopToggle(),

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

func (wm Manager) ScreenGetAll(ctx context.Context) (screens []Screen, err error) {

	err = safe.Run(func() error {
		fn := wm.ScreenGetAllFn
		if fn == nil {
			fn = wailsrun.ScreenGetAll
		}

		screens, err = fn(ctx)
		return err
	})

	return screens, err
}

func (wm Manager) WindowExecJS(ctx context.Context, js string) error {
	return safe.Run(func() error {
		fn := wm.WindowExecJSFn
		if fn == nil {
			fn = wailsrun.WindowExecJS
		}

		return fn(ctx, js)
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

func (wm *Manager) StateData(ctx context.Context) (*WindowData, error) {
	if wm == nil {
		return nil, nil
	}

	data := &WindowData{}

	if x, ok := wm.MaximiseManager.(MaximiseManagerDataProvider); ok {
		md, err := x.StateData(ctx)
		if err != nil {
			return nil, err
		}

		data.MaximiserData = md
	}

	if x, ok := wm.PositionManager.(PositionManagerDataProvider); ok {
		pd, err := x.StateData(ctx)
		if err != nil {
			return nil, err
		}
		data.PositionData = pd
	}

	if x, ok := wm.ThemeManager.(ThemeManagerDataProvider); ok {
		td, err := x.StateData(ctx)
		if err != nil {
			return nil, err
		}
		data.ThemeData = td
	}

	return data, nil
}

func (wm *Manager) RestoreWindows(ctx context.Context, data *WindowData) error {
	if wm == nil {
		return fmt.Errorf("manager is nil")
	}

	if data == nil {
		return fmt.Errorf("data is nil")
	}

	md := data.MaximiserData
	pd := data.PositionData
	td := data.ThemeData

	if md != nil {
		if err := wm.restoreMaximiser(ctx, md); err != nil {
			return err
		}
	}

	if pd != nil {
		if err := wm.restorePosition(ctx, pd); err != nil {
			return err
		}
	}

	if td != nil {
		if err := wm.restoreTheme(ctx, td); err != nil {
			return err
		}
	}

	return nil
}

func (wm *Manager) restoreMaximiser(ctx context.Context, data *MaximiserData) error {
	if x, ok := wm.MaximiseManager.(RestorableMaximiseManager); ok {
		return x.RestoreMaximiser(ctx, data)
	}
	return nil
}

func (wm *Manager) restorePosition(ctx context.Context, data *PositionData) error {
	if x, ok := wm.PositionManager.(RestorablePositionManager); ok {
		return x.RestorePosition(ctx, data)
	}
	return nil
}

func (wm *Manager) restoreTheme(ctx context.Context, data *ThemeData) error {
	if x, ok := wm.ThemeManager.(RestorableThemeManager); ok {
		return x.RestoreTheme(ctx, data)
	}
	return nil
}
