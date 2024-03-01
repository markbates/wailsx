package eventx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsOffAll(ctx context.Context) error {
	if em.OffAllFn != nil {
		return em.OffAllFn(ctx)
	}

	return wailsrun.EventsOffAll(ctx)
}
