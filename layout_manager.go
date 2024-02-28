package wailsx

import (
	"context"
	"fmt"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

type LayoutManager struct {
	GetPositionFn func(ctx context.Context) (x int, y int, err error)
	GetSizeFn     func(ctx context.Context) (w int, h int, err error)
	SetPositionFn func(ctx context.Context, x int, y int) error
	SetSizeFn     func(ctx context.Context, w int, h int) error
}

func (ly LayoutManager) WindowGetPosition(ctx context.Context) (x int, y int, err error) {
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

	if ly.GetPositionFn == nil {
		x, y = wailsrun.WindowGetPosition(ctx)
		return x, y, nil
	}

	return ly.GetPositionFn(ctx)
}

func (ly LayoutManager) WindowGetSize(ctx context.Context) (w int, h int, err error) {
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

	if ly.GetSizeFn == nil {
		w, h = wailsrun.WindowGetSize(ctx)
		return w, h, nil
	}

	return ly.GetSizeFn(ctx)
}

func (ly LayoutManager) WindowSetPosition(ctx context.Context, x int, y int) (err error) {
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

	if ly.SetPositionFn == nil {
		wailsrun.WindowSetPosition(ctx, x, y)
	}

	return ly.SetPositionFn(ctx, x, y)
}

func (ly LayoutManager) WindowSetSize(ctx context.Context, w int, h int) (err error) {
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

	if ly.SetSizeFn == nil {
		wailsrun.WindowSetSize(ctx, w, h)
		return nil
	}

	return ly.SetSizeFn(ctx, w, h)
}
