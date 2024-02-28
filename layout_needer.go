package wailsx

type LayoutNeeder interface {
	SetLayout(pos *Layout) error
}
