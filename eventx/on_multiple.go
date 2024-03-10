package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error) {
	if em == nil {
		return wailsrun.EventsOnMultiple(ctx, name, callback, counter)
	}

	var cancel CancelFn

	err := safe.Run(func() error {
		fn := em.EventsOnMultipleFn
		if fn == nil {
			fn = wailsrun.EventsOnMultiple
		}

		var err error
		cancel, err = fn(ctx, name, callback, counter)
		if err != nil {
			return err
		}

		return em.data.AddCallback(name, callback, counter)
	})

	if err != nil {
		return nil, err
	}

	return cancel, err
}
