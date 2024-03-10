package eventx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/markbates/wailsx/statedata"
)

var _ EventManagerDataProvider = &Manager{}

type Manager struct {
	DisableWildcardEmits bool
	DisableStateData     bool

	EventsEmitFn       func(ctx context.Context, name string, data ...any) error
	EventsOffAllFn     func(ctx context.Context) error
	EventsOffFn        func(ctx context.Context, name string, additional ...string) error
	EventsOnFn         func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	EventsOnMultipleFn func(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnceFn       func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)

	NowFn func() time.Time

	mu   sync.RWMutex
	data EventsData
}

func (em *Manager) StateData(ctx context.Context) (statedata.Data[*EventsData], error) {
	if em == nil {
		return statedata.Data[*EventsData]{}, fmt.Errorf("error manager is nil")
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

	return json.Marshal(em.data)
}

func (em *Manager) init(ctx context.Context) error {
	if em == nil {
		return fmt.Errorf("error manager is nil")
	}

	em.data.DisableStateData = em.DisableStateData

	return nil
}
