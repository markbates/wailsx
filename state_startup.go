package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/markbates/plugins"
)

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
