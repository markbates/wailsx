package eventx

import (
	"time"

	"github.com/markbates/wailsx/wailstest"
)

func newEventManager() (EventManager, *wailstest.EmitCatcher) {
	ec := &wailstest.EmitCatcher{}
	return EventManager{
		EmitFn:               ec.Emit,
		DisableWildcardEmits: true,
		nowFn:                nowTime,
	}, ec
}

func nowTime() time.Time {
	return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
}

func oldTime() time.Time {
	return time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC)
}
