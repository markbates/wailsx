package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/markbates/plugins"
)

const (
	PosX = 100
	PosY = 100
	PosW = 1200
	PosH = 800
)

var _ plugins.Plugin = &Layout{}
var _ StateDataProvider = &Layout{}

func NewLayout() *Layout {
	return &Layout{
		X: PosX,
		Y: PosY,
		W: PosW,
		H: PosH,
	}
}

type Layout struct {
	LayoutManager

	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`

	mu sync.RWMutex
}

func (ly *Layout) PosX() int {
	if ly == nil {
		return PosX
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.X == 0 {
		return PosX
	}

	return ly.X
}

func (ly *Layout) PosY() int {
	if ly == nil {
		return PosY
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.Y == 0 {
		return PosY
	}

	return ly.Y
}

func (ly *Layout) Width() int {
	if ly == nil {
		return PosW
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.W == 0 {
		return PosW
	}

	return ly.W
}

func (ly *Layout) Height() int {
	if ly == nil {
		return PosH
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.H == 0 {
		return PosH
	}

	return ly.H
}

func (ly *Layout) MarshalJSON() ([]byte, error) {
	if ly == nil {
		return json.Marshal(NewLayout())
	}

	return json.Marshal(map[string]int{
		"x": ly.PosX(),
		"y": ly.PosY(),
		"w": ly.Width(),
		"h": ly.Height(),
	})
}

func (ly *Layout) Update(ctx context.Context) error {
	if ly == nil {
		return fmt.Errorf("position is nil")
	}

	ly.mu.Lock()
	defer ly.mu.Unlock()

	x, y, err := ly.WindowGetPosition(ctx)
	if err != nil {
		return err
	}

	w, h, err := ly.WindowGetSize(ctx)
	if err != nil {
		return err
	}

	ly.X = x
	ly.Y = y
	ly.W = w
	ly.H = h

	return nil
}

func (ly *Layout) Layout(ctx context.Context) error {
	if ly == nil {
		ly = NewLayout()
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	err := ly.WindowSetPosition(ctx, ly.X, ly.Y)
	if err != nil {
		return err
	}

	err = ly.WindowSetSize(ctx, ly.W, ly.H)
	if err != nil {
		return err
	}

	return nil
}

func (ly *Layout) PluginName() string {
	return fmt.Sprintf("%T", ly)
}

func (ly *Layout) StateData() (StateData, error) {
	if ly == nil {
		ly = NewLayout()
	}

	return StateData{
		Name: "position",
		Data: ly,
	}, nil
}
