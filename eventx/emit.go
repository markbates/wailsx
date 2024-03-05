package eventx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/msgx"
	"github.com/markbates/wailsx/wailsrun"
)

func (em *Manager) EventsEmit(ctx context.Context, event string, args ...any) (err error) {
	if em == nil {
		return fmt.Errorf("error manager is nil")
	}

	if len(args) == 0 {
		args = []any{event}
	}

	args = em.handleArgs(event, args...)

	fn := em.EventsEmitFn
	if fn == nil {
		fn = wailsrun.EventsEmit
	}

	err = safe.Run(func() error {
		events := []string{event}
		if !em.DisableWildcardEmits {
			events = append(events, "*")
		}

		for _, e := range events {
			err := fn(ctx, e, args...)
			if err != nil {
				return err
			}

			err = em.data.EmitEvent(e, args...)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (em *Manager) handleArgs(event string, args ...any) []any {
	for i, a := range args {
		switch t := a.(type) {
		case msgx.Messenger:
			//  do nothing
			// it's already a message
		case error:
			args[i] = msgx.ErrorMessage{
				Err: t,
				Message: msgx.Message{
					Event: event,
					Text:  t.Error(),
					Time:  em.Now(),
					Data:  t,
				},
			}
		case string:
			args[i] = msgx.Message{
				Event: event,
				Text:  t,
				Time:  em.Now(),
				Data:  t,
			}
		default:
			args[i] = msgx.Message{
				Event: event,
				Data:  t,
				Time:  em.Now(),
			}
		}
	}
	return args
}
