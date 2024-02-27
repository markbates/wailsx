package wailsx

import (
	"context"
	"fmt"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Positioner struct {
	GetPositionFn func(ctx context.Context) (x int, y int, err error)
	GetSizeFn     func(ctx context.Context) (w int, h int, err error)
	SetPositionFn func(ctx context.Context, x int, y int) error
	SetSizeFn     func(ctx context.Context, w int, h int) error
}

func (pos Positioner) WindowGetPosition(ctx context.Context) (x int, y int, err error) {
	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if pos.GetPositionFn == nil {
		x, y = wailsrun.WindowGetPosition(ctx)
		return x, y, nil
	}

	return pos.GetPositionFn(ctx)
}

func (pos Positioner) WindowGetSize(ctx context.Context) (w int, h int, err error) {
	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if pos.GetSizeFn == nil {
		w, h = wailsrun.WindowGetSize(ctx)
		return w, h, nil
	}

	return pos.GetSizeFn(ctx)
}

func (pos Positioner) WindowSetPosition(ctx context.Context, x int, y int) (err error) {
	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if pos.SetPositionFn == nil {
		wailsrun.WindowSetPosition(ctx, x, y)
	}

	return pos.SetPositionFn(ctx, x, y)
}

func (pos Positioner) WindowSetSize(ctx context.Context, w int, h int) (err error) {
	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if pos.SetSizeFn == nil {
		wailsrun.WindowSetSize(ctx, w, h)
		return nil
	}

	return pos.SetSizeFn(ctx, w, h)
}
