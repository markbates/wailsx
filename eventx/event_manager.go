package eventx

import (
	"context"
)

type EventManager interface {
	EventsEmit(ctx context.Context, event string, args ...any) (err error)
	EventsOff(ctx context.Context, name string, additional ...string) error
	EventsOffAll(ctx context.Context) error
	EventsOn(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	EventsOnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnce(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
}

type EventManagerDataProvider interface {
	EventManager
	StateDataProvider
}
