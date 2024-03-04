//go:build wails || dev || desktop || production

// when not built with wails, the stubs are used
package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

func (mm *MaximiseManager) StateData(ctx context.Context) (statedata.StateData[*MaximiserData], error) {
	sd := statedata.StateData[*MaximiserData]{
		Name: MaximiserStateDataName,
	}

	if mm == nil {
		return sd, fmt.Errorf("maximiser manager is nil")
	}

	data := mm.data

	isFullscreen, err := wailsrun.WindowIsFullscreen(ctx)
	if err != nil {
		return sd, err
	}
	data.IsFullscreen = isFullscreen

	isMaximised, err := wailsrun.WindowIsMaximised(ctx)
	if err != nil {
		return sd, err
	}

	data.IsMaximised = isMaximised

	isMinimised, err := wailsrun.WindowIsMinimised(ctx)
	if err != nil {
		return sd, err
	}

	data.IsMinimised = isMinimised
	if !isMinimised {
		isNormal, err := wailsrun.WindowIsNormal(ctx)
		if err != nil {
			return sd, err
		}
		data.IsNormal = isNormal
	}

	sd.Data = &data
	return sd, nil
}

func (pm *PositionManger) StateData(ctx context.Context) (statedata.StateData[*PositionerData], error) {
	sd := statedata.StateData[*PositionerData]{
		Name: "positioner",
	}

	if pm == nil {
		return sd, fmt.Errorf("positioner manager is nil")
	}

	data := pm.data

	x, y, err := wailsrun.WindowGetPosition(ctx)
	if err != nil {
		return sd, err
	}

	data.X = x
	data.Y = y

	w, h, err := wailsrun.WindowGetSize(ctx)
	if err != nil {
		return sd, err
	}

	data.W = w
	data.H = h

	sd.Data = &data

	return sd, nil
}

func (th ThemeManager) StateData(ctx context.Context) (statedata.StateData[ThemerData], error) {
	return th.data.StateData(ctx)
}
