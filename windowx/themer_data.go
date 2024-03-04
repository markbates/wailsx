package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

type ThemerData struct {
	BackgroundColour struct {
		R uint8
		G uint8
		B uint8
		A uint8
	}
	IsDarkTheme   bool
	IsLightTheme  bool
	IsSystemTheme bool
}

func (th ThemerData) PluginName() string {
	return fmt.Sprintf("%T", th)
}

func (th ThemerData) StateData(ctx context.Context) (statedata.StateData[ThemerData], error) {
	return statedata.StateData[ThemerData]{
		Name: ThemerStateDataName,
		Data: th,
	}, nil
}
