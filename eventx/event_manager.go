package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

type EventManager interface {
	Emit(ctx context.Context, event string, args ...any) (err error)
	Off(ctx context.Context, name string, additional ...string) error
	OffAll(ctx context.Context) error
	On(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	OnMultiple(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	Once(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
}
