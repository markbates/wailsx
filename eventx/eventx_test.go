package eventx

import (
	"github.com/markbates/wailsx/wailstest"
)

func newTestManager() Manager {
	return Manager{
		DisableWildcardEmits: true,
		NowFn:                wailstest.NowTime,
	}
}
