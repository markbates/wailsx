package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) On(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if em.OnFn != nil {
		return em.OnFn(ctx, name, callback)
	}

	return wailsrun.EventsOn(ctx, name, callback)
}
