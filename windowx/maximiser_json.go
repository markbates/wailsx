package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.StateDataProvider[MaximiserData] = MaximiserData{}

type MaximiserData struct {
	IsFullscreen bool `json:"is_fullscreen,omitempty"`
	IsMaximised  bool `json:"is_maximised,omitempty"`
	IsMinimised  bool `json:"is_minimised,omitempty"`
	IsNormal     bool `json:"is_normal,omitempty"`
}

func (md MaximiserData) StateData(ctx context.Context) (statedata.StateData[MaximiserData], error) {

	const name = "maximiser"

	return statedata.StateData[MaximiserData]{
		Name: name,
		Data: md,
	}, nil
}

func (md MaximiserData) PluginName() string {
	return fmt.Sprintf("%T", md)
}

func (md *MaximiserData) SetFullscreen() error {
	if md == nil {
		return fmt.Errorf("maximiser data is nil")
	}

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

	md.IsFullscreen = false
	md.IsMaximised = false
	md.IsMinimised = false
	md.IsNormal = true

	return nil
}
