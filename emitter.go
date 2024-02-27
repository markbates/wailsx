package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

func NewEmitter() Emitter {
	return Emitter{
		EmitFn: wailsrun.EventsEmit,
	}
}

type Emitter struct {
	DisableWildcard bool

	EmitFn func(ctx context.Context, event string, args ...any) `json:"-"`
}

func (em Emitter) Emit(ctx context.Context, event string, args ...any) {
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
					Time:  time.Now(),
				},
			}
		case string:
			args[i] = Message{
				Event: event,
				Text:  t,
				Time:  time.Now(),
			}
		case Messenger:
			//  do nothing
			// it's already a message
		default:
			args[i] = Message{
				Event: event,
				Data:  t,
				Time:  time.Now(),
			}
		}
	}

	fn := em.EmitFn
	if fn == nil {
		fn = wailsrun.EventsEmit
	}

	go func() {
		b, _ := json.MarshalIndent(args, "", "  ")
		fmt.Printf("emitting event: %q: %s\n", event, string(b))
	}()

	if !em.DisableWildcard {
		fn(ctx, "*", args...)
	}

	fn(ctx, event, args...)
}
