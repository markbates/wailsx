package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em Manager) OffAll(ctx context.Context) error {
	if em.OffAllFn != nil {
		return em.OffAllFn(ctx)
	}

	wailsrun.EventsOffAll(ctx)
	return nil
}
