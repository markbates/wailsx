package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

type PositionManger struct {
	WindowCenterFn      func(ctx context.Context) error
	WindowGetPositionFn func(ctx context.Context) (int, int, error)
	WindowGetSizeFn     func(ctx context.Context) (int, int, error)
	WindowSetMaxSizeFn  func(ctx context.Context, width int, height int) error
	WindowSetMinSizeFn  func(ctx context.Context, width int, height int) error
	WindowSetPositionFn func(ctx context.Context, x int, y int) error
	WindowSetSizeFn     func(ctx context.Context, width int, height int) error
}

func (pm PositionManger) WindowCenter(ctx context.Context) error {
	return safe.Run(func() error {
		if pm.WindowCenterFn == nil {
			return wailsrun.WindowCenter(ctx)
		}
		return pm.WindowCenterFn(ctx)
	})
}

func (pm PositionManger) WindowGetPosition(ctx context.Context) (int, int, error) {
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

func (pm PositionManger) WindowGetSize(ctx context.Context) (int, int, error) {
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

func (pm PositionManger) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	return safe.Run(func() error {
		if pm.WindowSetMaxSizeFn == nil {
			return wailsrun.WindowSetMaxSize(ctx, width, height)
		}
		return pm.WindowSetMaxSizeFn(ctx, width, height)
	})
}

func (pm PositionManger) WindowSetMinSize(ctx context.Context, width int, height int) error {
	return safe.Run(func() error {
		if pm.WindowSetMinSizeFn == nil {
			return wailsrun.WindowSetMinSize(ctx, width, height)
		}
		return pm.WindowSetMinSizeFn(ctx, width, height)
	})
}

func (pm PositionManger) WindowSetPosition(ctx context.Context, x int, y int) error {
	return safe.Run(func() error {
		if pm.WindowSetPositionFn == nil {
			return wailsrun.WindowSetPosition(ctx, x, y)
		}
		return pm.WindowSetPositionFn(ctx, x, y)
	})
}

func (pm PositionManger) WindowSetSize(ctx context.Context, width int, height int) error {
	return safe.Run(func() error {
		if pm.WindowSetSizeFn == nil {
			return wailsrun.WindowSetSize(ctx, width, height)
		}
		return pm.WindowSetSizeFn(ctx, width, height)
	})
}
