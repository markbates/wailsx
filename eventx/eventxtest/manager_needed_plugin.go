package eventxtest

import "github.com/markbates/wailsx/eventx"

type ManagerNeededPlugin struct {
	Called bool
}

func (mnp *ManagerNeededPlugin) SetEventManager(em eventx.EventManager) error {
	mnp.Called = true
	return nil
}

func (mnp *ManagerNeededPlugin) WasCalled() bool {
	if mnp == nil {
		return false
	}
	return mnp.Called
}
