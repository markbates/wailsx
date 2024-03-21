package windowx

import (
	"context"
	"fmt"
	"sync"
)

const (
	THEME_DARK   = "dark"
	THEME_LIGHT  = "light"
	THEME_SYSTEM = ""
)

type ThemeData struct {
	BackgroundColour Colour `json:"background_colour,omitempty"`
	Theme            string `json:"theme,omitempty"`

	mu sync.RWMutex
}

func (th *ThemeData) SetDarkTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.Theme = THEME_DARK

	return nil
}

func (th *ThemeData) SetLightTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.Theme = THEME_LIGHT

	return nil
}

func (th *ThemeData) SetSystemTheme() error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	th.Theme = THEME_SYSTEM

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

func (th *ThemeData) StateData(ctx context.Context) (*ThemeData, error) {
	if th == nil {
		return th, fmt.Errorf("themer is nil")
	}

	th.mu.RLock()
	defer th.mu.RUnlock()

	sd := &ThemeData{
		BackgroundColour: th.BackgroundColour,
		Theme:            th.Theme,
	}

	return sd, nil
}
