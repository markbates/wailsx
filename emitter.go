package wailsx

import (
	"context"
	"fmt"
	"time"

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
	DisableWildcard bool

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
			args[i] = ErrorMessage{
				Err: t,
				Message: Message{
					Event: event,
					Text:  t.Error(),
					Time:  nowFn(),
					Data:  t,
				},
			}
		case string:
			args[i] = Message{
				Event: event,
				Text:  t,
				Time:  nowFn(),
				Data:  t,
			}
		case Messenger:
			//  do nothing
			// it's already a message
		default:
			args[i] = Message{
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

	// go func() {
	// 	b, _ := json.MarshalIndent(args, "", "  ")
	// 	fmt.Printf("emitting event: %q: %s\n", event, string(b))
	// }()

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
