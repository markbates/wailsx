package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.DataProvider[*MaximiserData] = &MaximiserData{}

type MaximiserData struct {
	IsFullscreen bool `json:"is_fullscreen,omitempty"`
	IsMaximised  bool `json:"is_maximised,omitempty"`
	IsMinimised  bool `json:"is_minimised,omitempty"`
	IsNormal     bool `json:"is_normal,omitempty"`

	mu sync.RWMutex
}

func (md *MaximiserData) StateData(ctx context.Context) (statedata.Data[*MaximiserData], error) {
	sd := statedata.Data[*MaximiserData]{
		Name: MaximiserStateDataName,
		Data: md,
	}

	if md == nil {
		return sd, fmt.Errorf("maximiser data is nil")
	}

	return sd, nil
}

func (md *MaximiserData) PluginName() string {
	return fmt.Sprintf("%T", md)
}

func (md *MaximiserData) SetFullscreen() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = true
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = false

	return nil
}

func (md *MaximiserData) SetMaximised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = true
	md.IsMinimised = false
	md.IsNormal = false

	return nil
}

func (md *MaximiserData) SetMinimised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = true
	md.IsNormal = false

	return nil
}

func (md *MaximiserData) SetNormal() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = true

	return nil
}

func (md *MaximiserData) SetUnfullscreen() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = true

	return nil
}

func (md *MaximiserData) SetUnmaximised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = true

	return nil
}

func (md *MaximiserData) SetUnminimised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = true

	return nil
}
