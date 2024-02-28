package wailsx

type PositionNeeder interface {
	SetPosition(pos *Position) error
}
