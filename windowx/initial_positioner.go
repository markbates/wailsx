package windowx

type InitialPositioner interface {
	InitHeight() int
	InitPosX() int
	InitPosY() int
	InitWidth() int
}
