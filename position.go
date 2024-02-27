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
		return json.Marshal(NewPosition())
	}

	return json.Marshal(map[string]int{
		"x": pos.PosX(),
		"y": pos.PosY(),
		"w": pos.Width(),
		"h": pos.Height(),
	})
}

func (pos *Position) Update(ctx context.Context) error {
	if pos == nil {
		return fmt.Errorf("position is nil")
	}

	pos.mu.Lock()
	defer pos.mu.Unlock()

	x, y, err := pos.WindowGetPosition(ctx)
	if err != nil {
		return err
	}

	w, h, err := pos.WindowGetSize(ctx)
	if err != nil {
		return err
	}

	pos.X = x
	pos.Y = y
	pos.W = w
	pos.H = h

	return nil
}

func (pos *Position) Layout(ctx context.Context) error {
	if pos == nil {
		pos = NewPosition()
	}

	pos.mu.RLock()
	defer pos.mu.RUnlock()

	err := pos.WindowSetPosition(ctx, pos.X, pos.Y)
	if err != nil {
		return err
	}

	err = pos.WindowSetSize(ctx, pos.W, pos.H)
	if err != nil {
		return err
	}

	return nil
}
