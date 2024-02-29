package eventx_test

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/wailstest"
)

type EmitCatcher struct {
	Events []CaughtEvent
	Error  bool
}

type CaughtEvent struct {
	Event string
	Args  []any
}

func (ec *EmitCatcher) Emit(ctx context.Context, event string, args ...any) error {
	if ec == nil {
		return fmt.Errorf("catcher is nil")
	}

	if ec.Error {
		return wailstest.ErrTest
	}

	ec.Events = append(ec.Events, CaughtEvent{
		Event: event,
		Args:  args,
	})

	return nil
}
