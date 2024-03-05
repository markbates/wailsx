package eventx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsOff(ctx context.Context, name string, additional ...string) error {
	if em == nil {
		return fmt.Errorf("error manager is nil")
	}

	return safe.Run(func() error {
		if em.EventsOffFn == nil {
			em.EventsOffFn = wailsrun.EventsOff
		}

		err := em.EventsOffFn(ctx, name, additional...)
		if err != nil {
			return err
		}

		names := append([]string{name}, additional...)
		return em.data.CallbacksOff(names...)
	})
}
