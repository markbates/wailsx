package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
)

type ThemeData struct {
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

func (th ThemeData) PluginName() string {
	return fmt.Sprintf("%T", th)
}

func (th ThemeData) StateData(ctx context.Context) (statedata.Data[ThemeData], error) {
	return statedata.Data[ThemeData]{
		Name: ThemeStataDataName,
		Data: th,
	}, nil
}
