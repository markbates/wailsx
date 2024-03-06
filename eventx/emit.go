package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/eventx/msgx"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsEmit(ctx context.Context, event string, args ...any) (err error) {
	if err := em.init(ctx); err != nil {
		return err
	}

	em.mu.RLock()
	defer em.mu.RUnlock()

	if len(args) == 0 {
		args = []any{event}
	}

	fn := em.EventsEmitFn
	if fn == nil {
		fn = wailsrun.EventsEmit
	}

	events := []string{event}
	if !em.DisableWildcardEmits {
		events = append(events, "*")
	}

	var wg safe.Group

	msgs := msgx.NewMessages(event, em.Now(), args...)
	args = msgs.Any()

	for _, e := range events {
		e := e
		wg.Go(func() error {
			return fn(ctx, e, args...)
		})

		wg.Go(func() error {
			return em.data.EmitEvent(e, em.Now(), args...)
		})
	}

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}
