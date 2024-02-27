package wailstest

import "context"

type PositionCatcher struct {
	X int
	Y int
	W int
	H int
}

func (pc *PositionCatcher) WindowGetPosition(ctx context.Context) (int, int) {
	return pc.X, pc.Y
}

func (pc *PositionCatcher) WindowGetSize(ctx context.Context) (int, int) {
	return pc.W, pc.H
}

func (pc *PositionCatcher) WindowSetPosition(ctx context.Context, x int, y int) {
	pc.X = x
	pc.Y = y
}

func (pc *PositionCatcher) WindowSetSize(ctx context.Context, w int, h int) {
	pc.W = w
	pc.H = h
}
