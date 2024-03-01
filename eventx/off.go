package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOff(ctx context.Context, name string, additional ...string) error {
	return safe.Run(func() error {
		if em.EventsOffFn == nil {
			em.EventsOffFn = wailsrun.EventsOff
		}

		return em.EventsOffFn(ctx, name, additional...)
	})
}
