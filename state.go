package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/markbates/plugins"
	"github.com/markbates/safe"
	"github.com/markbates/wailsx/eventx"
)

var _ Saver = &State{}
var _ Shutdowner = &State{}
var _ Startuper = &State{}
var _ plugins.Plugin = &State{}

type State struct {
	*eventx.Manager // emit events
	*Layout         // layout of the app

	Name    string          // application name
	Plugins plugins.Plugins // plugins for the state

	// save function, if nil, save to file in ~/.config/<name>/state.json
	// will call Saver plugins
	SaveFn func(ctx context.Context) error

	// startup function, if nil, load from file in ~/.config/<name>/state.json
	// will call Startuper plugins
	StartupFn func(ctx context.Context) error

	// shutdown function, if nil, call Save
	// will call Shutdowner plugins
	ShutdownFn func(ctx context.Context) error

	// dom ready function, if nil, do nothing
	// will call DomReadyer plugins
	DomReadyFn func(ctx context.Context) error

	// before close function, if nil, do nothing
	// will call BeforeCloser plugins
	BeforeCloseFn func(ctx context.Context) error

	mu sync.RWMutex
}

func NewState(name string, plugins ...plugins.Plugin) (*State, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}

	st := &State{
		Name:    name,
		Manager: eventx.NewManager(),
		Plugins: plugins,
		Layout:  NewLayout(),
	}

	return st, nil
}

func (st *State) MarshalJSON() ([]byte, error) {
	mm, err := st.JSONMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(mm)
}

func (st *State) JSONMap() (map[string]any, error) {
	if st == nil {
		return nil, fmt.Errorf("state is nil")
	}

	st.mu.RLock()
	defer st.mu.RUnlock()

	pos := st.Layout
	if pos == nil {
		pos = NewLayout()
	}

	mm := map[string]any{
		"name":     st.Name,
		"position": pos,
	}

	list, err := st.stateDataPlugins()
	if err != nil {
		return nil, err
	}

	for _, sd := range list {
		mm[sd.Name] = sd.Data
	}

	return mm, nil
}

func (st *State) PluginName() string {
	return fmt.Sprintf("%T: %s", st, st.Name)
}

func (st *State) stateDataPlugins() ([]StateData, error) {
	if st == nil {
		return nil, fmt.Errorf("state is nil")
	}

	var list []StateData

	var mu sync.Mutex

	var wg safe.Group

	sdps := plugins.ByType[StateDataProvider](st.Plugins)

	for _, s := range sdps {
		s := s
		wg.Go(func() error {
			sd, err := s.StateData()
			if err != nil {
				return err
			}

			mu.Lock()
			list = append(list, sd)
			mu.Unlock()
			return nil
		})
	}

	if err := wg.Wait(); err != nil {
		return nil, err
	}

	return list, nil
}
