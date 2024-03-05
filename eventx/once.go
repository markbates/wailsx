package eventx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOnce(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if em == nil {
		return nil, fmt.Errorf("error manager is nil")
	}

	var fn wailsrun.CancelFn
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
