package eventx_test

import (
	. "github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/eventx/eventxtest"
	"github.com/markbates/wailsx/wailstest"
)

func newEventManager() (EventManager, *eventxtest.EmitCatcher) {
	ec := &eventxtest.EmitCatcher{}
	return EventManager{
		EmitFn:               ec.Emit,
		DisableWildcardEmits: true,
		NowFn:                wailstest.NowTime,
	}, ec
}