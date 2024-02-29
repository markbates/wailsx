package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em Manager) On(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
	if em.OnFn != nil {
		return em.OnFn(ctx, name, callback)
	}

	fn := wailsrun.EventsOn(ctx, name, func(data ...any) {
		err := callback(data...)
		if err != nil {
			em.LogError(ctx, err.Error())
		}
	})

	return func() error {
		fn()
		return nil
	}, nil
}
