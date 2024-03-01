package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOff(ctx context.Context, name string, additional ...string) error {
	if em.OffFn != nil {
		return em.OffFn(ctx, name, additional...)
	}

	return wailsrun.EventsOff(ctx, name, additional...)
}
