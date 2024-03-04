//go:build wails || dev || desktop || production

// when not built with wails, the stubs are used
package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

func (mm MaximiserManager) StateData(ctx context.Context) (statedata.StateData[MaximiserData], error) {
	data := mm.data

	const name = "maximiser"
	sd := statedata.StateData[MaximiserData]{
		Name: name,
	}

	isFullscreen, err := mm.WindowIsFullscreen(ctx)
	if err != nil {
		return sd, err
	}
	data.IsFullscreen = isFullscreen

	isMaximised, err := mm.WindowIsMaximised(ctx)
	if err != nil {
		return sd, err
	}

	data.IsMaximised = isMaximised

	isMinimised, err := mm.WindowIsMinimised(ctx)
	if err != nil {
		return sd, err
	}

	data.IsMinimised = isMinimised
	if !isMinimised {
		isNormal, err := mm.WindowIsNormal(ctx)
		if err != nil {
			return sd, err
		}
		data.IsNormal = isNormal
	}

	sd.Data = data
	return sd, nil
}
