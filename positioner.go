package wailsx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Positioner struct {
	GetPositionFn func(ctx context.Context) (int, int)
	GetSizeFn     func(ctx context.Context) (int, int)
	SetPositionFn func(ctx context.Context, x int, y int)
	SetSizeFn     func(ctx context.Context, w int, h int)
}

func (pos Positioner) WindowGetPosition(ctx context.Context) (int, int) {
	if pos.GetPositionFn == nil {
		wailsrun.WindowGetPosition(ctx)
	}

	return pos.GetPositionFn(ctx)
}

func (pos Positioner) WindowGetSize(ctx context.Context) (int, int) {
	if pos.GetSizeFn == nil {
		wailsrun.WindowGetSize(ctx)
	}

	return pos.GetSizeFn(ctx)
}

func (pos Positioner) WindowSetPosition(ctx context.Context, x int, y int) {
	if pos.SetPositionFn == nil {
		wailsrun.WindowSetPosition(ctx, x, y)
	}

	pos.SetPositionFn(ctx, x, y)
}

func (pos Positioner) WindowSetSize(ctx context.Context, w int, h int) {
	if pos.SetSizeFn == nil {
		wailsrun.WindowSetSize(ctx, w, h)
	}

	pos.SetSizeFn(ctx, w, h)
}
