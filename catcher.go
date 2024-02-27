package wailsx

import "context"

// EmitCatcher is a test helper to catch emitted events
type EmitCatcher struct {
	Events []CaughtEvent
}

type CaughtEvent struct {
	Event string
	Args  []any
}

func (ec *EmitCatcher) Emit(ctx context.Context, event string, args ...any) {
	if ec == nil {
		return
	}

	ec.Events = append(ec.Events, CaughtEvent{
		Event: event,
		Args:  args,
	})
}
