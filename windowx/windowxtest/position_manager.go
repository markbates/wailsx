package windowxtest

import (
	"context"
	"fmt"
)

type PositionManger struct {
	IsCentered bool
	X          int
	Y          int
	W          int
	H          int
	MaxW       int
	MaxH       int
	MinW       int
	MinH       int
}

func (pm *PositionManger) WindowCenter(ctx context.Context) error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	pm.IsCentered = true
	return nil
}

func (pm *PositionManger) WindowGetPosition(ctx context.Context) (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("position manager is nil")
	}

	return pm.X, pm.Y, nil
}

func (pm *PositionManger) WindowGetSize(ctx context.Context) (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("position manager is nil")
	}

	return pm.W, pm.H, nil
}

func (pm *PositionManger) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	if width < 0 || height < 0 {
		return fmt.Errorf("width or height is less than 0: %d, %d", width, height)
	}

	pm.MaxW = width
	pm.MaxH = height
	return nil
}

func (pm *PositionManger) WindowSetMinSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	if width < 0 || height < 0 {
		return fmt.Errorf("width or height is less than 0: %d, %d", width, height)
	}

	pm.MinW = width
	pm.MinH = height
	return nil
}

func (pm *PositionManger) WindowSetPosition(ctx context.Context, x int, y int) error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	if x < 0 || y < 0 {
		return fmt.Errorf("x or y is less than 0: %d, %d", x, y)
	}

	pm.X = x
	pm.Y = y
	return nil
}

func (pm *PositionManger) WindowSetSize(ctx context.Context, width int, height int) error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	if width < 0 || height < 0 {
		return fmt.Errorf("width or height is less than 0: %d, %d", width, height)
	}

	pm.W = width
	pm.H = height
	return nil
}
