package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em EventManager) On(ctx context.Context, name string, callback CallbackFn) (func(), error) {
	if em.OnFn != nil {
		return em.OnFn(ctx, name, callback)
	}

	fn := wailsrun.EventsOn(ctx, name, func(data ...any) {
		err := callback(data...)
		if err != nil {
			em.LogError(ctx, err.Error())
		}
	})

	return fn, nil
}
