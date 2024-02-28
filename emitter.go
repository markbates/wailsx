package wailsx

import (
	"context"
	"fmt"
	"time"

	"github.com/markbates/wailsx/msgx"
	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewEmitter() Emitter {
	return Emitter{
		EmitFn: func(ctx context.Context, event string, args ...any) error {
			wailsrun.EventsEmit(ctx, event, args...)
			return nil
		},
		nowFn: time.Now,
	}
}

type Emitter struct {
	DisableWildcard bool `json:"disable_wildcard,omitempty"`

	EmitFn func(ctx context.Context, event string, args ...any) error `json:"-"`

	nowFn func() time.Time // for testing
}

func (em Emitter) Emit(ctx context.Context, event string, args ...any) (err error) {
	nowFn := em.nowFn
	if nowFn == nil {
		nowFn = time.Now
	}

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
					Time:  nowFn(),
					Data:  t,
				},
			}
		case string:
			args[i] = msgx.Message{
				Event: event,
				Text:  t,
				Time:  nowFn(),
				Data:  t,
			}
		case msgx.Messenger:
			//  do nothing
			// it's already a message
		default:
			args[i] = msgx.Message{
				Event: event,
				Data:  t,
				Time:  nowFn(),
			}
		}
	}

	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	fn := em.EmitFn
	if fn == nil {
		fn = em.wailsEmit
	}

	if !em.DisableWildcard {
		if err := fn(ctx, "*", args...); err != nil {
			return err
		}
	}

	return fn(ctx, event, args...)
}

func (em Emitter) wailsEmit(ctx context.Context, event string, args ...any) error {
	wailsrun.EventsEmit(ctx, event, args...)
	return nil
}
