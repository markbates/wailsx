package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/markbates/plugins"
	"golang.org/x/sync/errgroup"
)

var _ Saver = &State{}
var _ Shutdowner = &State{}
var _ Startuper = &State{}
var _ plugins.Plugin = &State{}

type State struct {
	Emitter // emit save events
	*Layout // layout of the app

	Name    string          // application name
	Plugins plugins.Plugins // plugins for the state

	// save function, if nil, save to file in ~/.config/<name>/state.json
	SaveFn func(ctx context.Context) error

	// startup function, if nil, load from file in ~/.config/<name>/state.json
	StartupFn func(ctx context.Context) error

	// shutdown function, if nil, call Save
	ShutdownFn func(ctx context.Context) error

	mu sync.RWMutex
}

func NewState(name string, plugins ...plugins.Plugin) (*State, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}

	st := &State{
		Name:    name,
		Emitter: NewEmitter(),
		Plugins: plugins,
		Layout:  NewLayout(),
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

	if err := fn(ctx); err != nil {
		return err
	}

	return st.saverPlugins(ctx)
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

	ems := plugins.ByType[EmitNeeder](st.Plugins)
	for _, em := range ems {
		if err := em.SetEmitter(st.Emitter); err != nil {
			return err
		}
	}

	lms := plugins.ByType[LayoutNeeder](st.Plugins)
	for _, lm := range lms {
		if err := lm.SetLayout(st.Layout); err != nil {
			return err
		}
	}

	sts := plugins.ByType[Startuper](st.Plugins)
	for _, s := range sts {
		if err := s.Startup(ctx); err != nil {
			return err
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

	if err := fn(ctx); err != nil {
		return err
	}

	sps := plugins.ByType[Shutdowner](st.Plugins)

	for _, sp := range sps {
		if err := sp.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
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

func (st *State) saverPlugins(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	var wg errgroup.Group

	sps := plugins.ByType[Saver](st.Plugins)

	for _, s := range sps {
		s := s
		wg.Go(func() error {
			return s.Save(ctx)
		})
	}

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}

func (st *State) stateDataPlugins() ([]StateData, error) {
	if st == nil {
		return nil, fmt.Errorf("state is nil")
	}

	var list []StateData

	var mu sync.Mutex

	var wg errgroup.Group

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
