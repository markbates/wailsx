package wailstest

import (
	"context"
	"fmt"
)

// StartuperPlugin is a test implementation of the Startuper interface
type StartuperPlugin struct {
	Called bool
	Error  bool
}

func (s *StartuperPlugin) PluginName() string {
	return fmt.Sprintf("%T", s)
}

// Save marks the plugin as called
// If Error is true, it returns the ERR error
func (s *StartuperPlugin) Startup(ctx context.Context) error {
	if s == nil {
		return fmt.Errorf("state is nil")
	}

	s.Called = true

	if s.Error {
		return ErrTest
	}

	return nil
}

// WasCalled returns true if the plugin was called
func (s *StartuperPlugin) WasCalled() bool {
	if s == nil {
		return false
	}
	return s.Called
}
