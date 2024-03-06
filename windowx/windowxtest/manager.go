package windowxtest

import "github.com/markbates/wailsx/windowx"

func NewManager() *windowx.Manager {
	m := &windowx.Manager{}
	m.MaximiseManager = &windowx.Maximiser{}
	m.PositionerManager = &PositionManger{}

	return m
}
