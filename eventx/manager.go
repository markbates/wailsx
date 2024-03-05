package eventx

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

func NewManager() *Manager {
	return &Manager{}
}

var _ EventManager = &Manager{}
var _ statedata.StateDataProvider[*EventsData] = &Manager{}

type Manager struct {
	DisableWildcardEmits bool

	EventsEmitFn       func(ctx context.Context, name string, data ...any) error
	EventsOffAllFn     func(ctx context.Context) error
	EventsOffFn        func(ctx context.Context, name string, additional ...string) error
	EventsOnFn         func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	EventsOnMultipleFn func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	EventsOnceFn       func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)

	NowFn func() time.Time

	data EventsData
}

func (em *Manager) StateData(ctx context.Context) (statedata.StateData[*EventsData], error) {
	if em == nil {
		return statedata.StateData[*EventsData]{}, fmt.Errorf("error manager is nil")
	}

	return em.data.StateData(ctx)
}

func (em *Manager) Now() time.Time {
	if em.NowFn != nil {
		return em.NowFn()
	}

	return time.Now()
}

func (em *Manager) MarshalJSON() ([]byte, error) {
	if em == nil {
		return nil, fmt.Errorf("error manager is nil")
	}

	return json.MarshalIndent(&em.data, "", "  ")
}
