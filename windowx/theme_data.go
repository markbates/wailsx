package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
)

type ThemeData struct {
	BackgroundColour Colour `json:"background_colour,omitempty"`
	IsDarkTheme      bool   `json:"is_dark_theme,omitempty"`
	IsLightTheme     bool   `json:"is_light_theme,omitempty"`
	IsSystemTheme    bool   `json:"is_system_theme,omitempty"`

	mu sync.RWMutex
}

func (th *ThemeData) SetDarkTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.IsDarkTheme = true
	th.IsLightTheme = false
	th.IsSystemTheme = false

	return nil
}

func (th *ThemeData) SetLightTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.IsDarkTheme = false
	th.IsLightTheme = true
	th.IsSystemTheme = false

	return nil
}

func (th *ThemeData) SetSystemTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.IsDarkTheme = false
	th.IsLightTheme = false
	th.IsSystemTheme = true

	return nil
}

func (th *ThemeData) SetBackgroundColour(R, G, B, A uint8) error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.BackgroundColour.R = R
	th.BackgroundColour.G = G
	th.BackgroundColour.B = B
	th.BackgroundColour.A = A

	return nil
}

func (th *ThemeData) PluginName() string {
	return fmt.Sprintf("%T", th)
}

func (th *ThemeData) StateData(ctx context.Context) (statedata.Data[*ThemeData], error) {
	return statedata.Data[*ThemeData]{
		Name: ThemeStataDataName,
		Data: th,
	}, nil
}
