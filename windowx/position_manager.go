package windowx

import (
	"context"
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
	StateData(ctx context.Context) (*PositionData, error)
}

type RestorablePositionManager interface {
	PositionManager
	RestorePosition(ctx context.Context, data *PositionData) error
}
