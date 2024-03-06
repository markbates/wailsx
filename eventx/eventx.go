package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/wailsrun"
)

const (
	EventManagerStateDataName = "events"
)

func NewManager() *Manager {
	return &Manager{
		NowFn: time.Now,
	}
}

func NopManager() *Manager {
	return &Manager{
		NowFn: time.Now,
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
