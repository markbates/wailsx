package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/markbates/plugins"
)

var _ Saver = &State{}
var _ Shutdowner = &State{}
var _ Startuper = &State{}
var _ plugins.Plugin = &State{}

type State struct {
	Emitter   // emit save events
	*Position // position of the state

	Name    string `json:"name,omitempty"` // application name
	Plugins plugins.Plugins

	// save function, if nil, save to file in ~/.config/<name>/state.json
	SaveFn func(ctx context.Context) error `json:"-"`

	// startup function, if nil, load from file in ~/.config/<name>/state.json
	StartupFn func(ctx context.Context) error `json:"-"`

	// shutdown function, if nil, call Save
	ShutdownFn func(ctx context.Context) error `json:"-"`

	mu sync.RWMutex
}

func NewState(name string, plugins ...plugins.Plugin) (*State, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}

	st := &State{
		Name:     name,
		Emitter:  NewEmitter(),
		Plugins:  plugins,
		Position: NewPosition(),
	}

	return st, nil
}

func (st *State) Save(ctx context.Context) (err error) {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	st.mu.RLock()
	name := st.Name
	fn := st.SaveFn
	st.mu.RUnlock()

	if len(name) == 0 {
		return fmt.Errorf("name is required: %+v", st)
	}

	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if fn == nil {
		fn = st.saveToFile
	}

	return fn(ctx)
}

func (st *State) Startup(ctx context.Context) (err error) {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	st.mu.RLock()
	name := st.Name
	fn := st.StartupFn
	st.mu.RUnlock()

	if len(name) == 0 {
		return fmt.Errorf("name is required: %+v", st)
	}

	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if fn == nil {
		fn = st.loadFromFile
	}

	if err := fn(ctx); err != nil {
		return err
	}

	for _, p := range st.Plugins {
		if p == nil {
			continue
		}

		if s, ok := p.(Startuper); ok {
			if err := s.Startup(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (st *State) Shutdown(ctx context.Context) (err error) {
	if st == nil {
		return nil
	}

	st.mu.RLock()
	fn := st.ShutdownFn
	st.mu.RUnlock()

	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if fn == nil {
		fn = st.Save
	}

	return fn(ctx)
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

	pos := st.Position
	if pos == nil {
		pos = NewPosition()
	}

	mm := map[string]any{
		"name":     st.Name,
		"position": pos,
	}

	for _, p := range st.Plugins {
		if p, ok := p.(StateDataPlugin); ok {
			sd, err := p.StateData()
			if err != nil {
				return nil, err
			}

			mm[sd.Name] = sd.Data
		}
	}

	return mm, nil
}

func (st *State) PluginName() string {
	return fmt.Sprintf("%T", st)
}

func (st *State) saveToFile(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fp := filepath.Join(home, ".config", st.Name)

	err = os.MkdirAll(fp, 0755)
	if err != nil {
		return err
	}

	fp = filepath.Join(fp, "state.json")

	f, err := os.Create(fp)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(st); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func (st *State) loadFromFile(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fp := filepath.Join(home, ".config", st.Name, "state.json")

	f, err := os.Open(fp)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(f)
	if err := dec.Decode(st); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}
