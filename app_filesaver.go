package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/flect"
)

var _ Saver = &AppFilesaver{}

type AppFilesaver struct {
	App  *App
	Path string

	mu sync.RWMutex
}

func (af *AppFilesaver) Save(ctx context.Context) error {
	if af == nil {
		return fmt.Errorf("AppFilesaver is nil")
	}

	if af.App == nil {
		return fmt.Errorf("AppFilesaver.App is nil")
	}

	s, err := af.filepath()
	if err != nil {
		return err
	}

	af.mu.Lock()
	defer af.mu.Unlock()

	sd, err := af.App.StateData(ctx)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(s), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(s)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(sd); err != nil {
		return err
	}

	return nil
}

func (af *AppFilesaver) filepath() (s string, err error) {
	if af == nil {
		af = &AppFilesaver{}
	}

	af.mu.RLock()
	defer af.mu.RUnlock()

	s = af.Path

	an := "wailsx.json"
	if af.App != nil && len(af.App.Name) > 0 {
		n := flect.Underscore(af.App.Name)
		an = n + ".json"
	}

	if len(s) > 0 {
		if ext := filepath.Ext(s); len(ext) == 0 {
			s = filepath.Join(s, an)
		}
		return s, nil
	}

	s, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}

	s = filepath.Join(s, ".config")

	s = filepath.Join(s, an)

	return s, nil
}
