package windowxtest

import "github.com/markbates/wailsx/windowx"

func NewManager() *windowx.Manager {
	m := &windowx.Manager{}
	m.Maximiser = &windowx.MaximiseManager{}
	m.Positioner = &PositionManger{}

	return m
}
