package eventx

import (
	"context"
	"time"

	"github.com/markbates/wailsx/logx"
)

func NewManager() Manager {
	return Manager{}
}

var _ EventManager = Manager{}

type Manager struct {
	logx.ErrorLoggable

	DisableWildcardEmits bool

	EmitFn       func(ctx context.Context, name string, data ...any) error
	OffAllFn     func(ctx context.Context) error
	OffFn        func(ctx context.Context, name string, additional ...string) error
	OnFn         func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	OnMultipleFn func(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	OnceFn       func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)

	NowFn func() time.Time
}

func (em Manager) Now() time.Time {
	if em.NowFn != nil {
		return em.NowFn()
	}

	return time.Now()
}

func (em Manager) LogError(ctx context.Context, message string) {
	if em.ErrorLoggable == nil {
		return
	}

	em.ErrorLoggable.LogError(ctx, message)
}

func (em Manager) LogErrorf(ctx context.Context, format string, args ...any) {
	if em.ErrorLoggable == nil {
		return
	}

	em.ErrorLoggable.LogErrorf(ctx, format, args...)
}
