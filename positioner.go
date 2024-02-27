package wailsx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

// type Positioner interface {
// 	WindowGetPosition(ctx context.Context) (int, int)
// 	WindowGetSize(ctx context.Context) (int, int)
// }

type Positioner struct {
	GetPositionFn func(ctx context.Context) (int, int)
	GetSizeFn     func(ctx context.Context) (int, int)
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
