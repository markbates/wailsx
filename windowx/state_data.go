//go:build !dev && !desktop && !production && !wails

// when not built with wails, the stubs are used
package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

func (mm *MaximiseManager) StateData(ctx context.Context) (statedata.Data[*MaximiserData], error) {
	return mm.data.StateData(ctx)
}

func (pm *PositionManger) StateData(ctx context.Context) (statedata.Data[*PositionerData], error) {
	return pm.data.StateData(ctx)
}

func (th Themer) StateData(ctx context.Context) (statedata.Data[ThemeData], error) {
	return th.data.StateData(ctx)
}
