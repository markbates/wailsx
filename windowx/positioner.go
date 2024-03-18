package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ PositionManager = &Positioner{}
var _ InitialPositioner = &Positioner{}

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

	data PositionData
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
