package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOffAll(ctx context.Context) error {
	if em == nil {
		return wailsrun.EventsOffAll(ctx)
	}

	return safe.Run(func() error {
		if em.EventsOffAllFn == nil {
			em.EventsOffAllFn = wailsrun.EventsOffAll
		}

		err := em.EventsOffAllFn(ctx)
		if err != nil {
			return err
		}

		return em.data.CallbacksOffAll()
	})
}
