package wailstest

import (
	"context"
	"fmt"
)

// SaverPlugin is a test implementation of the Saver interface
type SaverPlugin struct {
	Saved bool
	Error bool
}

func (s *SaverPlugin) PluginName() string {
	return fmt.Sprintf("%T", s)
}

// Save marks the state as saved
// If Error is true, it returns the ERR error
func (s *SaverPlugin) Save(ctx context.Context) error {
	if s == nil {
		return fmt.Errorf("state is nil")
	}

	if s.Error {
		return ERR
	}

	s.Saved = true
	return nil
}
