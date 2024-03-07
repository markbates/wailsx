package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOn(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if em == nil {
		return wailsrun.EventsOn(ctx, name, callback)
	}

	var fn wailsrun.CancelFn
	err := safe.Run(func() error {
		if em.EventsOnFn == nil {
			em.EventsOnFn = wailsrun.EventsOn
		}

		var err error
		fn, err = em.EventsOnFn(ctx, name, callback)
		if err != nil {
			return err
		}

		return em.data.AddCallback(name, callback, 0)
	})

	if err != nil {
		return nil, err
	}

	return fn, err
}
