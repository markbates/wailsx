package eventx

import (
	"github.com/markbates/wailsx/wailstest"
)

func newEventManager() (EventManager, *wailstest.EmitCatcher) {
	ec := &wailstest.EmitCatcher{}
	return EventManager{
		EmitFn:               ec.Emit,
		DisableWildcardEmits: true,
		NowFn:                wailstest.NowTime,
	}, ec
}
