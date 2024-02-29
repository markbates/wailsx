package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em Manager) OnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error) {
	if em.OnMultipleFn != nil {
		return em.OnMultipleFn(ctx, name, callback, counter)
	}

	cb := func(data ...any) {
		err := callback(data...)
		if err != nil {
			em.LogError(ctx, err.Error())
		}
	}

	fn := wailsrun.EventsOnMultiple(ctx, name, cb, counter)
	return func() error {
		fn()
		return nil
	}, nil

}
