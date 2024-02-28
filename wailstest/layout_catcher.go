package wailstest

import (
	"context"
	"fmt"
)

type LayoutCatcher struct {
	X int
	Y int
	W int
	H int
}

func (pc *LayoutCatcher) WindowGetPosition(ctx context.Context) (int, int, error) {
	if pc == nil {
		return 0, 0, fmt.Errorf("catcher is nil")
	}

	return pc.X, pc.Y, nil
}

func (pc *LayoutCatcher) WindowGetSize(ctx context.Context) (int, int, error) {
	if pc == nil {
		return 0, 0, fmt.Errorf("catcher is nil")
	}

	return pc.W, pc.H, nil
}

func (pc *LayoutCatcher) WindowSetPosition(ctx context.Context, x int, y int) error {
	if pc == nil {
		return fmt.Errorf("catcher is nil")
	}

	pc.X = x
	pc.Y = y
	return nil
}

func (pc *LayoutCatcher) WindowSetSize(ctx context.Context, w int, h int) error {
	if pc == nil {
		return fmt.Errorf("catcher is nil")
	}

	pc.W = w
	pc.H = h
	return nil
}