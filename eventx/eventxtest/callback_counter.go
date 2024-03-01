package eventxtest

import (
	"fmt"

	"github.com/markbates/wailsx/wailsrun"
)

type CallbackCounter struct {
	Callback wailsrun.CallbackFn
	Called   int
	MaxCalls int
	Off      bool
}

func (cc *CallbackCounter) Call(data ...any) error {
	if cc == nil {
		return fmt.Errorf("callback counter is nil")
	}

	if cc.Off {
		return nil
	}

	cb := cc.Callback
	if cb == nil {
		cb = func(...any) error {
			return nil
		}
	}

	if cc.MaxCalls == 0 {
		cc.Called++
		return cb(data...)
	}

	if cc.Called >= cc.MaxCalls {
		return nil
	}

	cc.Called++
	return cc.Callback(data...)
}
