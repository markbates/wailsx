package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

const (
	PosX = 100
	PosY = 100
	PosW = 1200
	PosH = 800
)

func NewPosition() *Position {
	return &Position{
		X: PosX,
		Y: PosY,
		W: PosW,
		H: PosH,
	}
}

type Position struct {
	Positioner

	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`

	mu sync.RWMutex
}

func (pos *Position) PosX() int {
	if pos == nil {
		return PosX
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	if pos.X == 0 {
		return PosX
	}

	return pos.X
}

func (pos *Position) PosY() int {
	if pos == nil {
		return PosY
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	if pos.Y == 0 {
		return PosY
	}

	return pos.Y
}

func (pos *Position) Width() int {
	if pos == nil {
		return PosW
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	if pos.W == 0 {
		return PosW
	}

	return pos.W
}

func (pos *Position) Height() int {
	if pos == nil {
		return PosH
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	if pos.H == 0 {
		return PosH
	}

	return pos.H
}

func (pos *Position) MarshalJSON() ([]byte, error) {
	if pos == nil {
		fmt.Println("nil position")
		return json.Marshal(NewPosition())
	}

	return json.Marshal(map[string]int{
		"x": pos.PosX(),
		"y": pos.PosY(),
		"w": pos.Width(),
		"h": pos.Height(),
	})
}

func (pos *Position) Update(ctx context.Context) {
	if pos == nil {
		return
	}

	pos.mu.Lock()
	defer pos.mu.Unlock()

	x, y := pos.WindowGetPosition(ctx)
	w, h := pos.WindowGetSize(ctx)

	pos.X = x
	pos.Y = y
	pos.W = w
	pos.H = h
}

func (pos *Position) Layout(ctx context.Context) {
	if pos == nil {
		pos = NewPosition()
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	pos.WindowSetPosition(ctx, pos.X, pos.Y)
	pos.WindowSetSize(ctx, pos.W, pos.H)
}
