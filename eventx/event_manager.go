package eventx

import (
	"context"
)

type EventManager interface {
	Emit(ctx context.Context, event string, args ...any) (err error)
	Off(ctx context.Context, name string, additional ...string) error
	OffAll(ctx context.Context) error
	On(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	OnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	Once(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
}
