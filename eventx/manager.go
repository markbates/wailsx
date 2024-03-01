package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/logx"
	"github.com/markbates/wailsx/wailsrun"
)

func NewManager() Manager {
	return Manager{}
}

var _ EventManager = Manager{}

type Manager struct {
	logx.ErrorLoggable

	DisableWildcardEmits bool

	EventsEmitFn       func(ctx context.Context, name string, data ...any) error
	EventsOffAllFn     func(ctx context.Context) error
	EventsOffFn        func(ctx context.Context, name string, additional ...string) error
	EventsOnFn         func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	EventsOnMultipleFn func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	EventsOnceFn       func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)

	NowFn func() time.Time
}

func (em Manager) Now() time.Time {
	if em.NowFn != nil {
		return em.NowFn()
	}

	return time.Now()
}

func (em Manager) LogError(ctx context.Context, message string) error {
	if em.ErrorLoggable == nil {
		return wailsrun.LogError(ctx, message)
	}

	return em.ErrorLoggable.LogError(ctx, message)
}

func (em Manager) LogErrorf(ctx context.Context, format string, args ...any) error {
	if em.ErrorLoggable == nil {
		return wailsrun.LogErrorf(ctx, format, args...)
	}

	return em.ErrorLoggable.LogErrorf(ctx, format, args...)
}
