package wailsx

import (
	"context"
	"fmt"
	"sync"
)

// var _ statedata.StateDataProvider[*AppData] = &AppData{}
var _ AppStateDataProvider = &AppData{}

type AppData struct {
	AppName string         `json:"app_name,omitempty"` // application name
	API     *APIData       `json:"api,omitempty"`
	Plugins map[string]any `json:"plugins,omitempty"`

	mu sync.RWMutex
}

func (ad *AppData) StateData(ctx context.Context) (*AppData, error) {
	return ad, nil
}

func (ad *AppData) PluginName() string {
	return fmt.Sprintf("%T", ad)
}
