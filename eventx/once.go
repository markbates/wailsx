package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOnce(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if em.OnceFn != nil {
		return em.OnceFn(ctx, name, callback)
	}

	return wailsrun.EventsOnce(ctx, name, callback)
}
