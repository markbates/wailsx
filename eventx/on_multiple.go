package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOnMultiple(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {

	var fn wailsrun.CancelFn

	err := safe.Run(func() error {
		if em.EventsOnMultipleFn == nil {
			em.EventsOnMultipleFn = wailsrun.EventsOnMultiple
		}

		var err error
		fn, err = em.EventsOnMultipleFn(ctx, name, callback, counter)
		return err
	})

	if err != nil {
		return nil, err
	}

	return fn, err
}
