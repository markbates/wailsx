package windowxtest

import (
	"context"
	"fmt"

	"github.com/markbates/wailsx/windowx"
)

var _ windowx.Maximiser = &MaximiserManager{}

type MaximiserManager struct {
	windowx.MaximiserData
}

func (mm *MaximiserManager) WindowFullscreen(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if mm.IsFullscreen {
		return fmt.Errorf("window is already fullscreen")
	}

	mm.IsFullscreen = true
	mm.IsMaximised = false
	mm.IsMinimised = false
	mm.IsNormal = false
	return nil
}

func (mm MaximiserManager) WindowIsFullscreen(ctx context.Context) (bool, error) {
	return mm.IsFullscreen, nil
}

func (mm MaximiserManager) WindowIsMaximised(ctx context.Context) (bool, error) {
	return mm.IsMaximised, nil
}

func (mm MaximiserManager) WindowIsMinimised(ctx context.Context) (bool, error) {
	return mm.IsMinimised, nil
}

func (mm *MaximiserManager) WindowMaximise(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if mm.IsMaximised {
		return fmt.Errorf("window is already maximised")
	}

	mm.IsFullscreen = false
	mm.IsMaximised = true
	mm.IsMinimised = false
	mm.IsNormal = false
	return nil
}

func (mm *MaximiserManager) WindowMinimise(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if mm.IsMinimised {
		return fmt.Errorf("window is already minimised")
	}

	mm.IsFullscreen = false
	mm.IsMaximised = false
	mm.IsMinimised = true
	mm.IsNormal = false
	return nil
}

func (mm *MaximiserManager) WindowUnfullscreen(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if !mm.IsFullscreen {
		return fmt.Errorf("window is not fullscreen")
	}

	mm.IsFullscreen = false
	mm.IsMaximised = false
	mm.IsMinimised = false
	mm.IsNormal = true
	return nil
}

func (mm *MaximiserManager) WindowUnmaximise(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if !mm.IsMaximised {
		return fmt.Errorf("window is not maximised")
	}

	mm.IsFullscreen = false
	mm.IsMaximised = false
	mm.IsMinimised = false
	mm.IsNormal = true
	return nil
}

func (mm *MaximiserManager) WindowUnminimise(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	if !mm.IsMinimised {
		return fmt.Errorf("window is not minimised")
	}

	mm.IsFullscreen = false
	mm.IsMaximised = false
	mm.IsMinimised = false
	mm.IsNormal = true
	return nil
}

func (mm MaximiserManager) WindowIsNormal(ctx context.Context) (bool, error) {
	return mm.IsNormal, nil
}
