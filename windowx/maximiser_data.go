package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.DataProvider[*MaximiserData] = &MaximiserData{}

const (
	WINDOW_FULLSCREEN = "fullscreen"
	WINDOW_MAXIMISED  = "maximised"
	WINDOW_MINIMISED  = "minimised"
	WINDOW_NORMAL     = ""
)

type MaximiserData struct {
	Layout string `json:"layout,omitempty"`

	mu sync.RWMutex
}

func (md *MaximiserData) SetFullscreen() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_FULLSCREEN

	return nil
}

func (md *MaximiserData) SetMaximised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_MAXIMISED

	return nil
}

func (md *MaximiserData) SetMinimised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_MINIMISED

	return nil
}

func (md *MaximiserData) SetNormal() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_NORMAL

	return nil
}

func (md *MaximiserData) SetUnfullscreen() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_NORMAL

	return nil
}

func (md *MaximiserData) SetUnmaximised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_NORMAL

	return nil
}

func (md *MaximiserData) SetUnminimised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	md.Layout = WINDOW_NORMAL

	return nil
}

func (md *MaximiserData) ToggleMaximised() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

	md.mu.Lock()
	defer md.mu.Unlock()

	if md.Layout == WINDOW_MAXIMISED {
		md.Layout = WINDOW_NORMAL
		return nil
	}

	md.Layout = WINDOW_MAXIMISED

	return nil
}

func (md *MaximiserData) StateData(ctx context.Context) (*MaximiserData, error) {
	if md == nil {
		return md, fmt.Errorf("maximiser data is nil")
	}

	md.mu.RLock()
	defer md.mu.RUnlock()

	sd := &MaximiserData{
		Layout: md.Layout,
	}

	return sd, nil
}

func (md *MaximiserData) PluginName() string {
	return fmt.Sprintf("%T", md)
}
