package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/logx"
)

type CallbackFn func(data ...any) error

func NewEventManager() EventManager {
	return EventManager{}
}

type EventManager struct {
	logx.ErrorLoggable

	DisableWildcardEmits bool

	EmitFn       func(ctx context.Context, name string, data ...any) error
	OffAllFn     func(ctx context.Context) error
	OffFn        func(ctx context.Context, name string, additional ...string) error
	OnFn         func(ctx context.Context, name string, callback CallbackFn) (func(), error)
	OnMultipleFn func(ctx context.Context, name string, callback CallbackFn, counter int) (func(), error)
	OnceFn       func(ctx context.Context, name string, callback CallbackFn) (func(), error)

	NowFn func() time.Time
}

func (em EventManager) Now() time.Time {
	if em.NowFn != nil {
		return em.NowFn()
	}

	return time.Now()
}

func (em EventManager) LogError(ctx context.Context, message string) {
	if em.ErrorLoggable == nil {
		return
	}

	em.ErrorLoggable.LogError(ctx, message)
}

func (em EventManager) LogErrorf(ctx context.Context, format string, args ...any) {
	if em.ErrorLoggable == nil {
		return
	}

	em.ErrorLoggable.LogErrorf(ctx, format, args...)
}
