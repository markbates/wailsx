package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.DataProvider[*PositionData] = &PositionData{}

type PositionData struct {
	IsCentered bool `json:"is_centered,omitempty"`
	X          int  `json:"x,omitempty"`
	Y          int  `json:"y,omitempty"`
	W          int  `json:"w,omitempty"`
	H          int  `json:"h,omitempty"`
	MaxW       int  `json:"max_w,omitempty"`
	MaxH       int  `json:"max_h,omitempty"`
	MinW       int  `json:"min_w,omitempty"`
	MinH       int  `json:"min_h,omitempty"`

	mu sync.Mutex
}

func (pd *PositionData) SetCentered() error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.IsCentered = true
	return nil
}

func (pd *PositionData) SetPosition(x, y int) error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	if x < 0 || y < 0 {
		return fmt.Errorf("x or y is less than 0: %d, %d", x, y)
	}

	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.X = x
	pd.Y = y
	return nil
}

func (pd *PositionData) SetSize(w, h int) error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	if w < 0 || h < 0 {
		return fmt.Errorf("w or h is less than 0: %d, %d", w, h)
	}

	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.W = w
	pd.H = h
	return nil
}

func (pd *PositionData) SetMaxSize(w, h int) error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	if w < 0 || h < 0 {
		return fmt.Errorf("w or h is less than 0: %d, %d", w, h)
	}

	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.MaxW = w
	pd.MaxH = h
	return nil
}

func (pd *PositionData) SetMinSize(w, h int) error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	if w < 0 || h < 0 {
		return fmt.Errorf("w or h is less than 0: %d, %d", w, h)
	}

	pd.mu.Lock()
	defer pd.mu.Unlock()

	pd.MinW = w
	pd.MinH = h
	return nil
}

func (pd *PositionData) PluginName() string {
	return fmt.Sprintf("%T", pd)
}

func (pd *PositionData) StateData(ctx context.Context) (statedata.Data[*PositionData], error) {
	return statedata.Data[*PositionData]{
		Data: pd,
	}, nil
}
