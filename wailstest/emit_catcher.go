package wailstest

import (
	"context"
	"fmt"
)

// EmitCatcher is a test helper to catch emitted events
type EmitCatcher struct {
	Events []CaughtEvent
}

type CaughtEvent struct {
	Event string
	Args  []any
}

func (ec *EmitCatcher) Emit(ctx context.Context, event string, args ...any) error {
	if ec == nil {
		return fmt.Errorf("catcher is nil")
	}

	ec.Events = append(ec.Events, CaughtEvent{
		Event: event,
		Args:  args,
	})

	return nil
}
