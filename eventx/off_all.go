package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOffAll(ctx context.Context) error {
	return safe.Run(func() error {
		if em.OffAllFn == nil {
			em.OffAllFn = wailsrun.EventsOffAll
		}

		return em.OffAllFn(ctx)
	})
}
