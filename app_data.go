package wailsx

import (
	"context"
	"fmt"
)

var _ AppStateDataProvider = &AppData{}

type AppData struct {
	AppName string         `json:"app_name,omitempty"` // application name
	API     *APIData       `json:"api,omitempty"`
	Plugins map[string]any `json:"plugins,omitempty"`
}

func (ad *AppData) StateData(ctx context.Context) (*AppData, error) {
	if ad == nil {
		return nil, fmt.Errorf("app data is nil")
	}

	return ad, nil
}

func (ad *AppData) PluginName() string {
	return fmt.Sprintf("%T", ad)
}
