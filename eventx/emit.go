package eventx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/msgx"
	"github.com/markbates/wailsx/wailsrun"
)

func (em Manager) EventsEmit(ctx context.Context, event string, args ...any) (err error) {
	if len(args) == 0 {
		args = []any{event}
	}

	for i, a := range args {
		switch t := a.(type) {
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
		case msgx.Messenger:
			//  do nothing
			// it's already a message
		default:
			args[i] = msgx.Message{
				Event: event,
				Data:  t,
				Time:  em.Now(),
			}
		}
	}

	fn := em.EmitFn
	if fn == nil {
		fn = em.wailsEmit
	}

	err = safe.Run(func() error {
		if !em.DisableWildcardEmits {
			if err := fn(ctx, "*", args...); err != nil {
				return err
			}
		}

		return fn(ctx, event, args...)
	})

	if err != nil {
		return err
	}

	return nil
}

func (em Manager) wailsEmit(ctx context.Context, event string, args ...any) error {
	return wailsrun.EventsEmit(ctx, event, args...)
}
