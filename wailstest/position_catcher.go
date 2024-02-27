package wailstest

import (
	"context"
	"fmt"
)

type PositionCatcher struct {
	X int
	Y int
	W int
	H int
}

func (pc *PositionCatcher) WindowGetPosition(ctx context.Context) (int, int, error) {
	if pc == nil {
		return 0, 0, fmt.Errorf("catcher is nil")
	}

	return pc.X, pc.Y, nil
}

func (pc *PositionCatcher) WindowGetSize(ctx context.Context) (int, int, error) {
	if pc == nil {
		return 0, 0, fmt.Errorf("catcher is nil")
	}

	return pc.W, pc.H, nil
}

func (pc *PositionCatcher) WindowSetPosition(ctx context.Context, x int, y int) error {
	if pc == nil {
		return fmt.Errorf("catcher is nil")
	}

	pc.X = x
	pc.Y = y
	return nil
}

func (pc *PositionCatcher) WindowSetSize(ctx context.Context, w int, h int) error {
	if pc == nil {
		return fmt.Errorf("catcher is nil")
	}

	pc.W = w
	pc.H = h
	return nil
}
