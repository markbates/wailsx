package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ PositionManager = &Positioner{}

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
	WindowCenterFn      func(ctx context.Context) error
	WindowGetPositionFn func(ctx context.Context) (int, int, error)
	WindowGetSizeFn     func(ctx context.Context) (int, int, error)
	WindowSetMaxSizeFn  func(ctx context.Context, width int, height int) error
	WindowSetMinSizeFn  func(ctx context.Context, width int, height int) error
	WindowSetPositionFn func(ctx context.Context, x int, y int) error
	WindowSetSizeFn     func(ctx context.Context, width int, height int) error

	data PositionData
}

func (pm *Positioner) WindowCenter(ctx context.Context) error {
	if pm == nil {
		return nil
	}

	return safe.Run(func() error {
		if pm.WindowCenterFn == nil {
			return wailsrun.WindowCenter(ctx)
		}
		return pm.WindowCenterFn(ctx)
	})
}

func (pm *Positioner) WindowGetPosition(ctx context.Context) (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("positioner manager is nil")
	}

	var x, y int
	err := safe.Run(func() error {
		var err error
		if pm.WindowGetPositionFn == nil {
			pm.WindowGetPositionFn = wailsrun.WindowGetPosition
		}
		x, y, err = pm.WindowGetPositionFn(ctx)
		return err
	})

	return x, y, err
}

func (pm *Positioner) WindowGetSize(ctx context.Context) (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("positioner manager is nil")
	}

	var w, h int
	err := safe.Run(func() error {
		var err error
		if pm.WindowGetSizeFn == nil {
			pm.WindowGetSizeFn = wailsrun.WindowGetSize
		}
		w, h, err = pm.WindowGetSizeFn(ctx)
		return err
	})

	return w, h, err
}

func (pm *Positioner) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("positioner manager is nil")
	}

	return safe.Run(func() error {
		if pm.WindowSetMaxSizeFn == nil {
			return wailsrun.WindowSetMaxSize(ctx, width, height)
		}
		return pm.WindowSetMaxSizeFn(ctx, width, height)
	})
}

func (pm *Positioner) WindowSetMinSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("positioner manager is nil")
	}

	return safe.Run(func() error {
		if pm.WindowSetMinSizeFn == nil {
			return wailsrun.WindowSetMinSize(ctx, width, height)
		}
		return pm.WindowSetMinSizeFn(ctx, width, height)
	})
}

func (pm *Positioner) WindowSetPosition(ctx context.Context, x int, y int) error {
	if pm == nil {
		return fmt.Errorf("positioner manager is nil")
	}

	return safe.Run(func() error {
		if pm.WindowSetPositionFn == nil {
			return wailsrun.WindowSetPosition(ctx, x, y)
		}
		return pm.WindowSetPositionFn(ctx, x, y)
	})
}

func (pm *Positioner) WindowSetSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("positioner manager is nil")
	}

	return safe.Run(func() error {
		if pm.WindowSetSizeFn == nil {
			return wailsrun.WindowSetSize(ctx, width, height)
		}
		return pm.WindowSetSizeFn(ctx, width, height)
	})
}
