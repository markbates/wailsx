package eventx

import (
	"fmt"
	"sync"
)

type CallbackCounter struct {
	Called   int  `json:"called"`
	MaxCalls int  `json:"max_calls"`
	Off      bool `json:"off"`

	my sync.RWMutex
}

func (cc *CallbackCounter) Catch(data ...any) (bool, error) {
	if cc == nil {
		return false, fmt.Errorf("callback counter is nil")
	}

	cc.my.Lock()
	defer cc.my.Unlock()

	if cc.Off {
		return false, nil
	}

	if cc.MaxCalls == 0 {
		cc.Called++
		return true, nil
	}

	if cc.Called >= cc.MaxCalls {
		return false, nil
	}

	cc.Called++
	return true, nil
}
