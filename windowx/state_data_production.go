//go:build wails || dev || desktop || production

// when not built with wails, the stubs are used
package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

func (mm *Maximiser) StateData(ctx context.Context) (statedata.Data[*MaximiserData], error) {
	sd := statedata.Data[*MaximiserData]{
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

func (pm *Positioner) StateData(ctx context.Context) (statedata.Data[*PositionData], error) {
	sd := statedata.Data[*PositionData]{
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
