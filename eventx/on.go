package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOn(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {

	var fn wailsrun.CancelFn
	err := safe.Run(func() error {
		if em.EventsOnFn == nil {
			em.EventsOnFn = wailsrun.EventsOn
		}

		var err error
		fn, err = em.EventsOnFn(ctx, name, callback)
		return err
	})

	if err != nil {
		return nil, err
	}

	return fn, err
}
