package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ PositionManager = &Positioner{}
var _ InitialPositioner = &Positioner{}
var _ RestorablePositionManager = &Positioner{}

func NopPositioner() *Positioner {
	return &Positioner{
		WindowCenterFn:      func(ctx context.Context) error { return nil },
		WindowGetPositionFn: func(ctx context.Context) (int, int, error) { return 0, 0, nil },
		WindowGetSizeFn:     func(ctx context.Context) (int, int, error) { return 0, 0, nil },
		WindowSetMaxSizeFn:  func(ctx context.Context, width int, height int) error { return nil },
		WindowSetMinSizeFn:  func(ctx context.Context, width int, height int) error { return nil },
		WindowSetPositionFn: func(ctx context.Context, x int, y int) error { return nil },
		WindowSetSizeFn:     func(ctx context.Context, width int, height int) error { return nil },
	}
}

type Positioner struct {
	WindowCenterFn      func(ctx context.Context) error                        `json:"-"`
	WindowGetPositionFn func(ctx context.Context) (int, int, error)            `json:"-"`
	WindowGetSizeFn     func(ctx context.Context) (int, int, error)            `json:"-"`
	WindowSetMaxSizeFn  func(ctx context.Context, width int, height int) error `json:"-"`
	WindowSetMinSizeFn  func(ctx context.Context, width int, height int) error `json:"-"`
	WindowSetPositionFn func(ctx context.Context, x int, y int) error          `json:"-"`
	WindowSetSizeFn     func(ctx context.Context, width int, height int) error `json:"-"`

	// default values for the window
	DefX int `json:"def_x"`
	DefY int `json:"def_y"`
	DefW int `json:"def_w"`
	DefH int `json:"def_h"`

	data PositionData

	mu sync.RWMutex
}

func (pm *Positioner) InitPosX() int {
	if pm == nil {
		return 0
	}

	return pm.DefX
}

func (pm *Positioner) InitPosY() int {
	if pm == nil {
		return 0
	}

	return pm.DefY
}

func (pm *Positioner) InitWidth() int {
	if pm == nil {
		return 1200
	}

	if pm.DefW > 0 {
		return pm.DefW
	}

	return 1200
}

func (pm *Positioner) InitHeight() int {
	if pm == nil {
		return 800
	}

	if pm.DefH > 0 {
		return pm.DefH
	}

	return 800
}

func (pm *Positioner) WindowCenter(ctx context.Context) error {
	if pm == nil {
		return wailsrun.WindowCenter(ctx)
	}

	return safe.Run(func() error {
		fn := pm.WindowCenterFn
		if fn == nil {
			fn = wailsrun.WindowCenter
		}

		if err := fn(ctx); err != nil {
			return err
		}

		return pm.data.SetCentered()
	})
}

func (pm *Positioner) WindowGetPosition(ctx context.Context) (int, int, error) {
	if pm == nil {
		return wailsrun.WindowGetPosition(ctx)
	}

	var x, y int
	err := safe.Run(func() error {
		var err error

		fn := pm.WindowGetPositionFn
		if fn == nil {
			fn = wailsrun.WindowGetPosition
		}

		x, y, err = fn(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return x, y, err
}

func (pm *Positioner) WindowGetSize(ctx context.Context) (int, int, error) {
	if pm == nil {
		return wailsrun.WindowGetSize(ctx)
	}

	var w, h int
	err := safe.Run(func() error {
		var err error

		fn := pm.WindowGetSizeFn
		if fn == nil {
			fn = wailsrun.WindowGetSize
		}

		w, h, err = fn(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return w, h, err
}

func (pm *Positioner) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return wailsrun.WindowSetMaxSize(ctx, width, height)
	}

	return safe.Run(func() error {
		fn := pm.WindowSetMaxSizeFn
		if fn == nil {
			fn = wailsrun.WindowSetMaxSize
		}

		if err := fn(ctx, width, height); err != nil {
			return err
		}

		return pm.data.SetMaxSize(width, height)
	})
}

func (pm *Positioner) WindowSetMinSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return wailsrun.WindowSetMinSize(ctx, width, height)
	}

	return safe.Run(func() error {
		fn := pm.WindowSetMinSizeFn
		if fn == nil {
			fn = wailsrun.WindowSetMinSize
		}

		if err := fn(ctx, width, height); err != nil {
			return err
		}

		return pm.data.SetMinSize(width, height)
	})
}

func (pm *Positioner) WindowSetPosition(ctx context.Context, x int, y int) error {
	if pm == nil {
		return wailsrun.WindowSetPosition(ctx, x, y)
	}

	return safe.Run(func() error {
		fn := pm.WindowSetPositionFn
		if fn == nil {
			fn = wailsrun.WindowSetPosition
		}

		if err := fn(ctx, x, y); err != nil {
			return err
		}

		return pm.data.SetPosition(x, y)
	})
}

func (pm *Positioner) WindowSetSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return wailsrun.WindowSetSize(ctx, width, height)
	}

	return safe.Run(func() error {
		fn := pm.WindowSetSizeFn
		if fn == nil {
			fn = wailsrun.WindowSetSize
		}

		if err := fn(ctx, width, height); err != nil {
			return err
		}

		return pm.data.SetSize(width, height)
	})
}

func (pm *Positioner) RestorePosition(ctx context.Context, data *PositionData) error {
	if pm == nil {
		return fmt.Errorf("positioner is nil")
	}

	if data == nil {
		return fmt.Errorf("position data is nil")
	}

	x := data.X
	if x == 0 {
		x = pm.InitPosX()
	}

	y := data.Y
	if y == 0 {
		y = pm.InitPosY()
	}

	if err := pm.WindowSetPosition(ctx, x, y); err != nil {
		return err
	}

	w := data.W
	if w == 0 {
		w = pm.InitWidth()
	}

	h := data.H
	if h == 0 {
		h = pm.InitHeight()
	}

	if err := pm.WindowSetSize(ctx, w, h); err != nil {
		return err
	}

	maxW := data.MaxW
	maxH := data.MaxH
	if maxW > 0 && maxH > 0 {
		if err := pm.WindowSetMaxSize(ctx, maxW, maxH); err != nil {
			return err
		}
	}

	minW := data.MinW
	minH := data.MinH
	if minW > 0 && minH > 0 {
		if err := pm.WindowSetMinSize(ctx, minW, minH); err != nil {
			return err
		}
	}

	return nil
}

func (pm *Positioner) StateData(ctx context.Context) (*PositionData, error) {
	if pm == nil {
		return nil, fmt.Errorf("positioner manager is nil")
	}

	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.data.StateData(ctx)
}
