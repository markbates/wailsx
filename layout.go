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
		x: PosX,
		y: PosY,
		w: PosW,
		h: PosH,
	}
}

type Layout struct {
	LayoutManager

	x int
	y int
	w int
	h int

	mu sync.RWMutex
}

// X returns the x position of the window
func (ly *Layout) X() int {
	if ly == nil {
		return PosX
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.x == 0 {
		return PosX
	}

	return ly.x
}

// Y returns the y position of the window
func (ly *Layout) Y() int {
	if ly == nil {
		return PosY
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.y == 0 {
		return PosY
	}

	return ly.y
}

// W returns the width of the window
func (ly *Layout) W() int {
	if ly == nil {
		return PosW
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.w == 0 {
		return PosW
	}

	return ly.w
}

// H returns the height of the window
func (ly *Layout) H() int {
	if ly == nil {
		return PosH
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	if ly.h == 0 {
		return PosH
	}

	return ly.h
}

// Set allows for manual setting of the position and size of the window
// this will call the Layout method to apply the changes
func (ly *Layout) Set(ctx context.Context, x, y, w, h int) error {
	if ly == nil {
		return fmt.Errorf("layout is nil")
	}

	ly.mu.Lock()

	ly.x = x
	ly.y = y
	ly.w = w
	ly.h = h

	ly.mu.Unlock()

	return ly.Layout(ctx)
}

// Update updates the position and size of the window
// based on the current running application.
func (ly *Layout) Update(ctx context.Context) error {
	if ly == nil {
		return fmt.Errorf("layout is nil")
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

	ly.x = x
	ly.y = y
	ly.w = w
	ly.h = h

	return nil
}

// Layout sets the position and size of the window
func (ly *Layout) Layout(ctx context.Context) error {
	if ly == nil {
		ly = NewLayout()
	}

	ly.mu.RLock()
	defer ly.mu.RUnlock()

	err := ly.WindowSetPosition(ctx, ly.x, ly.y)
	if err != nil {
		return err
	}

	err = ly.WindowSetSize(ctx, ly.w, ly.h)
	if err != nil {
		return err
	}

	return nil
}

// PluginName returns the name of the plugin
// implements the plugins.Plugin interface
func (ly *Layout) PluginName() string {
	return fmt.Sprintf("%T", ly)
}

// StateData returns the state data of the plugin
// implements the StateDataProvider interface
func (ly *Layout) StateData() (StateData, error) {
	if ly == nil {
		ly = NewLayout()
	}

	return StateData{
		Name: "position",
		Data: ly,
	}, nil
}

func (ly *Layout) MarshalJSON() ([]byte, error) {
	if ly == nil {
		return json.Marshal(NewLayout())
	}

	return json.Marshal(map[string]int{
		"x": ly.X(),
		"y": ly.Y(),
		"w": ly.W(),
		"h": ly.H(),
	})
}
