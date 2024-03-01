package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) OnMultiple(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
	if em.OnMultipleFn != nil {
		return em.OnMultipleFn(ctx, name, callback, counter)
	}

	return wailsrun.EventsOnMultiple(ctx, name, callback, counter)
}
