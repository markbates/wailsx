package windowxtest

import "fmt"

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

func (pm *PositionManger) WindowCenter() error {
	if pm == nil {
		return fmt.Errorf("position manager is nil")
	}

	pm.IsCentered = true
	return nil
}

func (pm *PositionManger) WindowGetPosition() (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("position manager is nil")
	}

	return pm.X, pm.Y, nil
}

func (pm *PositionManger) WindowGetSize() (int, int, error) {
	if pm == nil {
		return 0, 0, fmt.Errorf("position manager is nil")
	}

	return pm.W, pm.H, nil
}

func (pm *PositionManger) WindowSetMaxSize(width int, height int) error {
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

func (pm *PositionManger) WindowSetMinSize(width int, height int) error {
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

func (pm *PositionManger) WindowSetPosition(x int, y int) error {
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

func (pm *PositionManger) WindowSetSize(width int, height int) error {
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
