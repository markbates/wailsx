package windowx

import (
	"context"
)

type MaximiseManager interface {
	WindowFullscreen(ctx context.Context) error
	WindowIsFullscreen(ctx context.Context) (bool, error)
	WindowIsMaximised(ctx context.Context) (bool, error)
	WindowIsMinimised(ctx context.Context) (bool, error)
	WindowIsNormal(ctx context.Context) (bool, error)
	WindowMaximise(ctx context.Context) error
	WindowMinimise(ctx context.Context) error
	WindowToggleMaximise(ctx context.Context) error
	WindowUnfullscreen(ctx context.Context) error
	WindowUnmaximise(ctx context.Context) error
	WindowUnminimise(ctx context.Context) error
}

type MaximiseManagerDataProvider interface {
	MaximiseManager
	StateData(ctx context.Context) (*MaximiserData, error)
}

type RestorableMaximiseManager interface {
	MaximiseManager
	RestoreMaximiser(ctx context.Context, data *MaximiserData) error
}
