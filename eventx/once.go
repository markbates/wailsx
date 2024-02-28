package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em EventManager) Once(ctx context.Context, name string, callback CallbackFn) (func(), error) {
	if em.OnceFn != nil {
		return em.OnceFn(ctx, name, callback)
	}

	cb := func(data ...any) {
		err := callback(data...)
		if err != nil {
			em.LogError(ctx, err.Error())
		}
	}

	fn := wailsrun.EventsOnce(ctx, name, cb)
	return fn, nil
}
