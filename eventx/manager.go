package eventx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

func NewManager() *Manager {
	return &Manager{
		NowFn: time.Now,
	}
}

func NewNOOPManager() *Manager {
	return &Manager{
		NowFn: func() time.Time {
			return time.Time{}
		},
		EventsEmitFn: func(ctx context.Context, name string, data ...any) error {
			return nil
		},
		EventsOffAllFn: func(ctx context.Context) error {
			return nil
		},
		EventsOffFn: func(ctx context.Context, name string, additional ...string) error {
			return nil
		},
		EventsOnFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
		EventsOnMultipleFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
		EventsOnceFn: func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
			return func() error { return nil }, nil
		},
	}
}

var _ EventManager = &Manager{}
var _ statedata.DataProvider[*EventsData] = &Manager{}

type Manager struct {
	DisableWildcardEmits bool
	DisableStateData     bool

	EventsEmitFn       func(ctx context.Context, name string, data ...any) error
	EventsOffAllFn     func(ctx context.Context) error
	EventsOffFn        func(ctx context.Context, name string, additional ...string) error
	EventsOnFn         func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	EventsOnMultipleFn func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	EventsOnceFn       func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)

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

	return json.MarshalIndent(&em.data, "", "  ")
}

func (em *Manager) init(ctx context.Context) error {
	if em == nil {
		return fmt.Errorf("error manager is nil")
	}

	em.data.DisableStateData = em.DisableStateData

	return nil
}
