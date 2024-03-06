package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

type EventManager interface {
	EventsEmit(ctx context.Context, event string, args ...any) (err error)
	EventsOff(ctx context.Context, name string, additional ...string) error
	EventsOffAll(ctx context.Context) error
	EventsOn(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	EventsOnMultiple(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	EventsOnce(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
}

type EventManagerDataProvider interface {
	EventManager
	StateDataProvider
}
