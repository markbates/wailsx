package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOnce(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
	if em == nil {
		return wailsrun.EventsOnce(ctx, name, callback)
	}

	var fn CancelFn
	err := safe.Run(func() error {
		if em.EventsOnceFn == nil {
			em.EventsOnceFn = wailsrun.EventsOnce
		}

		var err error

		fn, err = em.EventsOnceFn(ctx, name, callback)
		if err != nil {
			return err
		}

		return em.data.AddCallback(name, callback, 1)
	})

	if err != nil {
		return nil, err
	}

	return fn, err
}
