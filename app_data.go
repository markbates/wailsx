package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

var _ AppStateDataProvider = AppData{}

type AppData struct {
	Name    string         `json:"name,omitempty"` // application name
	API     *APIData       `json:"api,omitempty"`
	Plugins map[string]any `json:"plugins,omitempty"`
}

func (ad AppData) StateData(ctx context.Context) (statedata.Data[AppData], error) {
	return statedata.Data[AppData]{
		Name: AppStateDataProviderName,
		Data: ad,
	}, nil
}
