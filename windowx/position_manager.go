package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type PositionManager interface {
	WindowCenter(ctx context.Context) error
	WindowGetPosition(ctx context.Context) (int, int, error)
	WindowGetSize(ctx context.Context) (int, int, error)
	WindowSetMaxSize(ctx context.Context, width int, height int) error
	WindowSetMinSize(ctx context.Context, width int, height int) error
	WindowSetPosition(ctx context.Context, x int, y int) error
	WindowSetSize(ctx context.Context, width int, height int) error
}

type PositionManagerDataProvider interface {
	PositionManager
	StateData(ctx context.Context) (statedata.Data[*PositionData], error)
}
