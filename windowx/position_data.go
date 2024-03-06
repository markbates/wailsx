package windowx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.DataProvider[*PositionData] = &PositionData{}

type PositionData struct {
	IsCentered bool
	X          int
	Y          int
	W          int
	H          int
	MaxW       int
	MaxH       int
	MinW       int
	MinH       int

	my sync.Mutex
}

func (pd *PositionData) SetCentered() error {
	if pd == nil {
		return fmt.Errorf("positioner data is nil")
	}

	pd.my.Lock()
	defer pd.my.Unlock()

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

	pd.my.Lock()
	defer pd.my.Unlock()

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

	pd.my.Lock()
	defer pd.my.Unlock()

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

	pd.my.Lock()
	defer pd.my.Unlock()

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

	pd.my.Lock()
	defer pd.my.Unlock()

	pd.MinW = w
	pd.MinH = h
	return nil
}

func (pd *PositionData) PluginName() string {
	return fmt.Sprintf("%T", pd)
}

func (pd *PositionData) StateData(ctx context.Context) (statedata.Data[*PositionData], error) {
	return statedata.Data[*PositionData]{
		Name: PositionerStateDataName,
		Data: pd,
	}, nil
}
