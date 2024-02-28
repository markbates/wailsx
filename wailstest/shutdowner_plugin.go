package wailstest

import (
	"context"
	"fmt"
)

// ShutdownerPlugin is a test implementation of the Shutdowner interface
type ShutdownerPlugin struct {
	Called bool
	Error  bool
}

func (s *ShutdownerPlugin) PluginName() string {
	return fmt.Sprintf("%T", s)
}

// Shutdown marks the plugin as called
// If Error is true, it returns the ERR error
func (s *ShutdownerPlugin) Shutdown(ctx context.Context) error {
	if s == nil {
		return fmt.Errorf("state is nil")
	}

	s.Called = true

	if s.Error {
		return ERR
	}

	return nil
}
