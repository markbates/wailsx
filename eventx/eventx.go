package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
)

const (
	EventManagerStateDataName = "events"
)

func NewManager() *Manager {
	return &Manager{
		NowFn: time.Now,
	}
}

// NopManager returns a new Manager with all the functions set to no-ops
// This is useful for testing. The NowFn is set to wailstest.NowTime
func NopManager() *Manager {
	return &Manager{
		NowFn: wailstest.NowTime,
		EventsEmitFn: func(ctx context.Context, name string, data ...any) error {
			return nil
		},
		EventsOffAllFn: func(ctx context.Context) error {
			return nil
		},
		EventsOffFn: func(ctx context.Context, name string, additional ...string) error {
			return nil
		},
		EventsOnFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
		EventsOnMultipleFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
		EventsOnceFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
	}
}
