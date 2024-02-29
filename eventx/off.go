package eventx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (em Manager) Off(ctx context.Context, name string, additional ...string) error {
	if em.OffFn != nil {
		return em.OffFn(ctx, name, additional...)
	}

	wailsrun.EventsOff(ctx, name, additional...)
	return nil
}
