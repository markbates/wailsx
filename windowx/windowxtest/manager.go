package windowxtest

import "github.com/markbates/wailsx/windowx"

func NewManager() *windowx.Manager {
	m := &windowx.Manager{}
	m.MaximiseManager = &windowx.Maximiser{}
	m.PositionManager = &PositionManger{}

	return m
}
