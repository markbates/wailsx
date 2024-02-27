package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type State struct {
	*Position

	Name string `json:"name,omitempty"` // application name

	// save function, if nil, save to file in ~/.config/<name>/state.json
	SaveFn func(ctx context.Context) error `json:"-"`

	// load function, if nil, load from file in ~/.config/<name>/state.json
	LoadFn func(ctx context.Context) error `json:"-"`

	mu sync.RWMutex
}

func NewState(name string) (*State, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}

	st := &State{
		Name:     name,
		Position: NewPosition(),
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

	mm := map[string]any{
		"name":     st.Name,
		"position": st.Position,
	}

	return mm, nil
}

func (st *State) Save(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	st.mu.Lock()
	defer st.mu.Unlock()

	if len(st.Name) == 0 {
		return fmt.Errorf("name is required: %+v", st)
	}

	if st.SaveFn != nil {
		return st.SaveFn(ctx)
	}

	return st.saveToFile()
}

func (st *State) Load(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	st.mu.Lock()
	defer st.mu.Unlock()

	if len(st.Name) == 0 {
		return fmt.Errorf("name is required: %+v", st)
	}

	if st.LoadFn != nil {
		return st.LoadFn(ctx)
	}

	return st.loadFromFile()
}

func (st *State) saveToFile() error {
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

func (st *State) loadFromFile() error {
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
