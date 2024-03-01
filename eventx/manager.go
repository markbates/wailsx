package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/wailsrun"
)

func NewManager() Manager {
	return Manager{}
}

var _ EventManager = Manager{}

type Manager struct {
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
