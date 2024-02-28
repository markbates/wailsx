package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/internal/safe"
)

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

	if fn == nil {
		fn = st.saveToFile
	}

	err = safe.Run(func() error {
		return fn(ctx)
	})

	if err != nil {
		return err
	}

	return st.saverPlugins(ctx)
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

func (st *State) saverPlugins(ctx context.Context) error {
	if st == nil {
		return fmt.Errorf("state is nil")
	}

	var wg safe.Group

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
