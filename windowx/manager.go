package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ WindowManagerDataProvider = &Manager{}

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

func (wm *Manager) StateData(ctx context.Context) (statedata.Data[*WindowData], error) {
	sd := statedata.Data[*WindowData]{
		Name: ManagerStateDataName,
	}

	if wm == nil {
		return sd, nil
	}

	data := &WindowData{}

	if x, ok := wm.MaximiseManager.(MaximiseManagerDataProvider); ok {
		md, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.MaximiserData = md.Data
	}

	if x, ok := wm.PositionManager.(PositionManagerDataProvider); ok {
		pd, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.PositionData = pd.Data
	}

	if x, ok := wm.ThemeManager.(ThemeManagerDataProvider); ok {
		td, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.ThemeData = td.Data
	}

	sd.Data = data

	return sd, nil
}
