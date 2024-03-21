package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

var _ AppStateDataProvider = AppData{}

type AppData struct {
	AppName string         `json:"app_name,omitempty"` // application name
	API     *APIData       `json:"api,omitempty"`
	Plugins map[string]any `json:"plugins,omitempty"`
}

func (ad AppData) StateData(ctx context.Context) (statedata.Data[AppData], error) {
	return statedata.Data[AppData]{
		Data: ad,
	}, nil
}

func (ad AppData) PluginName() string {
	return fmt.Sprintf("%T", ad)
}
